# set base image (host OS)
FROM python:3

# set the working directory in the container
WORKDIR /go/src/scripts

# copy the dependencies file to the working directory
COPY requirements.txt .

# install dependencies
RUN pip install -r requirements.txt

# copy the content of the local directory to the working directory
COPY . .
# COPY ./../db/data ./data

# command to run on container start
CMD [ "python", "builddb.py" ]
