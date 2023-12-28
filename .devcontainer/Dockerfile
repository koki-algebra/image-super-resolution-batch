FROM mcr.microsoft.com/devcontainers/base:jammy

USER vscode

# install rye
RUN curl -sSf https://rye-up.com/get | RYE_INSTALL_OPTION="--yes" bash \
	&& echo 'source "$HOME/.rye/env"' >> ~/.bashrc
