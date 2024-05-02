# Use an official Python runtime as a parent image
FROM python:3.9

# Install cron
RUN apt-get update && apt-get install -y cron

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the requirements.txt first to leverage Docker cache
COPY requirements.txt ./

# Install any needed packages specified in requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Copy the Gpt4all directory to the container
COPY Gpt4all/ /usr/src/app/Gpt4all

# Copy the Code directory to the container
COPY Code/ /usr/src/app/Code
COPY Pdf/ /usr/src/app/Pdf

# Copy cronjob file to the container
COPY crontab.txt /usr/src/app/

# Copy the rebuild_chatbot_container.sh script to the container
COPY rebuild_chatbot_container.sh /usr/src/app/

# Set execute permission to the script
RUN chmod +x /usr/src/app/rebuild_chatbot_container.sh

# Set up the cron job
RUN crontab /usr/src/app/crontab.txt

# Run the cron daemon in the background and the Python script in the foreground
CMD cron && python Code/main.py