FROM golang:1.14.2-buster
RUN set -eux; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
    # utilities for keeping Debian and OpenJDK CA certificates in sync
    ca-certificates p11-kit \
    ; \
    rm -rf /var/lib/apt/lists/*

# Default to UTF-8 file.encoding
ENV LANG C.UTF-8

ENV JAVA_HOME /usr/local/openjdk-11
ENV PATH $JAVA_HOME/bin:$PATH

# backwards compatibility shim
RUN { echo '#/bin/sh'; echo 'echo "$JAVA_HOME"'; } > /usr/local/bin/docker-java-home && chmod +x /usr/local/bin/docker-java-home && [ "$JAVA_HOME" = "$(docker-java-home)" ]

# https://adoptopenjdk.net/upstream.html
# >
# > What are these binaries?
# >
# > These binaries are built by Red Hat on their infrastructure on behalf of the OpenJDK jdk8u and jdk11u projects. The binaries are created from the unmodified source code at OpenJDK. Although no formal support agreement is provided, please report any bugs you may find to https://bugs.java.com/.
# >
ENV JAVA_VERSION 11.0.7

ENV PATH "$PATH:/opt/apktool"
ENV APKTOOL_VERSION 2.4.1

# apktool
RUN mkdir /opt/apktool && \
    wget https://bitbucket.org/iBotPeaches/apktool/downloads/apktool_${APKTOOL_VERSION}.jar -O /opt/apktool/apktool.jar && \
    wget https://raw.githubusercontent.com/iBotPeaches/Apktool/master/scripts/linux/apktool -O /opt/apktool/apktool && \
    chmod +x /opt/apktool/apktool

RUN mkdir /app
ADD .  /app/

WORKDIR /app
RUN go get -u github.com/ndelphit/apkurlgrep


#HOWTO
#docker build  . -t apkurlgrep
#docker run -it -v <the apk directory>:/work/ apkurlgrep apkurlgrep -a /work/com.flipkart.android.apk