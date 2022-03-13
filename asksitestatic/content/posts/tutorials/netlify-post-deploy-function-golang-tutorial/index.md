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

Netlify's documentation on Netlify Functions is OK, but not great. It took some jumping around and some additional digging to figure out everything i needed to do, but eventually I figured it all out. Here are the steps I completed as well as some helpful "missing documentation" I encountered along the way. 

### Create the File - Naming is important!

### Writing the execution steps

### Debugging the Payload

## Setting your Environment Variables

## Checking for Proper Execution