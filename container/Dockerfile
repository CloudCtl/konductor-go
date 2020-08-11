FROM registry.access.redhat.com/ubi8/ubi-minimal
COPY ./bin/konductor /root/
WORKDIR /root
RUN ./konductor init -h
CMD ["bash"]
