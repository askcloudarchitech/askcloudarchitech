---
title: I Made a Local Kubernetes Startup Script So You Don't Have To
date: 2022-04-23T14:24:42.088Z
description: sdf
draft: false
keywords:
  - sfd
tags:
  - sf
ID: 1650723882016
mainImage: local-kubernetes-startup-script
resources:
  - src: screen-shot-2022-04-23-at-10.45.18-am.png
    name: local-kubernetes-startup-script
    title: lk8s up
---
If you are getting started with Kubernetes development, one of the first things you are going to need is a way to run Kubernetes on your computer. There are a bunch of different ways to startup Kubernetes locally, but almost all of them either have a bunch of complicated steps or require you to understand Kubernetes in the first place (which kind of defeats the purpose). So to solve that problem for you (and myself, honestly) I wrote a simple Kubernetes startup script that will get you a running local cluster in no time. 

## What you get with this script

This Kubernetes startup script will install a local cluster on your machine, set up a local container repository, and install the NGINX Kubernetes ingress controller. With these three things, you should be able to install anything you want on the cluster and see it work. 

## Prerequisites

Before running the script, here are a few things you will need and how to install them

1. Docker - Install docker desktop from https://www.docker.com/products/docker-desktop/
2. Homebrew - This script is written for Mac. Check out [this article](https://askcloudarchitech.com/posts/tutorials/homebrew-setup-and-usage/) for how to install homebrew

If you are using a mac, this is all you need before running my script. If you are running windows, you will need to install [Kubernetes Kind](https://kind.sigs.k8s.io/) in WSL2 

## Downloading the script

You can download the [local Kubernetes startup script here](https://gist.githubusercontent.com/gmorse81/49a4d9998831aa98d604e015a817a078/raw/57129b79d9fd9d53fc53b9916ac767299a6ee030/lk8s)

## How to use the local Kubernetes startup script

After you have downloaded the script, save it somewhere that you can access it easily, or put it in your `/usr/local/bin` directory and name the file `lk8s`

Next, you can run the startup by running `lk8s up` 

If you are successful, it should look like this:

{{< img name="local-kubernetes-startup-script" size="medium" >}}

Should now be able to run `kubectl get pods` and see that the cluster responds. 

### Local repository

This script installs the Kind local registry. You can read more about there [here](https://kind.sigs.k8s.io/docs/user/local-registry/)

### Ingress controller