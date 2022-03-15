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
resources:
  - name: netlify-deploy-file
    title: File location netlify/functions/deploy-succeeded
    src: screen-shot-2022-03-14-at-8.38.22-pm.png
---
If you are hosting your website on Netlify.com, the Netlify Functions feature has the potential to add a ton of functionality to your JAMStack site for no additional cost. Netlify functions are simply serverless functions that can be triggered either by submitting data to a certain URL on your site or by triggered events in your site's build process. 

Recently I needed a way to execute some additional functionality after a successful deployment of my Hugo site hosted on Netlify. To accomplish this, I used the post-deploy Netlify function and it worked great!

## My Use Case

I've been working on a tool that will automatically take any new post I write on my website and post it to medium.com as a draft. I've written a few articles about this tool and you can read them here and here. The last step I needed to complete was the automation that would post my new article when I published it. Previously I had been running the tool manually, and while this worked fine, I am incredibly lazy and just wanted it to work without even thinking about it. 

## Creating the Netlify Function

Netlify's documentation on Netlify Functions is OK, but not great. It took some jumping around and some additional digging to figure out everything I needed to do, but eventually, I figured it all out. Here are the steps I completed as well as some helpful "missing documentation" I encountered along the way. 

### Create the File - Naming is important!

{{< img name="netlify-deploy-file" size="medium" >}}

### Writing the execution steps

{{< gist gmorse81 61c8257d0ccf14cedef078ecf3685ac3 >}}

### Debugging the Payload

{{< gist gmorse81 e3325aa2137578a71a4b73cbaa49f3b9 >}}

## Setting your Environment Variables

## Checking for Proper Execution

```
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
```