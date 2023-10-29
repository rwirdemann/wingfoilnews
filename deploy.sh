npm run build
scp -r build/* wingfoilnews@95.217.180.178:~/news.wingbuddies.de/build
scp -r package*.json wingfoilnews@95.217.180.178:~/news.wingbuddies.de
ssh wingfoilnews@95.217.180.178 "cd news.wingbuddies.de; npm ci --omit dev"
ssh wingfoilnews@95.217.180.178 "pkill node"
ssh wingfoilnews@95.217.180.178 "cd news.wingbuddies.de; sh -c 'ORIGIN=https://news.wingbuddies.de node build/index.js > news.wingbuddies.out 2>&1 &'"
