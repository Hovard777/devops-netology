#6.5. Elasticsearch
FROM centos:7
LABEL ElasticSearch Lab
MAINTAINER Ivan Febres <hovard777@gmail.com>
ENV PATH=/usr/lib:/usr/lib/jvm/jre-11/bin:$PATH

RUN yum install java-11-openjdk -y 
RUN yum install wget -y 

RUN wget https://fossies.org/linux/www/elasticsearch-8.6.0-linux-x86_64.tar.gz 
RUN yum install perl-Digest-SHA -y 
RUN tar -xzf elasticsearch-8.6.0-linux-x86_64.tar.gz \
    && yum upgrade -y
    
ADD elasticsearch.yml /elasticsearch-8.6.0/config/
ENV JAVA_HOME=/elasticsearch-8.6.0/jdk/
ENV ES_HOME=/elasticsearch-8.6.0
RUN groupadd elasticsearch \
    && useradd -g elasticsearch elasticsearch
    
RUN mkdir /var/lib/logs \
    && chown elasticsearch:elasticsearch /var/lib/logs \
    && mkdir /var/lib/data \
    && chown elasticsearch:elasticsearch /var/lib/data \
    && chown -R elasticsearch:elasticsearch /elasticsearch-8.6.0/
RUN mkdir /elasticsearch-8.6.0/snapshots &&\
    chown elasticsearch:elasticsearch /elasticsearch-8.6.0/snapshots
    
USER elasticsearch
CMD ["/usr/sbin/init"]
CMD ["/elasticsearch-8.6.0/bin/elasticsearch"]
