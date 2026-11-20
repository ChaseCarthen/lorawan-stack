go run ./cmd/ttn-lw-stack is-db create-oauth-client \
  --id localhost-console \
  --name "Local Development Console" \
  --owner admin \
  --secret "$TTN_LW_CONSOLE_OAUTH_CLIENT_SECRET" \
  --redirect-uri "http://localhost:8080/console/oauth/callback" \
  --redirect-uri "/console/oauth/callback" \
  --logout-redirect-uri "http://localhost:8080/console" \
  --logout-redirect-uri "/console"