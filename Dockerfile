FROM ubuntu:18.04

COPY bin/instacrawler /

EXPOSE 3000

RUN chmod +x instacrawler

CMD nohup ./instacrawler