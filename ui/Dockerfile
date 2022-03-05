FROM node:16 AS base

COPY package.json .
COPY package-lock.json .
RUN npm install

FROM base as dependencies
COPY pages /pages
COPY public /public
COPY styles /styles
COPY components /components
COPY next* /
COPY tsconfig.json /

EXPOSE 3000
RUN npm run build
CMD ["npm", "start"]