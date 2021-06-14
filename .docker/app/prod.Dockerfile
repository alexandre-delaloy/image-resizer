FROM node:lts-alpine as build

# ----- SETUP -----

# Set the current working with go absolute path
WORKDIR /app

# ----- INSTALL -----

# Copy package.json + package-lock.json for install before full copy
COPY app/package*.json ./

# Install all dependencies
RUN npm install

# ----- COPY + RUN -----

# Copy the source from the current directory to the container
COPY app/ .

# Build app
RUN npm run build

# ----- NGINX -----

FROM nginx:1.16.0-alpine

# Copy builded app
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

# Start service
CMD ["nginx", "-g", "daemon off;"]