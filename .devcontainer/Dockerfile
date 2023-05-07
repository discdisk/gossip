FROM docker.io/library/golang:1.20.4
RUN apt update && apt install -y openjdk-17-jdk graphviz gnuplot bzip2 zsh
RUN cd /tmp;wget https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2