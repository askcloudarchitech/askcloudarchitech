---
title: Introducing Medium-auto-post | Syndicate your Hugo Content to Medium.com
date: 2022-02-26T20:39:08.287Z
description: s
draft: false
keywords:
  - s
categories:
  - s
tags:
  - s
ID: 1645907948222
---
In a previous article, I outlined how to set up your Hugo site so you can easily syndicate your articles to medium.com. I also provided a Postman collection that would do the heavy lifting of posting your new content. While that worked pretty well, I wanted to do a little better. 

Namely, I wanted to make these improvements:

1. Remove the postman collection and replace it with a CLI tool that could be used in a terminal
2. Remove the need to provide the URL of the post you wanted to send to medium.com
3. Track and store which articles have already been syndicated to avoid posting duplicate content
4. Allow for all of this to work with one simple command 

So... to accomplish that I created mediumautopost! It's a CLI tool written in Go that accomplished all the improvements I listed above. This article will cover all of the one-time setup steps needed to get this all working. Let's get started!

## Step One: Install mediumautopost

## Step Two: Make a Place to Save Publishing Status

## Step Three: Populate the environment variables