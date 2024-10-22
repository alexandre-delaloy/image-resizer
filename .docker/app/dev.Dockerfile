FROM node:lts-alpine

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

EXPOSE 8010

# Start service
CMD ["npm", "run", "serve"]
