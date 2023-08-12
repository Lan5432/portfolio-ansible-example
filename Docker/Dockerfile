FROM ubuntu:22.04
ARG USER=test

RUN apt update && apt install openssh-server sudo -y
RUN useradd -rm -m -s /bin/bash -u 1000 -G sudo -U $USER
RUN echo "$USER:$USER" | chpasswd

RUN mkdir /home/$USER/.ssh
COPY authorized_keys /home/$USER/.ssh/authorized_keys
RUN chmod 700 -R /home/$USER/.ssh
RUN chown -R $USER:$USER /home/$USER/.ssh
RUN usermod -aG sudo $USER
RUN service ssh start

EXPOSE 22

ENTRYPOINT ["/usr/sbin/sshd"]
CMD [ "-D" ]