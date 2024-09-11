FROM node:19.0.0-slim

WORKDIR /app

COPY ./src/package*.json .

RUN npm install

COPY ./src .

EXPOSE 5173

CMD [ "npm", "run", "dev" ]
