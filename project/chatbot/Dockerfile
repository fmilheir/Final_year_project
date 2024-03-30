# Use an official Python runtime as a parent image
FROM python:3.9

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

#debug
RUN ls -l /usr/src/app/Gpt4all

# Command to run on container start
CMD ["python", "Code/main.py"]
