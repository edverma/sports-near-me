all:

server: nosql_db
	cd server && go run main.go

nosql_db:
	pgrep mongod || mongod --config /opt/homebrew/etc/mongod.conf --fork

client:
	cd client && npm run dev
