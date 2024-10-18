FROM node:18-alpine
WORKDIR /app
COPY src_svelte_frontend .
RUN npm ci
RUN npm run build
CMD ["npm", "run", "dev", "--", "--host"]