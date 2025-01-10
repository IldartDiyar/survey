FROM postgres:15

ENV POSTGRES_USER=postgres_user
ENV POSTGRES_PASSWORD=very_hard_password!
ENV POSTGRES_DB=survey_db

COPY ./init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
