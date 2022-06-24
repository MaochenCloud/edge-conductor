#!/bin/bash

# Copyright (c) 2022 Intel Corporation.
#
# SPDX-License-Identifier: Apache-2.0

# set -ex

# Output locatoin
# ${WEB_FILES}/${profile_name}/build/
# Example
# esp/data/usr/share/nginx/html/files/Ubuntu_20.04/build/
MIRROR_DIR=/opt/output
UBUNTU_VERSION=focal
UBUNTU_MIRROR_URL=http://archive.ubuntu.com/ubuntu

# Output file
# ubuntu_rootfs.tgz

# This script will build a ubuntu roofs tarball ubuntu_rootfs.img
# pre-request
#   minial ubuntu os at least

# step 0 : prepare
apt update -y
apt-get install -y debootstrap

# step 1 : create rootfs image and mount on fs 2G
umount rootfs/dev || true
umount rootfs/proc || true
umount rootfs/sys || true

# step 2 : mirror and make /tmpfile/ubuntu as a rootfs
## warning !
## /etc/sudoer
## Defaults env_keep += "http_proxy https_proxy ftp_proxy"
mkdir -p rootfs
debootstrap --arch amd64 --variant=minbase $UBUNTU_VERSION rootfs $UBUNTU_MIRROR_URL

# step 3 : bind system fs
mount --bind /dev rootfs/dev
mount -t proc proc rootfs/proc
mount -t sysfs sysfs rootfs/sys

# step 4 : !! Can not be updated !! update package and install extra packages
#cp /etc/apt/apt.conf  /tmpfile/ubuntu/etc/apt/apt.conf
#LANG=C.UTF-8 chroot /tmpfile/ubuntu/ sh -c \
#        "apt update -y && \
#         apt install -y wget openssh-server && \
#         apt clean"

rm rootfs/etc/apt/sources.list
echo "deb http://archive.ubuntu.com/ubuntu focal main restricted" >> rootfs/etc/apt/sources.list
echo "deb http://archive.ubuntu.com/ubuntu focal-updates main restricted" >> rootfs/etc/apt/sources.list
echo "deb http://archive.ubuntu.com/ubuntu focal universe" >> rootfs/etc/apt/sources.list
echo "deb http://archive.ubuntu.com/ubuntu focal-updates universe" >> rootfs/etc/apt/sources.list
echo "deb http://archive.ubuntu.com/ubuntu focal multiverse" >> rootfs/etc/apt/sources.list
echo "deb http://archive.ubuntu.com/ubuntu focal-updates multiverse" >> rootfs/etc/apt/sources.list
echo "deb http://archive.ubuntu.com/ubuntu focal-security main restricted" >> rootfs/etc/apt/sources.list
echo "deb http://archive.ubuntu.com/ubuntu focal-security universe" >> rootfs/etc/apt/sources.list
echo "deb http://archive.ubuntu.com/ubuntu focal-security multiverse" >> rootfs/etc/apt/sources.list


# step 5: chanfe root and install packages
LANG=C.UTF-8 chroot rootfs sh -c "
    export TERM=xterm-color && \
    export DEBIAN_FRONTEND=noninteractive && \
    export https_proxy=$https_proxy && \
    export no_proxy=$no_proxy && \
    export HTTPS_PROXY=$HTTPS_PROXY && \
    export NO_PROXY=$NO_PROXY && \
    apt update && \
    apt install -y sudo wget vim && \
    apt --download-only --assume-yes install systemd && \
    apt --download-only --assume-yes install locales && \
    apt --download-only --assume-yes install debconf && \
    apt --download-only --assume-yes install grub-efi shim && \
    apt --download-only --assume-yes install grub-pc && \
    apt --download-only --assume-yes install linux-image-generic && \
    apt --download-only --assume-yes install docker.io && \
    apt --download-only --assume-yes install pciutils && \
    apt --download-only --assume-yes install wget && \
    apt --download-only --assume-yes install openssh-server && \
    apt --download-only --assume-yes install socat  && \
    apt --download-only --assume-yes install ebtables  && \
    apt --download-only --assume-yes install ethtool  && \
    apt --download-only --assume-yes install conntrack  && \
    apt --download-only --assume-yes install ufw  && \
    apt --download-only --assume-yes install cloud-init  && \
    apt --download-only --assume-yes install pciutils  && \
    apt --download-only --assume-yes install net-tools  && \
    apt --download-only --assume-yes install nano  && \
    apt --download-only --assume-yes install init && \
    apt --download-only --assume-yes install vim && \
    apt --download-only --assume-yes install qemu-guest-agent && \
    apt --download-only --assume-yes install chrony && \
    apt --download-only --assume-yes install dkms && \
    apt --download-only --assume-yes install libgstreamer1.0-0 && \
    apt --download-only --assume-yes install gstreamer1.0-tools && \
    apt --download-only --assume-yes install gstreamer1.0-plugins-base && \
    apt --download-only --assume-yes install gstreamer1.0-plugins-good && \
    apt --download-only --assume-yes install gstreamer1.0-libav && \
    echo \"chroot install done !!!\""


# step 7 : unmount and put rootfs.tgz under tftp output path
umount rootfs/dev
umount rootfs/proc
umount rootfs/sys

rm -f rootfs.tgz
pushd rootfs
tar czvf ../rootfs.tgz *
popd
mv rootfs.tgz  $MIRROR_DIR

