
function start {
  Y='\033[1;33m'
  NC='\033[0m'
  printf "${Y}\n - $@...\n${NC}"
}
function complete {
  G='\033[0;32m'
  NC='\033[0m'
  printf "${G}\n - Completed: $@.\n${NC}"
}
function build_binary_mac {
  STEPNAME="Building macOS binary"
  start $STEPNAME
  go build -o bin/chat_mac .
  complete $STEPNAME
}
function build_binary_alpine {
  STEPNAME="Building Alpine Linux binary"
  start $STEPNAME
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/chat_alpine .
  complete $STEPNAME
}
function build_binaries {
  build_binary_mac
  build_binary_alpine
}

function build_image {
  STEPNAME="Building Docker image"
  start $STEPNAME
  docker build -t ivanturianytsia/bp-chat:latest .
  complete $STEPNAME
}
function push_image {
  STEPNAME="Push Docker image to Docker Hub"
  start $STEPNAME
  docker push ivanturianytsia/bp-chat:latest

  HASH=$(git rev-parse HEAD)
  docker tag ivanturianytsia/bp-chat:latest ivanturianytsia/bp-chat:$HASH
  docker push ivanturianytsia/bp-chat:$HASH
  complete $STEPNAME
}

function deploy {
  STEPNAME="Deploy Docker service localy"
  start $STEPNAME

  HASH=$(git rev-parse HEAD)
  docker service update \
    --image ivanturianytsia/bp-chat:$HASH \
    bp-chat

  complete $STEPNAME
}

case $1 in
  mac)
    build_binary_mac
    ;;
  alpine)
    build_binary_alpine
    ;;
  both)
    build_binaries
    ;;
  image)
    build_image
    ;;
  push)
    push_image
    ;;
  deploy)
    deploy
    ;;
  all)
    build_binaries
    build_image
    push_image
    ;;
esac
