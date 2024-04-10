#!/bin/bash

# Rebuild the chatbot service container
docker-compose up -d --no-deps --build chatbot
