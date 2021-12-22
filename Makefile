MESSAGE ?= "hello"

run:
	docker-compose up --build

message:
	PGPASSWORD=admin psql -h localhost -p 5432 -d message -U admin -c "INSERT INTO message (message_text) VALUES ('$(MESSAGE)');"

websocket:
	curl --include \
	--no-buffer \
	--header "Connection: Upgrade" \
	--header "Upgrade: websocket" \
	--header "Host: localhost:4837" \
	--header "Origin: http://localhost:3721" \
	--header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
	--header "Sec-WebSocket-Version: 13" \
	http://localhost:4837/messages
