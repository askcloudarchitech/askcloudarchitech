---
title: Your README is a Dumpster Fire - It's Time to Fix That
date: 2022-04-24T21:50:57.931Z
description: Let's be honest for a minute, all of your READMEs are
  afterthoughts. They are a chore that needs to be done. Well, today IS the day
  that you could turn all of that around. Today you could choose to give your
  READMEs all of the respect that they deserve! Come with me on a journey into
  your README dumpster fire. What it is now, what it could be, and what it could
  mean for you and your project when done CORRECTLY.
draft: false
keywords:
  - README
  - software development
  - open source
  - documentation
  - README.md
  - README SEO
  - SEO
  - github
  - git
  - documentation
  - coding
  - software
  - description
  - code examples
  - code contribution
  - open source
  - open source software
  - development
  - web development
tags:
  - README
  - development
  - documentation
  - best practices
ID: 1650837057854
mainImage: readme
resources:
  - name: readme
    src: reademe.md-2-.png
    title: README FIRE
---
It's time for a little self-reflection. 

Take a little time to think back over **every single project you have ever created**. OK, now right at this moment, what are you thinking about? Are you thinking about the code? Are you remembering the language, the functions, the classes? Or, are you think about what the project did? **Are you thinking about the problem that it solved?** 

Well, if you are like most people in this world, you are thinking about the latter. After all the code is written and the tool you poured your heart and soul into is done, **the code itself doesn't really matter, does it?** So that begs the question: if the problem you solved is more important than how you solved it, why is it that the **one single file** in your project that actually focuses on the functionality is the one that's most commonly ignored? You know what file I'm talking about, right? **The README!** 

Let's be honest for a minute, **all of your READMEs are afterthoughts**. They are a chore that needs to be done. Well, today IS the day that you could **turn all of that around**. Today you could choose to give your READMEs all of the **respect that they deserve**! Come with me on a journey into your README dumpster fire. What it is now, what it could be, and what it could mean for you and your project when **done CORRECTLY.** 

{{< img name="readme" size="large" >}}

## A brief history of bad READMEs

Before we dive into why you want a good README and how to make one, let's take a look at some of the most notoriously bad READMEs out there in the wastelands of forgotten projects. 

**The one-liner -** We all know this one because you essentially get this README for free just by checking the "create a README" box in Github when you create a repo. This is the one with **JUST the title** of the project and **NOTHING else.** If you have a one-liner README, at least you have one. I guess that's a win?

**The ghost -** If there is anything **worse than the one-liner**, it's the ghost. In this case, you couldn't even be bothered to check the box in Github to auto-generate your README. The **README literally doesn't exist.** If your project doesn't have a README, it's just a useless pile of code clogging up the interwebs. 

**The over-explainer -** Ah yes, the **exact opposite of the ghost** is the over-explainer. You know, the one where the README is so **long and complicated** that nobody would ever possibly take the time to read the whole thing? READMEs are made to be a summary, not all of your comments compiled into a doc.

**The out-of-date -** This one is **all-to-common**. You **started with good intentions** and wrote a README worth actually reading, but then, slowly, everything started to fall apart. A refactor of a whole portion of the app wasn't documented. A brand new set of config was completely forgotten. The architecture diagram is completely out of whack. This README is quite possibly the worst because it can be **completely miseading**. The poor soul who tried to follow your **three-year-old instructions** hates you right now. 

**The broken promise -** I'll spare you the dramatics and simply leave you with this: "README coming soon!" *Facepalm.*

## Benefits of a well-written README

I hope the section above made you feel both **ashamed of yourself** for the complete dumpster fire you have spread across the internet as well as **motivated** to make it right! A well-written README has the potential to be **so much more than just another piece of documentation**. Let's take a moment to consider the **massive benefits** of a README written with purpose.

**SEO benefit -** When people search for solutions, they don't exactly go to the old googly box and start typing code examples. What do they do, they type out the functionality they are looking for. If you can write a  good README with a strong, **keyword-rich explanation** of your project's functionality, there's a good chance **search engines will pick it up** and put it in front of your future users/contributors. 

**Gain users -** You want your super awesome project to be **popular**, right? Well, without a readme (or really any of those READMEs listed above) you won't exactly be pulling in your next biggest fans. As a matter of fact, you may actually be confusing them or scaring them away. **Write a solid README and the users will come!**

