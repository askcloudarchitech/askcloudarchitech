---
title: How To Turn Your Old Hardware Into A Kubernetes Cluster
date: 2023-03-11T18:31:42.966Z
description: Turning your old hardware into a kubernetes cluster is super easy. Learn how
  to use microk8s and revive your old hardware into a multi node kubernetes
  cluster.
draft: false
videoID: FD6x3vb3Fk4
keywords:
  - cluster
  - hardware
  - kubernetes
  - microk8s
  - tutorial
  - how-to
tags:
  - cloud architecture
  - kubernetes
  - microk8s
ID: 2023-03-11T18:31:42.966Z
previewImage: DSC00029.JPG
mainImage: ""
---

Today I’m going to show you how to set up Kubernetes on an old computer that you may have sitting around. 

{{< youtube FD6x3vb3Fk4 >}}

While setting up something like Kubernetes may seem complicated, It’s a surprisingly easy process once you have the steps and the right software to get it done fast. I think it’s the best choice when it comes to setting up a home server since you can add and remove applications without worrying about messing up the underlying operating system of the machine.

So let's jump in roll through the steps. 

No matter what kind of machine you are using, you’ll need some basics to get started so make sure you have the following handy.

1. The computer itself. I'm using an older Mac mini, but any computer will work.
2.  A basic USB flash drive. You will need this to house the software we are going to install.
3.  A network cable. Fiddling with Wi-Fi setup is a pain, so it's best to just wire the computer to the network at least initially.
4.  A mouse and keyboard that is either wired or has a dedicated dongle. Bluetooth probably won't work initially, so this is a must-have.

{{< img2 src="microk8s-install_1.6.1.jpg" alt="Hardware needed for setup" >}}

Once you have all your supplies together, grab the USB flash drive and plug it into a different, working computer. Not the one we are about to erase.  As you can see here, I’m using a very small 4 GB flash drive. It doesn't take much.

{{< img2 src="microk8s-install_2.12.1.jpg" alt="4GB flash drive" >}}

