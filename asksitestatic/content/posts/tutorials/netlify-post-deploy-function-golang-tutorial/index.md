---
title: How to Write a Netlify Post-Deploy Function in Golang
date: 2022-03-12T14:26:34.274Z
description: d
draft: false
keywords:
  - d
categories:
  - d
tags:
  - d
ID: 1647095194201
---
If you are hosting your website on Netlify.com, the Netlify Functions feature has the potential to add a ton of functionality to your JAMStack site for no additional cost. Netlify functions are simply serverless functions that can be triggered either by submitting data to a certain URL on your site or by triggered events in your site's build process. 

Recently I needed to a way to execute some additional functionality after a successful deploy of my Hugo site hosted on Netlify. To accomplish this, I used the post-deploy Netlify function and it worked great!

## My Use Case

I've been working on a tool that will automatically take any new post I write on my website and post it to medium.com as a draft. I've written a few articles about this tool and you can read them here and here. The last step i needed to complete was the automation that would post my new article when i published it. Previously I had been running the tool manually, and while this worked fine, I am incredibly lazy and just wanted it to work without even thinking about it. 

## Creating the Netlify Function

Netlify's documentation on Netlify Functions is OK, but not great. It took some jumping around and some additional digging to figure out everything i needed to do, but eventually, I figured it all out. Here are the steps I completed as well as some helpful "missing documentation" I encountered along the way. 

### Create the File - Naming is important!

### Writing the execution steps

{{< gist gmorse81 61c8257d0ccf14cedef078ecf3685ac3 >}}

### Debugging the Payload

{{< gist gmorse81 e3325aa2137578a71a4b73cbaa49f3b9 >}}

## Setting your Environment Variables

## Checking for Proper Execution

