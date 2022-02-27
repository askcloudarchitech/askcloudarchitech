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
resources:
  - name: thumbnail
    src: medium-auto-post-cli.001.jpeg
    title: Introducing MediumAutoPost
---
{{<img name="thumbnail" size="large">}}

In a previous article, I outlined how to set up your Hugo site so you can easily syndicate your articles to medium.com. I also provided a Postman collection that would do the heavy lifting of posting your new content. While that worked pretty well, I wanted to do a little better. 

Namely, I wanted to make these improvements:

1. Remove the postman collection and replace it with a CLI tool that could be used in a terminal
2. Remove the need to provide the URL of the post you wanted to send to medium.com
3. Track and store which articles have already been syndicated to avoid posting duplicate content
4. Allow for all of this to work with one simple command 

So... to accomplish that I created mediumautopost! It's a CLI tool written in Go that accomplishes all the improvements I listed above. This article will cover all of the one-time setup steps needed to get this all working. Let's get started!

## Step One: Install mediumautopost

Installing the tool is easy! If you have homebrew installed on your mac simply run these two steps:

```bash
brew tap askcloudarchitech/askcloudarchitech 
brew install mediumautopost
```

If you don't have homebrew or you would prefer to install the application manually you can also grab the binary [here](https://github.com/askcloudarchitech/mediumautopost/releases)

## Step Two: Make a Place to Save Publishing Status

One of the best features of MediumAutoPost is its ability to remember which posts it has already sent to medium.com. Since mediumautopost is designed to work with Hugo and Hugo doesn't really have the ability to save the state of something (since it's a static generated content tool), we need to set up a permanent place for storing the publication status of the individual posts. 

When building the tool I landed on storing this publication information in a git repo. I chose this for a couple of reasons:

1. Github is a pretty universal tool that many people use
2. It's free! If you haven't noticed, everything I'm doing here from web hosting to publications and automation is all completely free on a month-to-month basis. This is important to me for maintaining a very low-cost content publication solution. 

This process is pretty simple. Just go to github.com and 

## Step Three: Populate the environment variables