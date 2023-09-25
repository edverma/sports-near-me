all: shutdown server fe

fe:
	cd ./client && npm run dev &

server: cache nosql status
	cd ./server && go run . &

cache:
	redis-server --daemonize yes

nosql:
	mongod --config /opt/homebrew/etc/mongod.conf --fork

shutdown: status
	pgrep -f redis-server | xargs kill -9 && pgrep -f mongod | xargs kill -9 && pgrep -f "go run" | xargs kill -9 && pgrep -f "go-build" | xargs kill -9 && pgrep -f "node" |  xargs kill -9

status:
	ps aux | grep redis-server && ps aux | grep mongod && ps aux | grep "go run" && ps aux | grep "go-build"

