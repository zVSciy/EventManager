FROM node:18-alpine
WORKDIR /app
COPY notifications-svelte-frontend .
RUN npm ci
RUN npm run build
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]