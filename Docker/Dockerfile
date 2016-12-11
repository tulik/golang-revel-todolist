FROM buildpack-deps:wheezy-scm

# gcc for cgo
RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
	&& rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.6
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 5470eac05d273c74ff8bac7bef5bad0b5abbd1c4052efbdbc8db45332e836b0b

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
COPY go-wrapper /usr/local/bin/

ENV DEBIAN_FRONTEND noninteractive
ENV TERM xterm
# Install MongoDB
RUN mkdir -p /data/db
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927
RUN echo "deb http://repo.mongodb.org/apt/debian wheezy/mongodb-org/3.2 main" | tee /etc/apt/sources.list.d/mongodb-org-3.2.list
RUN apt-get update
RUN apt-get install -y --no-install-recommends --force-yes
RUN apt-get install dnsutils -y
RUN apt-get install wget build-essential python libkrb5-dev mongodb-org mongodb-org-server mongodb-org-shell mongodb-org-mongos mongodb-org-tools -y --force-yes
RUN echo "mongodb-org hold" | dpkg --set-selections && echo "mongodb-org-server hold" | dpkg --set-selections
RUN echo "mongodb-org-shell hold" | dpkg --set-selections
RUN echo "mongodb-org-mongos hold" | dpkg --set-selections
RUN echo "mongodb-org-tools hold" | dpkg --set-selections
# Install Node.js 4.x
ENV NODE_VERSION v4.4.4
RUN wget --no-check-certificate https://nodejs.org/dist/$NODE_VERSION/node-$NODE_VERSION-linux-x64.tar.gz && \
    tar -C /usr/local --strip-components 1 -xzf node-$NODE_VERSION-linux-x64.tar.gz && \
    rm node-$NODE_VERSION-linux-x64.tar.gz
# Standard setting (can be overwritten by -e while running)
ENV STORAGE_ENGINE wiredTiger
ENV JOURNALING yes
ENV NODE_ENV production
# Add dynamic configurator script
ADD mongodb-configurator /usr/local/mongodb-configurator
# Add run script for MongoDB
ADD run.sh /usr/local/bin/run.sh
RUN chmod +x /usr/local/bin/run.sh
WORKDIR /usr/local/mongodb-configurator
# Setup of the configurator
RUN chmod +x configurator.js && \
    npm install forever -g && \
    npm install
WORKDIR /
EXPOSE 3000
EXPOSE 27017
CMD ["/usr/local/bin/run.sh"]
#Revel golang framework
ENV GOPATH /go
# Install revel and the revel CLI.
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
ADD todolist /go/todolist
# Open up the port where the app is running.
EXPOSE 8080
WORKDIR $GOPATH
ENTRYPOINT service mongod start && /go/todolist/run.sh
