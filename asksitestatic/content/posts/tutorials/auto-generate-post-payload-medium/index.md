---
title: "Auto Generate a Medium.com REST API Payload to Syndicate Posts with Hugo"
slug: "auto-generate-post-payload-medium-com"
date: 2022-02-06T19:26:26-05:00
draft: false
ID: ashdjfklasdf978asdfjklasdfy78sadbfjkasd678
description: Easily syndicate your website or blog articles to medium.com with hugo by automatically generating an API payload for the Medium REST API.
keywords:
  - medium.com
  - medium REST API
  - content syndication
  - reusable content
  - golang
  - go hugo
  - hugo
categories:
  - Development
  - Tutorials
  - Hugo
tags:
  - medium.com
  - REST
  - content syndication
  - reusable content
  - golang
  - hugo
  - JAMSTACK
videoID: cY50uFTY7WM
mainImage:
resources:
  - name: collection-run-output
    src: "collection-run-output.png"
    title: Output after collection run
  - name: draft-screen
    src: "draft-screen.png"
    title: Medium.com drafts screen
  - name: integration-token
    src: "integration-token.png"
    title: Medium.com integration token screen
  - name: postman-collection-requests
    src: "postman-collection-requests.png"
    title: View of requests in postman collection
  - name: postman-collection-variables
    src: "postman-collection-variables.png"
    title: Postman collection variables to populate
  - name: run-collection-button
    src: "run-collection-button.png"
    title: Postman run collection button
  - name: run-collection-dropdown
    src: "run-collection-dropdown.png"
    title: Collection action menu with run collection option
---

{{< youtube cY50uFTY7WM >}}

If you are looking to get the most of every article you write, syndicating your content to Medium.com can help you get your content in front of more eyeballs regularly. When I write articles I usually start by posting them to my own website first, then taking the finished product and posting it to Medium.com. I then supply a canonical URL in the Medium.com article settings so search engines don’t see this post as duplicate content.

Medium has an auto-import tool that will attempt to read the content from your website and automatically create a Medium.com post based on what it finds. While this tool works sometimes, it’s not always the best. Sometimes it works perfectly, other times it completely fails. Because of this, I decided to create a simpler way to get my website articles onto Medium with minimum effort and far fewer errors.

My website is powered by Hugo, a static site generation CMS which allows you to have a fully functional site with no database or complicated hosting plan. My website is hosted for free on netlify.com because it is built with Hugo. Using Hugo’s powerful custom output formats feature, I was able to easily create the exact information Medium.com needs. Then it's just as simple as POSTing this information using the Medium API and voila! Your website post is now on Medium.com with very little additional tweaking required. So, without further ado, let’s jump into how to update your Hugo site to create this output format and start saving some time!

