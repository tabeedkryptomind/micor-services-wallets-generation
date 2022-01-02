# micor-services-wallets-generation

## Running the python app
```bash
cd  pyhton-wallets-btc 
docker build -t python .

# This will mount the volume into container and open a terminal
docker run -it -p 5000:5000 -v ${PWD}:/wallets python bash


# To run the application
python3 main.py
```

## Running the go app
```bash
cd go-wallets-eth
docker build --target dev . -t go
# This will mount the volume into container and open a terminal
docker run -it -p 8080:8080 -v ${PWD}:/wallets go sh
# To run the application 
go run main.go
```