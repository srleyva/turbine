# Build Image 
FROM node:alpine
RUN mkdir /tmp/wkdir
WORKDIR /tmp/wkdir
ADD . ./
RUN npm install -g @angular/cli && npm install
RUN ng build --prod --build-optimizer=false

# Runtime Image
FROM nginx:alpine
COPY nginx.conf /etc/nginx/nginx.conf
WORKDIR /usr/share/nginx/html
COPY --from=0 /tmp/wkdir/dist/ .
