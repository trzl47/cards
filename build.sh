#!/usr/bin/env bash
mkdir services/react-nginx/build && \
cp -r services/react/build/server services/react-nginx/build && \
cp -r services/react/build/client services/react-nginx/build && \
docker build -t trzl/portfolio:augment-react.1 -f ./services/react-nginx/Dockerfile.prod ./services/react-nginx && \
# docker push trzl/portfolio:augment-react.1 && \
docker build -t trzl/portfolio:augment-apollo.2 -f ./services/apollo/Dockerfile ./services/apollo && \
# docker push trzl/portfolio:augment-apollo.2 && \
rm -rf services/react-nginx/build && \
docker system prune
echo "Build complete!"