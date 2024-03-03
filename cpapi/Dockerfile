# For more information, please refer to https://aka.ms/vscode-docker-python
FROM ubuntu:latest

RUN apt update

RUN apt-get install unzip wget openjdk-11-jdk -y
RUN java --version
RUN apt install curl iproute2 vim -y

# # Creates a non-root user with an explicit UID and adds permission to access the /app folder
# # For more info, please refer to https://aka.ms/vscode-docker-python-configure-containers
ARG USERNAME=appuser
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && usermod -aG video $USERNAME \
    && apt-get update \
    && apt-get install -y sudo \
    && echo "$USERNAME ALL=(ALL:ALL) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

USER $USERNAME
RUN mkdir -p /home/$USERNAME/cpapi
WORKDIR /home/$USERNAME/cpapi

RUN wget https://download2.interactivebrokers.com/portal/clientportal.gw.zip
RUN /usr/bin/unzip clientportal.gw.zip

CMD ./bin/run.sh root/conf.yaml