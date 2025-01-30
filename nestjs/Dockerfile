# Use the official Node.js 20 Alpine image
FROM node:20-alpine

# Set the working directory
WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install --production

# Copy the rest of the application code
COPY . .

# USER node

# Expose the application port
EXPOSE 3000

# Start the application
CMD ["node", "dist/main"]
