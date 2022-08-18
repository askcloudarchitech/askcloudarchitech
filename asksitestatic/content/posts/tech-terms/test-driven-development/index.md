---
title: What Is Test Driven Development (TDD)?
date: 2022-08-18T04:26:38.999Z
description: There are a number of benefits to TDD all related to the idea that
  the tests are written ahead of time. When you practice TDD as a development
  methodology, it forces you to focus on individual features until they are
  complete (because your tests describe very specific scenarios and will
  continue to fail until you solve them). It can also make your code easier to
  refactor and maintain because it is written in a modular way where each
  feature has code and a related set of tests indicating if it still provides
  the intended functionality.
draft: false
keywords:
  - what is test driven development
  - tdd
  - test driven development tutorial
  - test driven development
  - software development
  - test driven development worth
  - test driven development explained
  - test-driven development
  - tech terms you should know
  - tech terms for dummies
  - tech terms explained
  - technical terms explained
tags:
  - what is test driven development
  - tdd
  - test driven development tutorial
  - test driven development
  - software development
  - test driven development worth
  - test driven development explained
  - test-driven development
  - tech terms you should know
  - tech terms for dummies
  - tech terms explained
  - technical terms explained
ID: 1660796798988
mainImage: TDD
videoID: pJTX-NcY95E
resources:
  - name: TDD
    title: TDD
    src: thumb.001.jpeg
---
TDD stands for test driven development and at the surface its a pretty simple concept. Essentially, when practicing test driven development you write failing tests first then write your code to solve the problem which makes the test pass.

{{< youtube pJTX-NcY95E >}}

## **Upsides of TDD**

There are a number of benefits to TDD all related to the idea that the tests are written ahead of time. When you practice TDD as a development methodology, it forces you to focus on individual features until they are complete (because your tests describe very specific scenarios and will continue to fail until you solve them). It can also make your code easier to refactor and maintain because it is written in a modular way where each feature has code and a related set of tests indicating if it still provides the intended functionality.

TDD also naturally produces well-documented features because each feature of the codebase has tests written with well explained scenarios.

Proponents of the TDD methodology argue that it reduces the cost of a project over time because it reduces the number of defects introduced into the project as time goes on, again because there is naturally a large set of test already in place which cover all of the applications specific logic and business use cases.

## **Downsides of TDD**

Now while all of those upsides are true, there is definitely another side to this coin.

Like I mentioned above, while TDD can reduce the cost of a project’s maintenance over time, it most certainly slows down a project in the beginning. If you are starting from scratch on a project it can be very time consuming to write failing tests for all of the basic boilerplate scenarios and will definitely slow things down.

But… Practicing the TDD methodology from the very beginning of a project is essentially required if you really want to truly follow the methodology. You can start with an already written codebase, but this would only provide tests and documentation for bugs and enhancements which would definitely reduce the overall effectiveness of the entire process.

So as you can see, this a bit of a catch 22. you are either starting from the very beginning of the development process with a TDD approach where it takes longer but you get long term benefits, or it may not be actually worth doing at all.

## The TDD mindset, VS methodology

Now, everything I’ve talked about until now is really describing the test driven development methodology which is definitely something you or an entire team would have to commit to for a very long time, but that doesn’t mean you cant use the idea of TDD in your day to day code, especially when fixing bugs.

I personally really like the idea of writing failing tests that describe a defect in a application then writing the code that fixes the defect. This is not really going all in with the whole TDD methodology, but it does ensure that you have truly fixed an issue as described and will ensure that this specific issue will never happen again because there is a test for it.

Final thoughts

Finally, if you think that the entire TDD methodology makes good sense, be prepared to have to sell the idea to stakeholders in the beginning of a project. Setting proper timeline expectations is absolutely key for a successful TDD based approach. Stakeholders who don’t truly understand the long term benefits may not support the process because of the time it takes in the beginning.