Now on your working computer, we will need to download a few things. First Is a copy of Ubuntu server. Head over to the [Ubuntu website](https://ubuntu.com/download/server) and download the latest version of Ubuntu server. 

After that’s finished, we will also need a program for flashing the Ubuntu image onto the flash drive. I used a program called etcher, and it's available for both Mac and Windows. Head over to [their website](https://www.balena.io/etcher) and download the program and install it.  

After downloading those two things, open up Etcher and choose the “flash from file” option. The file picker will open, and you should choose the Ubuntu ISO file that you downloaded earlier. 

In the next column set the target as your flash drive then hit the “flash button”. This will turn your flash drive into a bootable Ubuntu installer. After its finished, eject the flash drive and unplug it from your working computer. 

{{< img2 src="microk8s-install_2.18.1.jpg" alt="Balena Etcher" >}}

That's all we need to do for now on your working computer. Now let's shift focus over to the old machine that we’re going to install Kubernetes on.

Temporarily we are going to need to wire up this computer with a mouse, keyboard and monitor. Also plug in the network cable and your USB flash drive.

{{< img2 src="microk8s-install_2.23.1.jpg" alt="Connected mac mini for setup. keyboard dongle on left-most USB port" >}}

One quick note if you are doing this setup on a Mac mini: make sure to plug the keyboard into the left most USB port. Otherwise, the mac won't recognize the keyboard during its boot sequence. Ask me how I know that!

Once you get everything plugged in you need to get the machine to boot from the flash disk. The process to make this happen varies by machine. Some will automatically detect the boot disk and prompt you. Some will require that hit an F key at startup. In the case of this Mac mini, I had to hold the Alt key on my keyboard to get the boot menu. 

{{< img2 src="microk8s-install_4.1.1.jpg" alt="Mac Mini boot menu" >}}

No matter which one you have to do, just make sure you tell the machine to boot from the flash drive. Once you do that you should see this simple prompt asking you what you want to do. Choose the option to install Ubuntu server and hit enter.

{{< img2 src="microk8s-install_2.25.1.jpg" alt="install ubuntu server screen" >}}

At this point it spits out a bunch of text as it first starts. This process takes a few minutes. After all that’s done, you’ll get the Ubuntu server welcome screen, and you’ll need to provide some details.

Fair warning, I’m about to roll through 15 short, but required steps for this setup. I’m going to do this as quickly as possible without leaving out the few important bits scattered through this otherwise pretty self-explanatory wizard. Eyes on the prize, we’ll get through it. 

First, choose your language.

{{< img2 src="microk8s-install_4.2.1.jpg" alt="Choose language" >}}

Next, it may ask you if you want to upgrade to the latest release. I skipped this for now. You can always do it later.

{{< img2 src="microk8s-install_4.2.2.jpg" alt="option to update to latest version" >}}

Next, configure your keyboard.

{{< img2 src="microk8s-install_4.3.1.jpg" alt="Keyboard setup" >}}

Next it will ask what type of server install you want. Choose the minimized version unless you have a need to use this machine for other stuff.

{{< img2 src="microk8s-install_4.4.2.jpg" alt="Choose ubuntu server (minimized)" >}}

If all goes well it should automatically connect to the internet, and you can just hit enter. Before leaving this screen though, write down the IP address it was assigned. You’ll need this later. 

{{< img2 src="microk8s-install_4.5.2.jpg" alt="Auto detect IP address" >}}

Next leave the proxy empty and hit enter.

{{< img2 src="microk8s-install_4.6.1.jpg" alt="Leave proxy address empty" >}}

Next leave the mirror address alone and hit enter.

{{< img2 src="microk8s-install_4.6.2.jpg" alt="Leave mirror address as-is" >}}

Next on the disk setup, I chose to use the entire disk and set the disk up as an LVM group. These are the defaults and they are fine.

{{< img2 src="microk8s-install_4.7.1.jpg" alt="Allow to use the entire disk" >}}

Finally, it will show a summary. Hit enter to confirm then hit enter one more time to extra confirm that you are cool with it blowing away your hard drive. 

{{< img2 src="microk8s-install_4.8.2.jpg" alt="Confirm delete hard drive" >}}

On the next screen set up your user info. Keep a note of your username and password. You might need it later. 

{{< img2 src="microk8s-install_4.9.1.jpg" alt="Set up server and user information" >}}

Next skip the Ubuntu pro ad, unless you need enterprise support.

Next, enable the option to install OpenSSH server. This will allow you to access the machine remotely with no keyboard, mouse or monitor attached. Also on this screen, if you have a GitHub account with your public keys already stored, provide your GitHub username and the installer will pull them in as authorized keys for password-less access via SSH. 

If you are at all concerned about this not working, also check the box to allow password auth over SSH. This will allow you to log in with the password you just set. 

It will now pull your public key from GitHub and ask you to confirm.

{{< img2 src="microk8s-install_4.11.1.jpg" alt="Supply github user for password-less auth" >}}

Now, for the actual installation of Kubernetes. Its literally one check box. Select the option to install Microk8s. This is Ubuntu’s super popular lightweight version of Kubernetes, and it’s what we will be installing. After checking the box, go to the next step.

{{< img2 src="microk8s-install_4.12.2.jpg" alt="Choose Microk8s from the menu" >}}

Now we wait for the installation to finish…

Don't be like me and completely miss the text menu at the bottom of the screen. When it's done, that menu will show an option to “reboot now”.

{{< img2 src="microk8s-install_2.46.1.jpg" alt="Reboot now after install" >}}

Note that if you Don't take the flash drive out at this point it will probably yell at you. Take out the flash drive then reboot. 

{{< img2 src="microk8s-install_2.48.1.jpg" alt="Unplug the flash drive" >}}

And just in case you thought you were done waiting, you’re wrong. Now it will install microk8s. After its done it will show the login prompt. 

Ok, now it's done installing. You can shut the machine down and unplug all the unnecessary cables at this point and put your server in its final resting place. Then power it back on. 

Ok, now we have a running Kubernetes server and there’s only a couple more things to do before you can call it done. 

Now back at your working machine, open up a terminal and attempt to SSH into the machine. Your ssh command should be.

`ssh YOURUSERNAME@YOURSERVER_IP`

If you chose earlier to import your public keys from GitHub it should just let you without a password. If it didn't work, or you didn't choose that option, you will need to enter your password at this point. 

{{< img2 src="microk8s-install_2.55.1.jpg" alt="SSH into server" >}}

Now that we’ve confirmed we can connect to the server with SSH, you want to make sure that the IP address doesn’t change on you. Luckily most home routers these days have the ability to reserve an IP address and associate it with a specific mac address. On the screen is how I did this on my Netgear router, but every one is different. The important takeaway here is that you make sure the IP address of your server won't change when you lose power or restart your router. This will be important when you begin to use the cluster and need to reach it by IP address.

{{< img2 src="microk8s-install_2.57.1.jpg" alt="Pin the IP address of the server in the netgear manager" >}}

At this point you have a working single node Kubernetes cluster, but don't bail on my yet. Now I'm going to show you how you can expand the capacity of your cluster by adding more machines. If you have two or three or four old machines lying around they could all work together as a single Kubernetes cluster. 

And the best part is that it only take a couple commands to make that happen. So even if you don't have another machine to add now, don't you want to hang out just to see how it's done?

So obviously if you have another machine to add, you will need to do all the exact same steps to set up microk8s as we already did on the first one. The only difference will be the server name and the IP address. Everything else is exactly the same. 

Once you have both machines running as standalone clusters, SSH into both of them.

Below you are seeing the first machine’s terminal on the left, the second machine’s terminal on the right and finally my k9s Kubernetes dashboard at the bottom connected to the first machine. You will note that it only shows one node currently. 

{{< img2 src="microk8s-install_2.58.1.jpg" alt="update etc hosts with node 2 ip and name" >}}

1. With vim or another text editor, open the /etc/hosts file on each machine.
2. On both machines, enter the IP address and the name of the second Kubernetes server. This is required so microk8s can complete the join.
3. save the /etc/hosts file on each machine and exit vim
4. now, on the first machine enter microk8s add -node and hit enter. It will give you the command you need to enter on the second machine. 
5. Copy the command and paste it into the terminal of the second machine. 
6. The system will now join the two machines. You can see in the bottom k9s window, the cluster now shows two nodes. 

{{< img2 src="microk8s-install_2.60.1.jpg" alt="run microk8s add-node" >}}

And for demonstration purposes, I'm going to scale my installation of home assistant to two replicas. You can see that it is now running one replica on each node with no issues at all. 

{{< img2 src="microk8s-install_2.67.1.jpg" alt="Successful add of second node" >}}

Thanks for stopping by and until next time, happy coding.
