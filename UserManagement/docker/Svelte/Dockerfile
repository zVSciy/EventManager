FROM node:18-alpine

WORKDIR /app

# Copy package.json and package-lock.json first to leverage Docker caching
COPY frontend/package*.json ./

# Install dependencies
RUN npm install --legacy-peer-deps

# Copy the rest of the frontend source code
COPY frontend ./

# Expose the port used by the frontend application
EXPOSE 3000

# Command to run the development server
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]


# FROM node:18-alpine

# WORKDIR /app

# COPY package*.json /app

# RUN npm install -g npm@10.9.0

# COPY frontend /app

# # RUN npm install
# # RUN npm ci
# # RUN npm run build

# CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]