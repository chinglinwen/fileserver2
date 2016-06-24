#!/bin/sh
# registering a service

service="$( cat <<eof
{
  "Name": "fileserver2",
  "Tags": [
    "master"
  ],
  "Address":"192.168.100.94",
  "Port": 8000,
  "Check": {
    "HTTP": "http://localhost:8000",
    "Interval": "3s"
  }
}
eof
)"

url="http://localhost:8500/v1/agent/service/register"
curl -X PUT "$url" -d "$service"

# end.

