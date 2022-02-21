---
title: "Create a Custom CLI Tool and Distribute with HomeBrew Using Goreleaser and Github Actions"
slug: "create-homebrew-tap-golang-goreleaser-cobra-cli"
date: 2022-02-20T19:56:10-05:00
ID: 62283b7b1af6b2d659f02ec4d11f0dd7
draft: false
description: Learn how to create a command line tool with Go Cobra and distribute it to homebrew using goreleaser and github actions. 
keywords:
    - Go Cobra
    - Homebrew
    - homebrew Tap
    - Goreleaser
    - Github Actions
    - Custom CLI tools
    - golang development
categories: 
    - Development
tags:
    - Go Cobra
    - Homebrew
    - Homebrew Tap
    - Goreleaser
    - Github Actions
    - Custom CLI tools
    - golang development
resources:
    - name: create-secret
      src: "create-secret.png"
      title: Create a new secret
    - name: file-structure
      src: "file-structure.png"
      title: Go Cobra CLI file structure
    - name: new-personal-access-token
      src: "new-personal-access-token.png"
      title: Github personal access token screen
    - name: new-secret
      src: "new-secret.png"
      title: Github secrets screen
    - name: success-check
      src: "success-check.png"
      title: Green check showing success
    - name: actions-output
      src: "actions-output.png"
      title: Github actions screen
    - name: homebrew-tap-repo
      src: "homebrew-tap-repo.png"
      title: Homebrew Tap repo
---

