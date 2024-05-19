# Use the latest Keycloak image
FROM quay.io/keycloak/keycloak:latest

ENV KC_DB=postgres
ENV KC_DB_URL=jdbc:postgresql://host.docker.internal:5432/keycloak
ENV KC_DB_USERNAME=postgres
ENV KC_DB_PASSWORD=postgres
ENV KC_HOSTNAME=localhost

ENV KEYCLOAK_ADMIN=admin
ENV KEYCLOAK_ADMIN_PASSWORD=admin

EXPOSE 8080


ENTRYPOINT ["/opt/keycloak/bin/kc.sh", "start-dev", "--import-realm"]


