FROM busybox

RUN cd /
RUN apt-get update
RUN apt-install tree
RUN echo 'hoge'
