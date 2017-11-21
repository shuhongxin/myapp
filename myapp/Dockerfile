# This dockerfile uses the centos image
# VERSION 1 - EDITION 1
# Author: shuhongxin
# Command format: Instruction [arguments / command] ..
 
# Base image to use, this must be set as the first line
FROM hub.c.163.com/library/centos

# Maintainer: docker_user <docker_user at email.com> (@docker_user)
MAINTAINER shuhongxin@ghostcloud.cn

COPY ./   /usr/local/myapp
WORKDOR   /usr/local/myapp


RUN echo 'export PATH="$PATH:/usr/local/myapp"' >> ~/.bashrc



