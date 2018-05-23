# tower-slack

[![Build Status](https://travis-ci.org/leeeboo/tower-slack.svg?branch=master)](https://travis-ci.org/leeeboo/tower-slack)
[![Docker Status](https://dockerbuildbadges.quelltext.eu/status.svg?organization=leeeboo&repository=tower-slack)](https://hub.docker.com/r/leeeboo/tower-slack/builds/)

## BUILD

docker build -t YOUR_IMAGE_NAME .

## RUN 

docker run -d -p 8080:8080 YOUR_IMAGE_NAME --secret=TOWER_WEBHOOK_SECRET

TOWER_WEBHOOK_SECRET can be empty.

## USAGE

Go to Tower.im and set your project's webhook as the url which your tower-slack running. And set send all events.

For example:

    If your tower-slack running at http://tower-slack.foo.com.

    You should set the webhook to  http://tower-slack.foo.com/$PATH.

    The `$PATH` can be found in your slack incoming webhook setting, just like `https://hooks.slack.com/services/$PATH`.

    Then your tower events will be send to `https://hooks.slack.com/services/$PATH`.
