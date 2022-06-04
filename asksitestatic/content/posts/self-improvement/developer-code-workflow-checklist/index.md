---
title: Always Do These Eight Things Before Pushing a Commit
date: 2022-04-24T22:10:58.666Z
description: >-
  Submitting code for review can be a nerve-wracking experience, especially if
  you are early in your coding career or new to a larger company requiring a
  code review process. 


  Having a development career and working on side projects are two completely different things. If you are working on your own project it's easy enough to have an idea, slap some code together, and push it up. Nobody else cares and life is simple. When your job is coding, your code probably has a tangible impact on the bottom line of the company or the customer’s experience. This usually means that your code will be scrutinized much more thoroughly before it actually makes it to production.
draft: false
keywords:
  - web development
  - software engineering
  - software development
  - coding best practices
  - learn to code
  - beginner coding
  - pull request
  - pull request github
  - code review
  - code review best practices
  - code faster
  - efficient engineer
  - efficient developer
  - web development tutorial
  - learn web development
  - askcloudarchitech
  - dev setup
  - vscode
  - Write better code
  - Write clean code
  - Pass code review
  - code review checklist
tags:
  - web development
  - software engineering
  - software development
  - coding best practices
  - learn to code
  - beginner coding
  - pull request
  - pull request github
  - code review
  - code review best practices
  - code faster
  - efficient engineer
  - efficient developer
  - web development tutorial
  - learn web development
  - askcloudarchitech
  - dev setup
  - vscode
  - Write better code
  - Write clean code
  - Pass code review
  - code review checklist
ID: 1650838258655
mainImage: wait-pull-request-stop
resources:
  - src: thumb.001.jpeg
    name: wait-pull-request-stop
    title: Do these things before submitting your PR
---
Submitting code for review can be a nerve-wracking experience, especially if you are early in your coding career or new to a larger company requiring a code review process. 

Having a development career and working on side projects are two completely different things. If you are working on your own project it's easy enough to have an idea, slap some code together, and push it up. Nobody else cares and life is simple. When your job is coding, your code probably has a tangible impact on the bottom line of the company or the customer’s experience. This usually means that your code will be scrutinized much more thoroughly before it actually makes it to production. 

This is why it’s super important to do your own code review before pushing the big green button on GitHub. Before submitting your next code contribution, try these steps. They’ll not only make you look like a hero but also save precious time for both you and the reviewer which is excellent for your career.

Before you even start coding, read and understand the requirements of the task you are being asked to complete. If you don’t completely understand what you are supposed to be doing, take some time and ask questions. Taking 10 minutes to ask questions will save you hours of time on the back end trying to fix code you wrote based on assumptions. 

Alright, got the requirements all figured out? Next, make sure you have your development space set up for running the code and debugging. Nobody writes perfect code without trying it out constantly and iterating on changes. Figure out the most efficient way to set up your local environment so you can write code, try it, fail fast, and continue moving forward without wasting a bunch of time. 

As you write your code, use descriptive variable names and add comments where needed if the code is complex or may need explanation. Comments are useful for future devs working on your project but also for the reviewer so they can actually understand your approach while reviewing the code. 

You should also be writing tests for your code now as you write the code. here’s a good rule of thumb for writing tests. 

As you are writing, there will be natural spots where you will need to run the code to see if it works as expected. These spots are also the best test scenarios. Instead of dumping a debug statement to see if your code produces the correct results, write a unit test that passes if it returns the expected results. By the time you are done coding, you will already have a ton of tests that prove the function of the code. 

OK, so now you’re finished writing but before you submit your changes there are a few more steps before you actually submit your pull request.

First and most importantly, does your code actually solve the problem that was described in the request? This might seem like a dumb question, but it’s really easy to lose track of the specifics of the request and miss a finer detail. Literally, go back to the ticket and read it again. Does your code actually solve the problem?

Next, do your own code review before anyone else does. Check for the basic stuff. Did you leave debugging log statements in the code? Are your variable names descriptive? Do your function names make sense? Do you have all your tabs and spaces in order? 

Next, try to break your code. Don’t wait for QA to find that edge case. Save everyone some time and think critically about what could potentially break your implementation. Everyone’s code is different, but since you just wrote it you probably know best about what could potentially break it. Remember that the end-user won’t always use your tools exactly as intended so throw some random stuff at it and make sure it still works. 

While you are at it, those types of edge cases also make great automated test cases. you might want to write those as tests as well. 

Finally, does the codebase you’re working already have an automated test suite? If so, run the tests and make sure everything is passing. 

Ok, now it’s finally time. Push up your changes, submit your PR, and feel confident that you have done everything within your power to submit the best code you can possibly produce. If all goes well, the only feedback you’ll receive is “LGTM”!