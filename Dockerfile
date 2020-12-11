FROM ubuntu:18.04

COPY bin/instacrawler /

ENV PORT 3000
EXPOSE 3000

RUN chmod +x instacrawler

CMD nohup ./instacrawler
