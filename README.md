# tower-slack

make

export DOCKER_USER="YOUR_NAME"

make docker-image

docker run -d -p 8080:8080 $DOCKER_USER/tower-slack:$VERSION --secret=TOWER_WEBHOOK_SECRET

TOWER_WEBHOOK_SECRET can be empty.
