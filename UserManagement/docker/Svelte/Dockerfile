FROM node:lts-alpine

WORKDIR /app

COPY frontend/package*.json ./

<<<<<<< HEAD
# Install dependencies
RUN npm install --legacy-peer-deps
=======
>>>>>>> main
RUN npm install

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]