---
title: Taking Developer Efficiency to the Extreme
date: 2023-03-16T04:00:55.131Z
description: You know what developers like to do? Write code. Do you know what business
  like their developers doing? Writing code.
draft: ""
videoID: ""
keywords:
  - developer
  - developer efficiency
  - efficiency
  - managers
  - time
  - save time
tags:
  - continuous delivery
  - Development
  - automation
ID: 2023-03-16T04:00:55.131Z
previewImage: joshua-reddekopp-SyYmXSDnJ54-unsplash.jpg
mainImage: ""
---

{{< img2 src="joshua-reddekopp-SyYmXSDnJ54-unsplash.jpg" alt="Be Obsessed with developer efficiency" >}}

You know what developers like to do? Write code. Do you know what business like their developers doing? Writing code. Do you know what both businesses and developers absolutely hate? Inefficiency! 

Was that too obvious? Well, you might think so at the surface, but efficiency is a deep rabbit hole and can be defined in a ton of different ways in the context of development. As a matter of fact, it’s a topic I am deeply passionate about, and you should be too. So today, let me ask for a few minutes of your time, so we can go down the efficiency rabbit hole together. 

Are you ready?

I don’t think I could ever sum up just how obsessed I am with developer efficiency in just one article. I think that spending time making it, so a developer can do more coding and less other garbage is the secret to success in software projects. So to attempt to scratch the surface, today I am going to discuss the top three places where devs unnecessarily waste most of their time and usually end up absolutely thrashing the timeline of a development project. 

## Three Ways to Improve Developer Efficiency

### 1. Local Development Environment Setup

At the risk of having you all leave and stop reading this article, I’d like you to go to GitHub and look at some random projects. How many of them have the most incredibly ass backwards instructions on how to actually get a project working locally? Not only is this probably the number one reason a developer would choose not to contribute to a project, but it's also a process that collectively adds hundreds of hours to a business driven software project. The most terrible examples go one of two ways:

1. There are no instructions and the only way you will ever get up and running locally is to talk to someone else that already knows how or…
2. There are instructions, but they are outdated, only work on certain machines, or only work half the time. 

From a developer's perspective, if you spend half a day (or even half an hour for that matter) fiddling around with dependencies and local development tools before you can even start the task you’ve been assigned, you start your day on the completely wrong foot. You also end up looking like you code super slowly because of all the time you are wasting just getting set up.

From a management perspective, projects that don’t have a simple setup process might as well be burning your money. Its repetitive toil that results in 0 gained output and never gets any better. 

So here’s my advice on how to fix the local environment setup. Actually here are two options. 

1. Use docker containers - Docker containers for development standardize the dependency stack and lock in all the dependency versions, so nobody has to worry about it. As an additional bonus it also greatly reduces (but doesn’t eliminate) the “works on my machine” problem. 
2. Makefile or setup scripts - If the nature of your project isn’t container friendly, take the time to write a script that will set up the project for a first time developer. This process can be a little more complicated than it seems at the surface but its absolutely worth the effort. Try to think about things like dependency versions, different machine architectures like mac vs windows, and the fact that the person doing the work probably has to work on multiple projects at once.

Now before I move on from this one, I think there are a couple very important call-outs to make. The steps I listed above should not be confused with “local build instructions”. Local build instructions are a blueprint for how to take the source code and turn it into a final working project on your machine. Just being able to build from source is NOT what I’m talking about here. I’m talking about the ability to set up a project to actually contribute to it. This usually includes well written ways to build the project locally, but it’s more than that.

When you set up local development instructions, think about how a developer would actually work on the code. Usually with a containerized approach this means having the source volume mounted to the container. This way, when changes are made, there can be an even quicker way of just seeing the changes without having to completely rebuild the container. 

### 2. Testing and Preparing Demos

After completing a code task, there’s nothing more frustrating than having to manually prepare a demo environment or a space for QA to do their thing. This type of work is super time-consuming and its just another distraction that pulls devs away from their actual job of writing code. I think for a lot of companies this is part of a dev’s job description but why? If they can complete the task to spec and then move on to the next, that is a much better use of their time and a company’s dollars. 

So… how do you avoid wasting a bunch of time preparing a nice place to test? Automate!

Let me set the bar for you. A developer should be able to finish their task, open a pull request and walk away. If they have to do anything else you are wasting time on repetitive work that is getting you nowhere. Now where the magic happens is with all the automated stuff that happens after the dev walks away. 

1. Automated unit tests are run
2. Static code analysis is done
3. A test environment is automatically built 
4. Integration and end-to-end tests are automatically run
5. The product manager is notified that the environment is ready for a “self demo”

If any of these processes fail, the developer is notified and can come back to fix whatever happened. If the self demo leads to questions or concerns, the product manager can reach out to the dev with questions. 

Believe me, the last thing anyone wants is to submit work that doesn’t meet requirements. You’d be surprised how much effort someone will put in to avoid having to re-explain their work. You should try it.

### 3. Requirement Misunderstandings

This one is definitely one of those “an ounce prevention is worth a pound of cure” scenarios. It’s also the only scenario in this list where I am recommending that a developer take their hands off the keyboard and just have a simple conversation. 

It’s amazing to me how many companies do the exact opposite of that. They write requirements in a vacuum with no developer input or Q and A session. They then follow up in a demo (see above) to clarify the finer points of the request AFTER the work as been completed. This is completely upside down. 

For a manager, the solution to this one should be pretty self-explanatory. When you want to make something, talk to the person who will be making it. Have a real meaningful conversation about the goals of the project. Talk about what you want, what you don’t want and where you may be flexible. 

For a developer, trying to get this started at your place of work might be a challenge. If you do manage to get the team to try it, don’t blow it by being a jerk. Really listen to what the company is trying to accomplish. Make suggestions where time could be saved or where doing work now will pay off in the long run, without sounding like you are over-engineering (but that’s a different topic)

## This is Just The Beginning

If you just read all of this and thought “hey, these sound like good ideas” I promise this is just scratching the surface. Each of the three methods above have a ton of finer details that just couldn’t go into this post (due to attention span constraints). If you really start to look hard at where developers are spending time, you can begin to see the wasted hours, minutes and even seconds that all add up. So forge forward, and take a dive down the rabbit hole. Just see how deep it goes.