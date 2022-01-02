# micor-services-wallets-generation
## Running the python app
docker build -t python .

# this will mount the volume into container and open a terminal
docker run -it -p 5000:5000 -v ${PWD}:/wallets python bash

# in terminal just type
python3 main.py


## Running the go app
docker build --target dev . -t go
# this will mount the volume into container and open a terminal
docker run -it -p 8080:8080 -v ${PWD}:/wallets go sh
# in terminal just type
go run main.go