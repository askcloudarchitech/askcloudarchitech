---
title: "What Is a CDN and How Do They Work? - Content Delivery Networks Explained"
date: 2022-09-24T10:54:17-04:00
ID: 51f65d5e509a8526cab6f2e669f119e9
draft: false
description: |
  Content delivery networks or CDNs are, at their most basic level, multi-function services designed to deliver the content on your website to users fast.

  No matter which company you choose as a CDN provider, they all offer one service in common: the ability to deliver certain parts of your website’s content faster than if it were served from its origin where you actually host it.

keywords:
  - cdn content delivery networks
  - content delivery network
  - cdn
  - what is cdn
  - tech terms you should know
  - tech terms explained
  - what is a CDN
  - what is a cdn and how does it work
  - what is a cdn
  - content delivery network explained
  - cloudflare
  - cloudfront
  - akamai cdn
  - fastly
  - google cdn
  - how to use a cdn
  - how to set up a CDN
  - do i need a CDN
  - askcloudarchitech
mainImage: CDN 001
videoID: R2UjCU4cd04
resources:
  - name: CDN 001
    title: CDN 001
    src: CDN.001.jpeg
  - name: presentation 001
    title: presentation 001
    src: presentation.001.jpeg
  - name: presentation 002
    title: presentation 002
    src: presentation.002.jpeg
  - name: presentation 003
    title: presentation 003
    src: presentation.003.jpeg
  - name: presentation 004
    title: presentation 004
    src: presentation.004.jpeg
  - name: presentation 005
    title: presentation 005
    src: presentation.005.jpeg
  - name: presentation 006
    title: presentation 006
    src: presentation.006.jpeg
---

Content delivery networks or CDNs are, at their most basic level, multi-function services designed to deliver the content on your website to users fast. There are many CDNs out on the market these days. Some of them are offered as stand-alone services like Cloudflare, Fastly, or Akamai others are offered directly by the big cloud companies themselves like google CDN or AWS Cloudfront.

{{< youtube R2UjCU4cd04 >}}

No matter which company you choose as a CDN provider, they all offer one service in common and that is the ability to deliver certain parts of your website’s content faster than if it were served from its origin where you host it.

Essentially, when your site loads in a browser, all kinds of other stuff like images, CSS files, javascript files, and any other static content is fetched behind the scenes. If you don't have a CDN, all of this content needs to come from your server. This, as I'm sure you can imagine, can put a heavy strain on your server and take a long time to be delivered (especially if the files are large or they are geographically far away from the machine that’s making the request).

{{< img name="presentation 003" showcaption=false >}}

When you introduce a CDN into the mix, requests for these individual static files are instead requested from the CDN network. The first time these files are requested, the CDN fetches them from your server like a normal request and then delivers them to the requesting machine. But… the next time someone requests the same file, the CDN doesn't request it from your server and instead delivers it from its caches.

The benefits of this are twofold. First off, your server now only has to serve these files occasionally since the CDN is serving them for you. Second, the CDN provider will distribute your file to many edge endpoints all around the world so the amount of time it takes to be delivered is dramatically shortened.

So that all sounds pretty cool, but how does that work from a technical standpoint?

There are a couple of different ways that CDNs are set up depending on which one you choose. Some CDNs work by giving you a completely different domain name specifically for delivering your static content. For example, if you had “myawesomesite.com”, you may set up static.myawesomesite.com. The static subdomain points to the CDN and then you instruct the CDN to fetch the content from your actual site and serve the content from its edge. Usually in this scenario when someone loads your site at myawesomesite.com, you would then change the URLs of all of the static content like images and CSS to point to the static subdomain. This allows your dynamic content to be served directly from your server while all of the additional assets are pulled from the CDN network.

{{< img name="presentation 005" showcaption=false >}}

The other way some CDNs work is by actually proxying your whole website, but only caching the things that are cachable. In this scenario, you would point myawsomesite.com to the CDN network and then the CDN provider would evaluate each request and determine if that request should receive a cached file or if it should be passed back to your origin machine to get content served from your server directly.

{{< img name="presentation 006" showcaption=false >}}

No matter which way your CDN provider operates, there are certain things to look out for before you just flip the switch and start enjoying the benefits. mainly, you need to make sure the CDN knows how long to cache a file and if it should be cached in the first place. Every CDN provider has slightly different ways to control their caching rules and most of them offer some sort of control panel for making custom rules. But most of them will also honor the cache headers served by your origin server. Now, caching headers is a completely different topic worthy of another article, but essentially you can specify how long a file should be cached by setting headers in the HTTP response. When a CDN sees these headers, they then know how long the file can be kept before fetching it again.

Finally, while every CDN is primarily a way to deliver content faster, some CDNs offer more robust features than that. For example, Cloudflare has expanded its offering of basic CDN into a full-blown application firewall that allows you to not only speed up your site but also protect it from surges in traffic and even block common exploits like SQL injection.

Ok cool, so CDNs sound great, but are they expensive?

Well, that all depends on what you are looking for. While many of the CDNs I have mentioned here have a cost and often charge based on the amount of data consumed, they aren't necessarily expensive. For example, AWS Cloudfront charges by the volume of data processed. If you don't have a lot of traffic, you don't pay a lot. Pretty simple. But probably even better than that, Cloudflare has a free offering for websites that offers a subset of their total features. You may not get everything they offer, but you do get the basics of DDOS protection and content delivery from the edge which is what most basic sites need anyways. If your site grows, just add more features and cough up the cash.

Alright, I think I'll leave it there. Until next time. Happy coding.
