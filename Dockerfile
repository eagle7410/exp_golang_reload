FROM debian:stretch-slim
RUN apt-get -qq update
RUN apt-get -qq -y install curl
RUN mkdir -p /app
WORKDIR /app
ADD ./apps/Back ./Back
CMD /app/Back

EXPOSE 8080
