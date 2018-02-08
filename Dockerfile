FROM alpine:3.4

RUN echo "===> Installing sudo to emulate normal OS behavior..."  && \
    apk --update add sudo                                         && \
    \
    \
    echo "===> Adding Python runtime..."  && \
    apk --update add python py-pip openssl ca-certificates    && \
    apk --update add --virtual build-dependencies \
                python-dev libffi-dev openssl-dev build-base  && \
    pip install --upgrade pip cffi                            && \
    \
    \
    echo "===> Installing Ansible..."  && \
    pip install ansible                && \
    \
    \
    echo "===> Installing handy tools (not absolutely required)..."  && \
    apk --update add sshpass openssh-client rsync git  && \
    \
    \
    echo "===> Removing package list..."  && \
    apk del build-dependencies            && \
    rm -rf /var/cache/apk/*               && \
    \
    \
    echo "===> Adding hosts for convenience..."  && \
    mkdir -p /etc/ansible                        && \
    echo 'localhost' > /etc/ansible/hosts

RUN pip install requests

WORKDIR /opt/pansible

ADD build/linux_amd64/pansible .
ADD docker/ansible.cfg .
ADD docker/callback_plugins ./callback_plugins
ADD pansible.yml ./pansible.yml

# certificate to be generated
ADD server.crt .
ADD server.key .

ADD ssh_config /root/.ssh/config
ADD docker/ping.yml .
