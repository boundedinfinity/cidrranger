FROM boundedinfinity/go-glide:1.0.0
MAINTAINER brad.babb@boundedinfinity.com

ENV APP_DIR=/app
ENV DIST_DIR=/dist
ENV GOPATH=$APP_DIR

RUN mkdir -p $APP_DIR && mkdir -p $DIST_DIR

COPY . $APP_DIR

RUN cd $APP_DIR && \
    make go-bootstrap && \
    make go-install && \
    make beego-package docker_dist_dir=$DIST_DIR && \
    cd $DIST_DIR && \
    tar zxvf echo.tar.gz && \
    rm -rf echo.tar.gz

EXPOSE 8080

WORKDIR $DIST_DIR
CMD ["./echo"]
