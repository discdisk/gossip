set -e
go build -o out/gossip
cd out
./maelstrom/maelstrom test -w echo --bin gossip --node-count 1 --time-limit 10