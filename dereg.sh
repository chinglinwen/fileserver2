#!/bin/sh
# de-register a service

service="fileserver2"

url="http://localhost:8500/v1/agent/service/deregister/$service"
curl "$url"

# end.

