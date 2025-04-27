FROM ollama/ollama:latest

WORKDIR /root

COPY ./run-ollama.sh .

RUN chmod +x run-ollama.sh \
  && ./run-ollama.sh

EXPOSE 11434
