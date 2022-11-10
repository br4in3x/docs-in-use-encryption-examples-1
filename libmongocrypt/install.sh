#!/bin/sh
export LIBMONGOCRYPT_VERSION=1.6
export LINUX_RELEASE=bionic

sudo sh -c 'curl -s --location https://www.mongodb.org/static/pgp/libmongocrypt.asc | gpg --dearmor >/etc/apt/trusted.gpg.d/libmongocrypt.gpg'
echo "deb https://libmongocrypt.s3.amazonaws.com/apt/ubuntu ${LINUX_RELEASE}/libmongocrypt/${LIBMONGOCRYPT_VERSION} universe" | sudo tee /etc/apt/sources.list.d/libmongocrypt.list

apt-get update
apt-get install -y libmongocrypt-dev
which libmongocrypt