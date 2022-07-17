---
title: Is That Open Source Project Safe To Use? - Here’s How To Know
date: 2022-06-13T04:31:17.768Z
description: >-
  Developing software of any type will inevitably involve pulling in an
  open-source dependency. But, given that there are millions of open source
  projects out there and you don’t personally know anything about their source
  code or their security practices, how do you know if that code that you just
  pulled into your project is safe? 


  When it comes to security, there is no such thing as perfectly safe software. It’s all about managing risk and making informed decisions on the software you are choosing to include in your project. It would be super easy for me to sit here and say “You need to review every last line of code going into your project”, but the simple truth is that’s practically impossible. So instead, id like to share with you the five phases of software review that I follow when evaluating open source packages, APIs and SDKs.


  Summary of my five phases for open dependency review:

  1. Is it popular?

  2. Is it maintained?

  3. Is security considered?

  4. Is it even necessary?

  5. Maintenance.
draft: false
keywords:
  - open source
  - open source security
  - open source dependencies
  - opensource
  - secure web development
  - software dependencies
  - open source software
  - security
  - open source security risks
  - open source projects
  - what is open source
  - npm
  - composer
  - golang security
  - security review
  - secure code review
  - secure code review tutorial
  - dependabot pull request
  - dependabot
  - github security
tags:
  - open source
  - open source security
  - open source dependencies
  - opensource
  - secure web development
  - software dependencies
  - open source software
  - security
  - open source security risks
  - open source projects
  - what is open source
  - npm
  - composer
  - golang security
  - security review
  - secure code review
  - secure code review tutorial
  - dependabot pull request
  - dependabot
  - github security
ID: 1655094677759
videoID: 8ygb6rIgoKs
mainImage: open-source-security
resources:
  - name: open-source-security
    title: open source security
    src: thumb.001.jpeg
---

Developing software of any type will inevitably involve pulling in an open-source dependency. But, given that there are millions of open source projects out there and you don’t personally know anything about their source code or their security practices, how do you know if that code that you just pulled into your project is safe?

When it comes to security, there is no such thing as perfectly safe software. It’s all about managing risk and making informed decisions on the software you are choosing to include in your project. It would be super easy for me to sit here and say “You need to review every last line of code going into your project”, but the simple truth is that’s practically impossible. So instead, I'd like to share with you the five phases of software review that I follow when evaluating open source packages, APIs and SDKs.

{{< youtube 8ygb6rIgoKs >}}

## **This first part of the review is what I’m going to call the “is it popular” phase**

Because, let’s be honest here, if the project gets a lot of use it’s way more likely that it’s well maintained and that others may find, report, and fix security issues in the project that would otherwise go uncorrected.

Almost all languages written today have some sort of system for pulling in dependencies. For Go, it’s Go modules. For javascript projects it’s NPM. For Php it’s composer, for java it’s maven or Gradle or whatever other java junk you may pull out of the woodwork (I’m kidding java people. the internet still loves you). These websites for these package managers are a great place to start with your investigation into the security and behavior of your dependency.

Take a look at the number of other projects using the package. This number isn’t always 100% accurate, but if it’s a significant number, you are much better off using the dependency compared to that random package you found in the deep corners of the internet with 5 users and 0 reported issues. 0 reported issues is not a good thing, I promise!

Also, make sure the package has been used recently. Ten thousand downloads a year ago but none in the last six months is a sure sign that this package is outdated or unmaintained and you might want to look elsewhere for alternatives.

## **Ok, so now you can see how popular this project is. The next step is the “is it maintained?” phase**

From the project’s page, find the link to the source code. Click the link and go to the source repo. Most projects will be using GitHub, so that’s what I’m going to use as my example, but you can find many of the same details on almost any git repo site like bitbucket or GitLab.

There are a few key things to check on the project’s GitHub page.

First, how many contributors are listed on the project. Again this is a good indicator of the activity on the project.

Next, how many stars and watchers does it have? Do a lot of people have this project bookmarked so they can get back to the source code?

Next, does the project have an openly available issue queue? If so take a look. See what stuff is there. There will always be bugs and feature requests but does anything stand out as fishy? Any unaddressed security issues or any devs screaming about how nobody has merged their stuff in months? These are all signs to look out for.

Before you leave the issue queue, take a quick look at the closed issues. See how long it’s been since an issue was closed. See what types of issues have been resolved. You are specifically looking for a healthy amount of activity on the project which indicates it is actively maintained and people actually care about the security and functionality of the package.

## **The next step is the “is security even considered?” phase**

Take a look to see if the project has the “security” tab activated on their project. Some will openly show known vulnerabilities. others will have a list of security advisories or a security policy. Anything here in this tab is probably a good sign that the maintainers care about security.

Some projects will even have a note in their readme about how to report security issues. Most open-source projects would prefer that they get a shot at fixing an issue before it is publicly announced (for good reason). If the project you are reviewing has this, that’s also a very good sign.

As a final step, you could also do a review of the actual source code of the project. Now, I know sometimes these projects are huge and there is no way you could possibly review the whole thing, but it doesn’t hurt to look at the parts you want to use. Especially if the functionality you need is mission-critical.

As the last step, and this doesn’t apply to every project, you can check third-party sites for bounty programs. Sites like <https://huntr.dev> report known issues that need resolution.

Finally and maybe most importantly, check one of the many CVE databases on the internet. My favorite is <https://osv.dev/>. Just type the name of the package to see a history of CVEs filed against the project and most importantly if it was fixed and how long it took. the site will list the URL of the original report so you can look at more details and the resolutions.

## **The last step is the “is it even necessary?” Phase**

Well, that’s a good question, and it’s worth investigating. Just earlier I mentioned that many dependencies are large a have a ton of functionality. In general, the purpose of pulling in another project is to take advantage of its capabilities instead of re-inventing the wheel. in that case, it’s almost better if the project is large because if it’s small maybe you don’t even need it.

If you find something you think you need because you don’t know how to do it yourself, maybe you can just look at the source code of that project, pick out the pieces you need, and write them yourself with less complexity. For example: do you really need the whole Axios javascript package if you only need to do a simple POST? Probably not.

## **And finally, now we have the “maintenance” phase**

Ok, so you’ve now you’ve done your homework and installed the package. Now what? Well, security never ends so make sure you are doing something to monitor all the dependencies you have installed. The easiest way to do this is to enable Dependabot alerts on your project. It will automatically scan your lock files to get a list of all projects you use and produce a report of any known reported security issues in your dependencies. It will even offer to automatically update your project to the newest version of the projects and resolve security issues for you.

OK, I think I’ll leave it there. If you can take a small amount of time and do these steps before importing a package, it might just save you a ton of headaches in the future. And if you are interested in open source, maybe check out [this article](https://askcloudarchitech.com/posts/self-improvement/how-contribute-open-source-software/). Until next time happy coding!
