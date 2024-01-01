FROM ubuntu:22.04

RUN apt-get update && apt-get install -y curl \
	&& curl -sSf https://rye-up.com/get | RYE_INSTALL_OPTION="--yes" bash

ENV PATH="$PATH:/root/.rye/shims"

WORKDIR /app

COPY . .

RUN rm -rf .venv && rye sync --no-dev

CMD [ "rye", "run", "main" ]
