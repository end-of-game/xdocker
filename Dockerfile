FROM debian:jessie
ENV VERSION 2.25
RUN apt-get update -q
RUN apt-get install -qy curl build-essential
ADD xdocker /xdocker
ADD installer /installer
CMD /installer
