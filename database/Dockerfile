FROM postgres:14.1

ENV POSTGRES_USER admin
ENV POSTGRES_PASSWORD admin
ENV POSTGRES_DB message
ADD init-message.sql /docker-entrypoint-initdb.d/