```
12:11:09 AM: Build ready to start
12:11:11 AM: build-image version: 6ebfb829398e07eb99ad5455777162ff19838822 (xenial)
12:11:11 AM: build-image tag: v3.10.5
12:11:11 AM: buildbot version: 7eafb394e33f42f945c880ce4ac17c149867813a
12:11:11 AM: Fetching cached dependencies
12:11:11 AM: Failed to fetch cache, continuing with build
12:11:11 AM: Starting to prepare the repo for build
12:11:11 AM: No cached dependencies found. Cloning fresh repo
12:11:11 AM: git clone https://github.com/askcloudarchitech/askcloudarchitech
12:11:13 AM: Preparing Git Reference pull/11/head
12:11:13 AM: Parsing package.json dependencies
12:11:13 AM: Different publish path detected, going to use the one specified in the Netlify configuration file: 'asksitestatic/public' versus 'public' in the Netlify UI
12:11:13 AM: Different functions path detected, going to use the one specified in the Netlify configuration file: 'asksitestatic/netlify/functions' versus 'netlify/functions' in the Netlify UI
12:11:14 AM: Starting build script
12:11:14 AM: Installing dependencies
12:11:14 AM: Python version set to 2.7
12:11:15 AM: v12.18.0 is already installed.
12:11:15 AM: Now using node v12.18.0 (npm v6.14.4)
12:11:15 AM: Started restoring cached build plugins
12:11:15 AM: Finished restoring cached build plugins
12:11:15 AM: Attempting ruby version 2.7.2, read from environment
12:11:16 AM: Using ruby version 2.7.2
12:11:16 AM: Using PHP version 5.6
12:11:16 AM: Installing Hugo 0.89.4
12:11:17 AM: hugo v0.89.4-AB01BA6E+extended linux/amd64 BuildDate=2021-11-17T08:24:09Z VendorInfo=gohugoio
12:11:17 AM: Started restoring cached go cache
12:11:17 AM: Finished restoring cached go cache
12:11:17 AM: go version go1.14.4 linux/amd64
12:11:17 AM: go version go1.14.4 linux/amd64
12:11:17 AM: Installing missing commands
12:11:17 AM: Verify run directory
12:11:19 AM: ​
12:11:19 AM: ────────────────────────────────────────────────────────────────
12:11:19 AM:   Netlify Build                                                 
12:11:19 AM: ────────────────────────────────────────────────────────────────
12:11:19 AM: ​
12:11:19 AM: ❯ Version
12:11:19 AM:   @netlify/build 26.4.0
12:11:19 AM: ​
12:11:19 AM: ❯ Flags
12:11:19 AM:   baseRelDir: true
12:11:19 AM:   buildId: 622ec05de2c402000813fa88
12:11:19 AM:   deployId: 622ec05de2c402000813fa8a
12:11:19 AM: ​
12:11:19 AM: ❯ Current directory
12:11:19 AM:   /opt/build/repo/asksitestatic
12:11:19 AM: ​
12:11:19 AM: ❯ Config file
12:11:19 AM:   No config file was defined: using default values.
12:11:19 AM: ​
12:11:19 AM: ❯ Context
12:11:19 AM:   deploy-preview
12:11:19 AM: ​
12:11:19 AM: ────────────────────────────────────────────────────────────────
12:11:19 AM:   1. Build command from Netlify app                             
12:11:19 AM: ────────────────────────────────────────────────────────────────
12:11:19 AM: ​
12:11:19 AM: $ hugo
12:11:19 AM: Start building sites …
12:11:19 AM: hugo v0.89.4-AB01BA6E+extended linux/amd64 BuildDate=2021-11-17T08:24:09Z VendorInfo=gohugoio
12:11:25 AM:                    | EN
12:11:25 AM: -------------------+------
12:11:25 AM:   Pages            |  95
12:11:25 AM:   Paginator pages  |   0
12:11:25 AM:   Non-page files   |  22
12:11:25 AM:   Static files     | 111
12:11:25 AM:   Processed images |  88
12:11:25 AM:   Aliases          |  41
12:11:25 AM:   Sitemaps         |   1
12:11:25 AM:   Cleaned          |   0
12:11:25 AM: Total in 5676 ms
12:11:25 AM: ​
12:11:25 AM: (build.command completed in 5.7s)
12:11:25 AM: ​
12:11:25 AM: ────────────────────────────────────────────────────────────────
12:11:25 AM:   2. Functions bundling                                         
12:11:25 AM: ────────────────────────────────────────────────────────────────
12:11:25 AM: ​
12:11:25 AM: Packaging Functions from netlify/functions directory:
12:11:25 AM:  - deploy-succeeded/main.go
12:11:25 AM: ​
12:11:32 AM: ​
12:11:32 AM: (Functions bundling completed in 7.5s)
12:11:32 AM: ​
12:11:32 AM: ────────────────────────────────────────────────────────────────
12:11:32 AM:   3. Deploy site                                                
12:11:32 AM: ────────────────────────────────────────────────────────────────
12:11:32 AM: ​
12:11:32 AM: Starting to deploy site from 'asksitestatic/public'
12:11:33 AM: Creating deploy tree 
12:11:33 AM: Creating deploy upload records
12:11:33 AM: 0 new files to upload
12:11:33 AM: 1 new functions to upload
12:11:39 AM: Site deploy was successfully initiated
12:11:39 AM: ​
12:11:39 AM: (Deploy site completed in 6.9s)
12:11:39 AM: ​
12:11:39 AM: ────────────────────────────────────────────────────────────────
12:11:39 AM: Starting post processing
12:11:39 AM:   Netlify Build Complete                                        
12:11:39 AM: ────────────────────────────────────────────────────────────────
12:11:39 AM: ​
12:11:39 AM: (Netlify Build completed in 20.2s)
12:11:39 AM: Post processing - HTML
12:11:39 AM: Caching artifacts
12:11:39 AM: Started saving build plugins
12:11:39 AM: Finished saving build plugins
12:11:39 AM: Started saving pip cache
12:11:39 AM: Finished saving pip cache
12:11:39 AM: Started saving emacs cask dependencies
12:11:39 AM: Finished saving emacs cask dependencies
12:11:39 AM: Started saving maven dependencies
12:11:39 AM: Finished saving maven dependencies
12:11:39 AM: Started saving boot dependencies
12:11:40 AM: Processing form - questionform
12:11:40 AM: Finished saving boot dependencies
12:11:40 AM: Started saving rust rustup cache
12:11:40 AM: Finished saving rust rustup cache
12:11:40 AM: Detected form fields:
12:11:40 AM: Started saving go dependencies
12:11:40 AM: Finished saving go dependencies
12:11:40 AM:  - name
12:11:40 AM:  - email
12:11:40 AM:  - hp1
12:11:40 AM:  - question
12:11:41 AM: Post processing - header rules
12:11:41 AM: Post processing - redirect rules
12:11:41 AM: Post processing done
12:11:41 AM: Build script success
12:11:44 AM: Site is live ✨
12:11:56 AM: Finished processing build request in 44.868646347s
```