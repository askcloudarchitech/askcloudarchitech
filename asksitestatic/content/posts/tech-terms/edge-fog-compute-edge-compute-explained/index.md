---
title: What even is "the edge"? Fog Compute / Edge Compute Explained
date: 2023-03-14T22:46:34.881Z
description: The Edge is all the rage these days and every tech company out there is trying
  to find a way to tell their consumers that they using edge computing.
draft: ""
videoID: ""
keywords:
  - cloudflare
  - edge computing
  - edge networking
  - fog compute
  - vercel
  - CDN
tags:
  - cloud
  - edge
  - CDN
ID: 2023-03-14T22:46:34.881Z
previewImage: thumb.001.jpeg
mainImage: ""
slug: 
---
{{< img2 src="thumb.001.jpeg" alt="What is 'The Edge'?" >}}

“The Edge” is all the rage these days, and it kinda seems like every tech company out there is trying to find a way to tell their consumers that they are on the bandwagon of “edge computing”. 

As with most of these buzzword tech inventions, the name doesn’t really matter. What matters more is the concept behind it. Once you understand what they mean you may even find that you are using “the edge” or “the fog” (I’ll get into this later) in your day-to-day life, and it’s really not that big of a deal. You may even find that this tech existed way before it had this fancy name. 

So if you put all the hype aside what, exactly, is “the edge”?

Putting all the boring technical definitions aside, the entirety of the whole edge and fog hype can be summed up like this:

> Edge computing is all about processing data in a geographically closer location before sending it back to a centralized location.
> 

That's it. So… that’s the end of my article. Goodnight everyone!

Nah, just kidding. Let’s discuss some examples of edge computing and where the term fog computing comes in. 

Like I mentioned above in my super simplified definition, Edge computing really just defines the idea that computing doesn’t need to be done in one centralized location and then sent or received to or from a network connected device. The term edge computing started being used around the same time as the term Content Delivery Network (CDN). The idea here is that traditional PCs or servers aren’t the only devices connected to the network anymore. But, things have evolved since then and the term “Edge Computing” has kinda separated from the term “Edge Networking”. 

Edge computing these days is now more associated with the Internet of Things (IoT). These devices could be literally anything. Phones, cameras, sensors, household appliances and vehicles are all “things” in IoT. Since all of these things gather data and are expected to do something fun and exciting, they all need to either process information locally or send the information they capture to a server which does some processing and then acts based on what it has computed. 

So with all of these “things” out there, you could imagine that some of them need to react quickly or possibly process data that shouldn’t be passed across the internet to a centralized server. To add to that, the centralized server is probably geographically far away from the thing which adds additional problems with that whole “quickly” part. 

Edge computing aims to solve these issues, and the way it’s solved is by processing data somewhere other than a far away server somewhere up in the cloud. 

## Examples of Edge Computing

Here’s a couple examples:

1. Chick-fil-A - No matter what you think of their politics, Chick-fil-A has some sweet networking and edge computing stuff going on. If you don’t follow the [Chick-fil-A tech medium publication](https://medium.com/chick-fil-atech), you are missing out on some pretty cool stuff. They have built and shipped mini Kubernetes clusters to each of their locations to allow for local compute to happen at the store and call back to home with the details. This is an excellent example of edge computing at its best. 
2. Connected cars - I don’t think that most people really consider it, but your car is probably an edge compute node all on it’t own. For some more technologically advanced cars like Tesla, an absolute ton of data is processed right inside the vehicle. This includes all the advanced safety and convenience features that use cameras and sensors all over the car. Most of this data never leaves the vehicle. Instead, it's all computed right there at the edge with only the distilled data being sent back to the mothership. 

## What is “Fog” Computing?

To be completely honest, this is kinda splitting hairs. It feels like another buzzword that could generate some revenue, so it got jammed into the tech term glossary. Very Simply, Fog Computing is somewhere between edge computing and traditional cloud computing at data centers. The best example i could give would be a company network with a bunch of different sensors like motion sensors, smart lights, monitoring etc. All of these devices may not provide much computing themselves and even tying them together at the edge might be too close to get the big picture. By taking one step out and bring the data together on the network (but not as far as a data center), you can get valuable data and make decisions without consuming as much general internet bandwidth. 

## Confusing Overloaded Use of “Edge”

Another common place where you hear the word “edge” come into play is with edge networking. This term generally applies the same concepts, but is slightly different from your traditional edge computing. Edge networking is often related to Content Delivery Networks that will physically distribute caching servers around the globe. 

With edge networking, when you make a request, for example, to read this article, you are probably not getting the article served to you from the original source. It’s way more likely that you are getting served a cached copy of the article (or at least part of it) from a cache server at the network’s edge. 

Over the last few years, edge networking has gotten more and more advanced. Services like Cloudflare now offer way more than just caching at their edge. They now offer (with some configuration) the ability to distribute parts of your website logic out to the edge as well. This takes the strain off of your poor server that in the past would have to serve all the dynamic content on its own. 

Another edge provider that many web developers will be very familiar with is [Vercel.](https://vercel.com/) Vercel can take an application written with the popular meta-framework [NextJs](https://nextjs.org/) and distribute to it’s edge networks to make your website highly available and distributed with very little extra work. 

## Is it All Just Hype?

Well, yes and no I suppose. All this edge stuff is certainly real and useful tech. If you lump it all together and forget all the hyped up tech lingo, it's really just a way of saying “move the compute power closer to the user”. Getting caught up in the technical definitions seems a little silly when in reality it’s best to just think of the optimal way to make a great experience for the end user. So, if you are asked about edge technology, just think “distributed”. Gone are the days of a simple server sitting on a rack serving an entire experience. To truly compete, you need to make things fast and efficient. To do that, you probably need to use (or are already using) “the Edge”.