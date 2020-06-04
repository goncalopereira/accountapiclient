FROM form3tech/interview-accountapi:v1.0.0-4-g63cf8434
USER root
RUN apk add --no-cache bash # required for wait-for-it.sh
USER appuser
COPY --chown=appuser build/ci/wait-for-it.sh /app
ENTRYPOINT ["./wait-for-it.sh", "postgresql:5432", "--", "./entrypoint.sh"]
