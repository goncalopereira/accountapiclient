#!/bin/bash
curl -v -H "Content-Type: application/vnd.api+json" \
   --data-urlencode "filter[bank_id]=123456" \
  "http://localhost:8080/v1/organisation/accounts"