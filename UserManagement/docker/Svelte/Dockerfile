FROM node:18-alpine

WORKDIR /app

COPY ./frontend /app

RUN npm install -g npm@10.9.0
# RUN npm ci
# RUN npm run build

CMD ["npm", "run", "dev", "--", "--host"]