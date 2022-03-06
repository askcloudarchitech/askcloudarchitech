---
title: Introducing Medium-auto-post | Syndicate your Hugo Content to Medium.com
date: 2022-02-26T20:39:08.287Z
description: Download and install mediumautopost to automatically publish
  articles from your go-hugo website to medium.com with one simple command. This
  article explains how to set up mediumautopost and start syndicating your
  content to medium.com today.
draft: false
keywords:
  - mediumautopost
  - medium.com
  - content syndication
  - go-hugo
  - hugo
  - hugo content syndication
categories:
  - development
  - tutorials
  - hugo
tags:
  - go-hugo
  - hugo
  - content syndication
  - time saving tools
  - medium REST API
ID: 1645907948222
resources:
  - name: thumbnail
    src: medium-auto-post-cli.001.jpeg
    title: Introducing MediumAutoPost
  - name: new-personal-access-token
    title: Create a personal access token on github
    src: new-personal-access-token.png
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

## Step Two: Make a Place to Save the Publishing Status

One of the best features of MediumAutoPost is its ability to remember which posts it has already sent to medium.com. Since mediumautopost is designed to work with Hugo and Hugo doesn't really have the ability to save the state of something (since it's a static generated content tool), we need to set up a permanent place for storing the publication status of the individual posts. 

When building the tool I landed on storing this publication information in a git repo. I chose this for a couple of reasons:

1. Github is a pretty universal tool that many people use
2. It's free! If you haven't noticed, everything I'm doing here from web hosting to publications and automation is all completely free on a month-to-month basis. This is important to me for maintaining a very low-cost content publication solution. 

This process is pretty simple. Just go to github.com and create a new repo. I called mine medium-publish-status and you can [see it here](https://github.com/askcloudarchitech/medium-publish-status) if you would like to see an example. You don't even need to add any files to this repo. MediumAutoPost will do all of that for you. 

Next, you will need to create a Github personal access token so the program has permission to update the repo. Go to https://github.com/settings/tokens and hit the "generate new token" button. Provide a name in the note field, select an expiration date and check the box next to Repo then hit "Generate token" at the bottom. Copy the token and put it somewhere safe for now. NOTE: Keep this token a secret. 

{{<img name="new-personal-access-token" size="medium">}}

I should also add, if you don't want to do the Github status setup, you can also store the status file locally on your computer. I will cover that in the next step. Just a word of warning: be careful if you choose the local storage option. If you lose the file, the program will not know what posts it has already posted to medium.com. 

## Step Three: Populate the environment variables

The final step before running the command is to create the settings file. Here's an example of the file:

```shell
MEDIUM_ENDPOINT_PREFIX="https://api.medium.com/v1"
MEDIUM_BEARER_TOKEN="get token from medium. paste here"
WEBSITE_JSON_INDEX_URL="path to your json index file"
GITHUB_PERSONAL_TOKEN="generate a personal access token and paste here"
GITHUB_STATUS_REPO_OWNER="your Github account name"
GITHUB_STATUS_REPO="repo name for storing the status of posts to medium.com"
STORAGE_TYPE=""
STORAGE_FILE_PATH="/OPTIONAL/PATH/TO/STATUS/FILE.json"
```

These settings instruct MediumAutoPost on how to pull your articles and post them to medium.com. Let's run through the lines and discuss each one. 

MEDIUM_ENDPOINT_PREFIX: This is just the beginning of the URL for the medium API. You can leave it as-is  \
MEDIUM_BEARER_TOKEN: [In a previous article](https://askcloudarchitech.com/posts/tutorials/auto-generate-post-payload-medium-com/#step-1-generate-a-mediumcom-access-token), I explained how to get a medium bearer token. Put your token here  \
WEBSITE_JSON_INDEX_URL: Again, in the first article, I outlined this step. Put the URL of your index file here  \
GITHUB_PERSONAL_TOKEN: If you are using Github for storing your publication status as mentioned above, put your personal access token here  \
GITHUB_STATUS_REPO_OWNER: This should be your Github user name. In my case, for example, it's askcloudarchitech  \
GITHUB_STATUS_REPO: This should be the name of the repo you created using the steps outlined above  \
STORAGE_TYPE: If using Github to store your status, you can leave this empty. If you want to use the local file method, set this to "FILE"  \
STORAGE_FILE_PATH: If using the local file method, type the path of where you want the status file to be saved

Now that you have the settings file created, just save it somewhere and note the path to the file. 

## Step Four: Run the tool and post your content

The last step is to just run the command. From your terminal, run `mediumautopost -e /path/to/your/.env`, replacing the path in the command with the actual path of your settings file you created above. If all goes well, you should see some output telling you that your posts have been submitted!

Now head over to medium.com and log into your account. Go to your drafts sections and you should see your posts there. You can make any necessary tweaks then publish. 

Finally, check your status repo. There should be a file named status.json which contains the details of each article that was posted to medium. This file will be used on future runs to ensure these same posts are not published twice.