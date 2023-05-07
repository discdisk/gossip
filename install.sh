# sh -c "$(wget https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh -O -)"
mkdir out
cd out
cp /tmp/maelstrom.tar.bz2 .
tar -jxf maelstrom.tar.bz2
rm -rf maelstrom.tar.bz2