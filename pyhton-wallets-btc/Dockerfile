FROM ubuntu:latest

# Update package manager (apt-get) 
# Python and Pip
RUN apt-get update
RUN apt-get install -y python3.9 
RUN apt-get install -y python3-pip

WORKDIR /wallets


# Copy our script into the container
COPY . .


# Install our Python dependencies
RUN pip install -r requirements.txt

EXPOSE 5000