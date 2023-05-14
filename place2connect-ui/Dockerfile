# This tells us the version of node image we want to use.

# FROM node:alpine

FROM node:20-alpine3.16

# This tells Docker to create a directory so that other commands will use it
WORKDIR /app

#Copy Our package.json and package-lock.json file into the app directory to tell node the module we want to use
COPY /place2connect-ui/package.json /app
COPY .env  /app
COPY .env .

COPY   start.sh .
COPY   wait-for.sh .
# COPY   /place2connect-ui/start.sh .
# COPY   /place2connect-ui/wait-for.sh .
COPY   /place2connect-ui/ .
#To install the dependencies inside our image

RUN npm i
RUN npm run build

# Copy everything from ourlocal directory to the image in the code directory
COPY . /app
COPY .env  /app
COPY .env .

# Navigate to the application entry point and run the image
# EXPOSE 3000

CMD [ "npm", "run", "dev" ]
ENTRYPOINT ["/app/start.sh" ]

