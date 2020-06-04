#!/bin/bash
curl -v -X POST -H "Content-Type: application/vnd.api+json" \
  http://localhost:8080/v1/organisation/accounts \
  -d @./test/integration/data/create.json