**Gain contributors -** You can't continue to make this totally awesome project of yours completely on your own (yeah I'm talking to you), **you need contributors!** And... if you want to gain contributors you need people who like what you made so much that **they want to help** make it better. READMEs serve **two purposes** here: 1 **attract** your future contributor and, 2 **explain** how to actually contribute (but more on that later)

**Remind yourself later -** I'm not the only one with this problem, right? You made something cool and you know generally that it exists, but **forget some of the important details**. If there is one reason to write a README it might be this. Write it for your forgetful **future self**. Seriously, you'll get old too. And you will forget stuff. 

**Reduce reported issues -** Steam is starting to come out of my ears just thinking about this. Do you want to know the most **aggravating** part of software development? When you know that your project has a bug or **known issue** and someone reports it to you like you don't already know. Having someone explain to you in great detail that known issue you already wasted like 30 hours on is THE WORST. So... avoid that situation altogether by **simply adding a "known issues"** section to your README. This will work like a charm (unless that annoying issue reporter also doesn't read).

**Make money? -** hey, why not?! You can totally make money from a README. More on that in the next section. 

This last one isn't exactly a benefit, but just some **advice**. Just In case you are a Ghoster or a One-liner from above, I think it's important to point out that README files are written in Markdown. **Markdown is super easy to learn** and might even be worth committing to memory. Don't know markdown? [Here's a great doc](https://www.markdownguide.org/) on how it all works.

## How to craft a useful, well-written README

I think if you've made it this far through this post, you have gotta be super PUMPED about getting into this README action! Now that you've seen the failures and you know all the benefits, are you ready to learn, exactly, how to structure a README masterpiece? Let's (finally) get into the details. Here's the list, in order, of **elements you should have in your README**.

1. **A strong H1 title and an H2 subtitle** - Just like writing an article or a blog post, you need a great title and subtitle to attract **search engines and humans**. It doesn't need to be the name of your project, but it does help if your title includes the name of the project. 
2. **An intro paragraph focused on what the project *does*** *\-* Write an intro paragraph about what this project is, what it does, and how it's used. This section is still for SEO purposes and for keeping it simple about the **value your project provides** to the user who is searching for it. 
3. **Diagram (optional) -**  If necessary, add a **diagram** showing where this project fits and how it works. If it's a CLI tool or a graphical tool, this would be a great opportunity to add an **animated GIF** of your project in action. Even better, adding a **youtube video** demo of your project to your README could be very beneficial to gaining more users. 
4. **Installation and usage instructions (for end-users) -** Now it's time to get a little bit nerdier. If a user has gotten this far into your README, you bet there's a chance they actually **want to use your project**. Give **instructions** on how to install or use the tool. Don't get this confused with how to contribute to this project (like help improve the code), that's the next section. This section should only talk about how to be a **consumer of the project**.
5. **Installation and usage instructions (for contributors) -** Ya know the best part of open source projects? If you make something really cool, **others will want to help make it better**! In this section of the README, give instructions on how to pull the code down and start up the tool for **development purposes**. This section is usually pretty technical and may require instruction on how to build from source, but hopefully, you have a script for MAKEFILE from stuff like that. Anything you can do to make the development experience easier will help you **gain more contributors**. 
6. **Contributor expectations -** If you are looking for contributors, make sure you set the **ground rules**. There's nothing worse than getting someone who **wants to help you but they don't know how**! This section of the README gives the guidelines for contributions. Do you expect someone to create an issue in the issue queue and then resolve it with a pull request? Do you want squashed commits? Do you have a pull requests template? **Explain it all** here. 
7. **Known issues -** I already talked about this README section above so I'll keep it short. Make a brief list of known issues here so people don't report bugs you already know about!
8. **Begging for money! -** Don't be ashamed to ask for money. Seriously! You put a lot of effort into this project and if someone likes it they might just **throw you a couple of bucks**. You can add a link to [Buy Me a Coffee!](https://www.buymeacoffee.com/askcloudtech)

## TLDR

Too many words? Here's a README outline

* Intro
* about paragraph(s)
* diagram
* installation instructions for users
* installation instructions for contributors
* contributor expectations
* known issues
* beg for money 

Oh and one last thing. Just do yourself a favor and keep your README up to date on EVERY commit!