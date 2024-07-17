# Use Node.js LTS version as the base image
FROM node:lts AS build

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY frontend/package.json frontend/yarn.lock ./

# Install dependencies
RUN yarn install --frozen-lockfile

# Copy the rest of the application code
COPY frontend ./

# Build the application
RUN yarn build

# Stage 2: Serve the built application with Nginx
FROM nginx:alpine

# Copy custom nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy the built app from the previous stage into the Nginx image
COPY --from=build /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start Nginx and keep it running in the foreground
CMD ["nginx", "-g", "daemon off;"]