NOTE: Even if you don't use Hugo for your website's CMS, you can most likely create the same results on your CMS of choice. You can [skip to the second half](#two) of this article if you don't use Hugo.

## Modifying my Hugo site to generate the proper JSON output

### Step 1: In your Hugo site's config file, add the new custom output format.

```toml
[outputFormats]
  [outputFormats.MediumJSON]
    baseName = 'medium'
    isPlainText = true
    mediaType = 'application/json'

[outputs]
  page = ['HTML', 'MediumJSON']
  section = ['HTML', 'MediumJSON']
```

On my site, you can find this file here: https://github.com/askcloudarchitech/askcloudarchitech/blob/master/asksitestatic/config.toml. If you are new to Hugo, this file contains all of the site's global settings and some settings that control the theme.

By adding the `[outputFormats]` snippet of configuration you are instructing the Hugo site generator to look for `mediumjson.json` layout files which the site will use to generate the necessary output. To learn more about Hugo custom output formats, check out the [Hugo docs site](https://gohugo.io/templates/output-formats/). The second section called `[outputs]` instructs Hugo to generate the mediumjson output for both pages and sections of the site. This is all we will need in the settings to make this work.

### Step 2: Add the ID field to your posts archetypes

```yaml
---
title: "{{ replace .Name "-" " " | title }}"
slug: ""
date: {{ .Date }}
ID: {{ .UniqueID }}
```

https://github.com/askcloudarchitech/askcloudarchitech/blob/master/asksitestatic/archetypes/posts.md

This is a bit of foreshadowing to the next article in this series. For each of my website posts, I wanted to generate a unique ID that would permanently identify a specific post. This way later if I wanted to automate the process of publishing all of my posts to Medium.com (hint hint), a program could keep track of the IDs so it doesn't post the same article twice. You don't need to do this step, but it takes about 10 seconds so it certainly doesn't hurt.

### Step 3: Create the layout template which outputs the medium article POST data

Before diving into the how-to part of this step, it's important to first look at the [medium REST API docs](https://github.com/Medium/medium-api-docs#33-posts). The docs clearly outline what POST data you need in the body of your request to create a new article on Medium.com. Here's the example from the docs:

```
POST /v1/users/5303d74c64f66366f00cb9b2a94f3251bf5/posts HTTP/1.1
Host: api.medium.com
Authorization: Bearer 181d415f34379af07b2c11d144dfbe35d
Content-Type: application/json
Accept: application/json
Accept-Charset: utf-8

{
  "title": "Liverpool FC",
  "contentFormat": "html",
  "content": "<h1>Liverpool FC</h1><p>You’ll never walk alone.</p>",
  "canonicalUrl": "http://jamietalbot.com/posts/liverpool-fc",
  "tags": ["football", "sport", "Liverpool"],
  "publishStatus": "public"
}
```

The JSON body in the request above is what I am trying to replicate for each of my website's articles. To do this with Hugo, you create a new layout in the `layouts` directory and there's not much to it. You can see the layout [here](https://github.com/askcloudarchitech/askcloudarchitech/blob/master/asksitestatic/layouts/_default/single.mediumjson.json) in my website's git repo if you'd like. Here's what the contents of the file look like:

```
{
  "title": {{ .Title | jsonify }},
  "contentFormat": "html",
  "content": {{ printf "<h1>%s</h1> %s" .Title .Content | jsonify }},
  "canonicalUrl": {{ .Permalink | jsonify }},
  "tags": {{ .Params.tags | jsonify }}
}
```

As you can see, this is the same exact layout as the example from the Medium docs, but the values have been replaced with Hugo template data. Additionally, each of the variable values has been piped through the `jsonify` function to ensure they are properly escaped. This file should be named `single.mediumjson.json` and placed into your `layouts/_default` directory. This will ensure a `medium.json` file will be generated for every article created on your site.

### Step 4: Create an index layout template to get a full list of all articles

The second and final layout file for the site is a Hugo `list` template that will create an index that references each of the articles. This will not be used right now but will be very important later in part two of this series when we want to automate this entire process.

Again in the `layouts/_default` directory, create a new file called `list.mediumjson.json` and put the following content in the file.

```
[
    {{ range $i, $e := .RegularPagesRecursive -}}
    {{- if $i -}},
    {{ end -}}
    {
        "url": "{{ .Permalink }}medium.json",
        "id": "{{ if .Params.Id }}{{ .Params.Id }}{{end}}"
    }{{ end }}
]
```

This template file will loop through all the content on the site making a full listing of all the articles' full URLs and their unique IDs. You can see this file [here in my website's GIT repo](https://github.com/askcloudarchitech/askcloudarchitech/blob/master/asksitestatic/layouts/_default/list.mediumjson.json).

### Step 5: Publish your changes and try it out

OK. All that's left is to try it. You can either run `hugo serve` or push your changes up to your live site and then try it out. Here are the links to some examples on my site.

Index page: [https://askcloudarchitech.com/posts/medium.json]()  
Article page: [https://askcloudarchitech.com/posts/tutorials/homebrew-setup-and-usage/medium.json]()

<hr/> <a name="two"></a>

## Using the generated JSON output and Postman to post your article using the Medium REST API

As mentioned at the top of this article, if you aren't using Hugo as your CMS, you can likely still accomplish this same result using your blogging platform or CMS. Just set up your CMS of choice to generate the following files:

1. A JSON index file of your contents
   1. Example JSON format: https://askcloudarchitech.com/posts/medium.json
   1. Example HTML format (just for reference): https://askcloudarchitech.com/posts
1. A JSON output for each of your articles in Medium POST format
   1. Example JSON format: https://askcloudarchitech.com/posts/tutorials/homebrew-setup-and-usage/medium.json
   1. Example HTML article format (just for reference): https://askcloudarchitech.com/posts/tutorials/homebrew-setup-and-usage/

The next steps will explain how to take that generated data and easily post it to medium.com in a couple easy steps.

### Step 1: Generate a Medium.com access token

To use the Medium.com API, you first need a Medium access token. This token will be used in the next step by Postman so it can act on your behalf and post your article. Just follow these steps to get your token.

1. Go to https://medium.com/me/settings and then click on "integration tokens"
1. Enter a short description of the token so you remember what its for. Something simple like "Medium autopost tool" will work just fine.
1. Click "Get integration token" and the token will be generated and printed on the page.
1. Highlight and copy the token. Put it somewhere safe for now. NOTE: Do not share this token.  
   {{<img name="integration-token" size="medium" >}}

### Step 2: Download my pre-made Postman collection which does everything for you.

Instead of giving you a bunch of command line CURL commands to make all of this work, I have created a Postman collection that ties together the three necessary requests and does it all with one click of a button.

If you don't know what Postman is, it's a tool generally used for testing APIs or other web requests to see if they fit your needs. I have repurposed the tool to run the necessary Medium.com API requests in sequence and gather the necessary information along with way.

You can download Postman here: https://www.postman.com/downloads/  
You can get my workspace here: https://www.postman.com/morsecode/workspace/askcloudarchitech

Before we jump into using it, let me first explain what this collection does.

1. The first step in the collection uses your Medium integration token and connects to the `User` endpoint to get your unique user ID. It stores this in a variable for a future step.
2. The second step in the collection connects to your website and downloads the JSON payload data of the article you want to post to Medium.com. This is using the JSON info you created in the first half of this article. It then stores the data it retrieved to variables for the final step.
3. The third step in the collection puts all the pieces together. It uses your integration token and your Medium.com user ID to connect to the Medium `posts` API endpoint. It sends the information it collected in step two. If all is successful, it gets a `201` response indicating that the post was created

Note: I have set this collection up to only create drafts on Medium.com so don't worry about automatically posting stuff publicly!

### Step 3: Add the necessary values to the Postman collection variables and run the collection

If you have successfully downloaded my workspace, you should have something that looks like this in Postman:  
{{<img name="postman-collection-requests" size="medium" >}}

To get this whole thing to work, you only need to populate two variables. Click on the collection name then click on the variables tab. You will see two variables to fill in. It should look like this:  
{{<img name="postman-collection-variables" size="medium">}}

In the `Current Value` column on the `mediumToken` row, paste your Medium integration token (I have hidden mine in the screenshot). In the `Current Value` column on the `articleURL` row, past the full URL of the JSON Medium output of one of your articles. Save your changes if necessary.

Next, hover over the collection name in the right column of Postman and click the three dots. Choose the option to "Run collection".  
{{<img name="run-collection-dropdown" size="medium">}}

A new window will open on the left with some options.  
{{<img name="run-collection-button" size="medium">}}

On this screen make sure all three requests are checked as shown in the screenshot then click the "Run Medium API POST" button. This will kick off the process.

After this has finished you should see a screen showing that all the requests ran. I all went well, the last step will show a status code of 201 indicating that the article was successfully posted.  
{{<img name="collection-run-output" size="medium">}}

### Step 4: Check Medium.com for your draft and do your final tweaking

Now that this has finished, head over to https://medium.com/me/stories/drafts to see your draft. You should see something like this:  
{{<img name="draft-screen" size="medium">}}

There are a few things that are not supported by the API, so make sure to check the following.

1. Add a subtitle if necessary
1. Add an SEO title
1. Check that your tags applied correctly
1. Proofread your content to make sure it all came over correctly
1. Publish your draft so it will be available publicly

<hr/>

## A few gotchas and what's next

When I was first setting up this process, I noticed that my images were broken and that some links didn't work. I found that my CMS was using relative URLs for images and links on the site. To fix this I updated the code on my site to ensure it always uses full URLs. This may also be an issue for you so keep an eye out for that.

As I mentioned above, this is part one of a three-part series. In the next part, I will be eliminating the need for using Postman altogether and replacing it with a custom program. This program will:

1. Keep track of articles that have already been posted to Medium.com
2. Replace all the Postman stuff with one simple command in the terminal
3. Eliminate the need for you to supply the URL of the article you want to post. Using the index listing it will auto-discover new articles.

Of course, I will share this program as an open-source tool for all to use as well as go over all the code to show how it works.