Previously, [I wrote about installing and using Homebrew](https://medium.com/@askcloudarchitech/everything-you-need-to-know-about-homebrew-e91d1e82959a). In that article, I briefly mentioned that you could pretty easily create and distribute your own apps using HomeBrew. Today, I'm going to cover how to do that step-by-step. In this article I will cover the basics of creating a command line app using the excellent Go package Cobra then show you how to easily and automatically publish your tool to a HomeBrew Tap so others can install it with a couple simple commands.  Let's get started.

## Part 1. Create a CLI tool using Go Cobra
Since the primary purpose of this article is to show how to distribute a CLI tool with Homebrew, we obviously need to have a CLI tool to distribute! To accomplish this, I will be using a CLI tool that I have already written. I am not going to go really in depth into the functionality of the tool I created and will instead cover that in another post later on. The important part that we need here is just the basics of creating a CLI tool and the file layout. Obviously you will be creating your own tool so the details of mine don't really matter. 

I have written a tool called mediumautopost that is designed to post articles to medium.com. This tool is written in the Go programming language and is using the go-cobra package. According to the go-cobra github page "Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files". 

https://github.com/spf13/cobra

To get started creating a new CLI tool with go-cobra, you should first install the cobra tool. BTW, go-cobra is written in Go and we will be writing this CLI in Go, so install that too if you don't have it.

`brew install go`  
`go get -u github.com/spf13/cobra`

Now that it's installed. Create a new directory for your CLI tool's git repo. You can really put this anywhere. 

`mkdir myCLItool && cd myCLItool`

Then initialize it as a Git repo

`git init`

At this point you should have an empty directory with git initialized. Now it's time to create your github repo to match this project. Creating a github repo is outside of the scope of this project, but head over to github.com and create your project (likely with the same name as the folder you created above). You can push up your initial commit if you'd like at this point but it's not required.

Next, initialize your project as a go module

`go mod init github.com/YOURGITHUBNAME/YOURMODULENAME` - again, replace YOURMODULENAME with the name of the github repo you created. 

The next step is to use the `cobra` command line tool to build the scaffolding of your new CLI tool. It's super simple and all you need to do is run this command: 

`cobra init`

In your project, you should have a set of files and a directory that look something like this:

{{<img name="file-structure" size="tiny">}}

From this point I'm going to start sharing code from the repo I already created. You can see the full source code here: https://github.com/askcloudarchitech/mediumautopost. I used the exact same steps I listed above when creating this project and I will be sharing the contents of the files below showing the important parts of the CLI. 

First is the `main.go` file.  

{{< gist gmorse81 b3021d5eb9e3e15d683e06fa298942bb >}}

Not much to it, right? That's on purpose. When creating a CLI using cobra, the main.go file simply imports the `cmd` package. The `func main()` just calls the `cmd.Execute()` function in the `cmd/root.go` file.  This file contains the meat of the CLI functionality, so let's take a look at that.

{{< gist gmorse81 5a3e13d6b8b84a300b32bb9435537925 >}}

Let's review the above file in chunks by the line numbers. 

Line 1: Just the package declaration. This is usually the same name as the folder containing the files. 

Lines 3-8: Importing all the necessary packages for this to work. `os` is used for returning exit codes when the command completes. `mediumautopost` is the actual logic of what the command does when you run it. Its best practice to separate the command details from the logic of the command. This allows you to later use the logic in places other than this command. This follows the coding principle of "don't repeat yourself". The final import is the `cobra` package itself which makes all this work. 

Line 10: Declaring a variable which will store the path to the .env file. This CLI uses environment variables for configuration and you can supply a file which contains your config. Below we will read in a flag that populates this variable. 

Lines 13-26: These are the instructions for the CLI command itself. When this is finished, running `mediumautopost` or `mediumautopost -h` (for help listing) will follow the instructions you put into these lines. As you can see its pretty straight forward. There is the name of the command, the short description, the long description and the `RunE` property which is function that gets executed. The function declared for `RunE` is kept pretty simple. It invokes a function named `Do` within the `mediumautopost` package. Again, I'm not going to go into the details of my specific CLI tool but just know that all the rest of what this thing actually does it kicked off by that call to `Do`. If you want to just get something working, you could easily replace line 23 with `fmt.PrintLn("Hello World!")` and then the command would print "Hello World!" to your terminal. 

Lines 30-35: This is the `Execute()` function that is called by `main.go`. You can see that it invokes the rootCmd declared above and then just returns a exit code of `1` if something went wrong or by default an exit code of `0` if all goes well. 

Lines 37-39: For those of you unfamiliar with Golang, this is the `init` function and it is automatically invoked when this package is initialized. In this function, we are creating the flag which allows you to pass the location of your .env file. `rootCmd.Flags().StringVarP` is a function which allows you to declare a flag, its description (for help text), its long name (--envfilepath), its short name (-e), and its default value if the flag is not used. In this case just an empty string. 

So... All of that together in this file automatically builds the CLI and it's available commands. These are the different ways it can be used. 

`mediumautopost -h` Prints help text  
`mediumautopost` Runs the program with no .env location provided  
`mediumautopost -e .myenv` Runs the program with a .env file of .myenv  
`mediumautopost --envfilepath=.myenv` Same as above, but full length flag name

Here's what the help text looks like:

```bash
For details on how to set up your site to use this program please visit 
https://askcloudarchitech.com/posts/tutorials/auto-generate-post-payload-medium-com/
Ensure you have set up your env file as shown in the .env.example
Example command: mediumautopost --envfilepath=.env

Usage:
  mediumautopost [flags]

Flags:
  -e, --envfilepath string   Path to your environment file. if left empty, the program will only use system environment variables.
  -h, --help                 help for medium-auto-post
```

At this point you should be able to test your program. With your terminal, run `go run main.go -h` and it should print the help text of your CLI. While this is great (yay, working code!), it's not exactly what we want in the end. If you want to easily install this tool and distribute it, you actually need to compile it and put it somewhere publicly. The next steps will outline how to automatically build your tool, create well documented releases on github and finally make the tool available to install using HomeBrew. 

## Part 2. Create a Github Repository for your homebrew tap

This next part was slightly confusing to me when I was first trying to understand it, so I think it deserves some explanation before diving into the steps. To distribute your tool with homebrew, you need what's called a homebrew Tap. A Tap is a separate github repo which only contains instruction for HomeBrew (so it knows how to install your stuff). When you create a Tap, you don't need to make a new one for each and every tool you wish to distribute. You can simply make one Tap and then distribute as many tools as you'd like through that single Tap. So it's best practice to make a single github repo for all of your stuff and name it with your github username (or organization name). In my case, I've created this repo which will be my homebrew tap for anything I make now and into the future: https://github.com/askcloudarchitech/homebrew-askcloudarchitech.

If you clicked the link above, you will see that there are only two files in this Tap: A README.md file (obviously) and a `mediumautopost.rb` file. I created the readme manually and pushed it up to this repo. The `mediumautopost.rb` is automatically generated and the steps outlined below will explain how to make all of that magic happen. 

For now, just go to your github account and create a new repo. Make sure the name of your repo starts with `homebrew-` and then whatever you want your Tap to be named. If you want to read more about the naming conventions, the homebrew site explains it here: https://docs.brew.sh/How-to-Create-and-Maintain-a-Tap#creating-a-tap

## Part 3. Create a Github personal access token and set up as a Actions Secret

A quick sanity check. You should at this point have two github repositories: The first one containing the code for your CLI tool and the second one for housing your homebrew Tap. Got it? Good. Everything I will be explaining from now on will be done in the first repo (the one with your CLI code). You don't need to actually do anything with the Tap repo as that is all automated. 

Since the first repo will need to publish the installation instructions to the second repo, you are going to need a github personal token. Go to https://github.com/settings/tokens and hit the "generate new token" button. Provide a name in the note field, select an expiration date and check the box next to Repo then hit "Generate token" at the bottom. Copy the token and put it somewhere safe for now. NOTE: Keep this token a secret. 

{{<img name="new-personal-access-token" size="medium">}}

Now that you have the token and we are already on github, let's quickly set this token as a secret that we will use in a later step. Go ahead and navigate to the repo of your CLI tool. From there click on settings. On the settings page on the left hand side, click secrets -> actions. Next, click the "new repository secret" button.

{{<img name="new-secret" size="medium">}}

On the "New Secret" screen, add your personal access token. You can name it whatever you like, but if you are following along with the code examples, name it `PUBLISHER_TOKEN` since that will be the secret name we reference later.

{{<img name="create-secret" size="medium">}}

## Part 4. Create a .goreleaser.yml file

Now onto the actual automation! We will be using [goreleaser](https://goreleaser.com/) to do all the heavy lifting. Here's what it will do for us with one simple config file:

1. Compile your CLI tool into a binary for multiple operating systems and architectures
1. Publish the binary along with release notes to the github releases tab on your project
1. Create the Homebrew formula (installation instructions) for your project and publish it to your Homebrew Tap repository.

Here's my .goreleaser file. You can mostly copy and paste this. Just create a file named .goreleaser in the root of your project. 

{{< gist gmorse81 e63b4919c0c062b4286de3934fc0e984 >}}

The contents of this file is somewhat self-explanatory, but let's review the sections so they all make sense. 

* **builds** - The builds section of the file instructs goreleaser to compile your code to a binary of the name you provide. You can specify the OS and the architecture as well as any additional flags needed. For my setup, I am compiling this program to work on Mac and Linux in both ARM and AMD64 architecture. I have also chosen to vendor my dependencies so I pass the flag to use the vendored files when compiling. This is a topic in itself, but if you don't want to deal with it, just remove that flag. [View full build config options here](https://goreleaser.com/customization/build/?h=builds)
* **release** - I use almost all the defaults here except to set the `prerelease` property to auto. This tells goreleaser to not release tags which are formatted as prerelease versions. [View full release options here](https://goreleaser.com/customization/release/?h=release)
* **universal_binaries** - This feature tells goreleaser to build universal binaries for MacOS. The `replace` option instructs the tool to replace the individual ARM and AMD64 binaries with the universal one. 
* **brews** - Finally the brews section tells goreleaser to publish this project to a homebrew Tap. Replace my details with your Tap information as well as the details of your project. [View full brews options here](https://goreleaser.com/customization/homebrew/?h=brews). NOTE that I do not provide the token here. That will happen in the next step. 

## Part 5. Create your Github Actions workflow

The final file needed to complete this project is the github actions config file. Create a file named `release.yml` in the directory `.github/workflows`. 

Here's mine:

{{< gist gmorse81 7ad16674c5693740495e39ea6b23ea47 >}}

This file is practically a copy/paste from the goreleaser docs. Goreleaser has created an action specifically for github actions and you can see the full instructions here: https://goreleaser.com/ci/actions/?h=github+ac. I have made just a couple modifications to this for my needs. First, I modified the `on` section at the top of the file which tells github actions to activate this action when a tag is pushed to this project. Second, I modified the `version` value so it sets the version number to the tag I created. Finally, I changed the `GITHUB_TOKEN` environment variable value to `${{ secrets.PUBLISHER_TOKEN }}` which will use the secret we created in step 3.

## Part 6. Publish and try it out

WHEW! Thats the whole setup. The final step is to push up a change and watch it work. I will first commit a change then tag it and push the tag. Here are the commands and full output from the steps:

```bash
➜  mediumAutoPost git:(main) ✗ git add .
➜  mediumAutoPost git:(main) ✗ git commit -m "update readme"
[main 94a760e] update readme
 1 file changed, 1 insertion(+), 1 deletion(-)
➜  mediumAutoPost git:(main) git push
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 286 bytes | 286.00 KiB/s, done.
Total 3 (delta 2), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (2/2), completed with 2 local objects.
To github.com:askcloudarchitech/mediumautopost.git
   dcf2695..94a760e  main -> main
➜  mediumAutoPost git:(main) git tag -a v0.4 -m "README bump for demo"
➜  mediumAutoPost git:(main) git push origin v0.4                     
Enumerating objects: 1, done.
Counting objects: 100% (1/1), done.
Writing objects: 100% (1/1), 171 bytes | 171.00 KiB/s, done.
Total 1 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:askcloudarchitech/mediumautopost.git
 * [new tag]         v0.4 -> v0.4
```

To simplify, here's just the commands without the git output

```bash
git add .
git commit -m "update readme"
git push
git tag -a v0.4 -m "README bump for demo"
git push origin v0.4
```

After pushing the tag, head over to Github and see if it worked. You should see a few things that would indicate success. 

1. On your CLI repo's home screen you should see a green check next to the most recent commit  
{{<img name="success-check" size="medium">}}
2. If you click on that green check, you should see the details of your action indicating that all steps ran successfully.  
{{<img name="actions-output" size="medium">}}
3. If you go to your *other* repo (the HomeBrew Tap) you should now see a new commit with a .rb file containing the instructions necessary to install your CLI.  
{{<img name="homebrew-tap-repo" size="medium">}}
4. Finally (and most importantly) you should be able to actually install your tool with HomeBrew! The details of your Tap will be different, but this is what I ran to install my program:

```bash
brew tap askcloudarchitech/askcloudarchitech
brew install mediumautopost
```

After it's installed, try it out!

```bash
mediumautopost -h

For details on how to set up your site to use this program please visit 
https://askcloudarchitech.com/posts/tutorials/auto-generate-post-payload-medium-com/
Ensure you have set up your env file as shown in the .env.example
Example command: mediumautopost --envfilepath=.env

Usage:
  mediumautopost [flags]

Flags:
  -e, --envfilepath string   Path to your environment file. if left empty, the program will only use system environment variables.
  -h, --help                 help for mediumautopost
```

<hr/>

If you've made it this far you probably got something out of this post. If you did, I would really appreciate if you would subscribe on medium.com and youtube.

<https://medium.com/@askcloudarchitech>  
<https://www.youtube.com/channel/UClSv7tWDA4wkCTLhZl1YBlw>