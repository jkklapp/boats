## Boats

Forked from https://github.com/typenil/go-docker

# Running the container

To build the images run `docker-compose build`
To run the containers `docker-compose up`

I got stuck trying to figure out how to communicate
between the containers, I have tried different versions in the compose
file with no luck.

If I had more time I would solve this plus adding unit tests to both
packages. Then I would think on extending both the ingestor (to accept different
types of data) and the datastore (to be able to use different databases to 
persist info)