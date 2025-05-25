#!/usr/bin/env bash

echo "Starting Ollama server..."
ollama serve &


echo "Waiting for Ollama server to be active..."
while [ "$(ollama list | grep 'NAME')" == "" ]; do
  sleep 1
done

# ollama pull deepseek-r1:1.5b
ollama create perfume-assistant -f ./Modelfile
