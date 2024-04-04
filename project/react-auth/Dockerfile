# Stage 1: Build the React application
FROM node:16-alpine as build

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to work directory
COPY package*.json ./

# Copy the rest of your app's source code from your host to your image filesystem.
COPY . .

# Install dependencies
RUN npm install

# Build the app
RUN npm run build

# Stage 2: Serve the React application from Nginx server
FROM nginx:alpine

# Copy the build output to replace the default nginx contents.
COPY --from=build /app/build /usr/share/nginx/html

# Copy the nginx configuration file to the container
COPY default.conf /etc/nginx/conf.d/default.conf

# Start Nginx and keep it running in the foreground
CMD ["nginx", "-g", "daemon off;"]
