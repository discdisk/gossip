set -e
go build -o out/gossip
cd out
# ./maelstrom/maelstrom test -w echo --bin gossip --node-count 1 --time-limit 10
# ./maelstrom/maelstrom test -w unique-ids --bin gossip --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition
./maelstrom/maelstrom test -w broadcast --bin gossip --node-count 1 --time-limit 20 --rate 10
# ./maelstrom/maelstrom test -w broadcast --bin gossip --node-count 5 --time-limit 20 --rate 10