# Twirp-Paperclips

## A demo app in golang using the twirp rpc library

More information on this demo and some context on twirp is available in this blog post:

## Installation

Download the protobuf binaries from https://github.com/google/protobuf/releases
Copy the downloaded protoc binary to your GOPATH bin folder
Install the go packages for twirp & protobuf:
```
go get github.com/twitchtv/twirp/protoc-gen-twirp
go get github.com/golang/protobuf/protoc-gen-go
```

## What does this service do?

Inspired by the existential dread of universal paperclips (http://www.decisionproblem.com/paperclips/index2.html), this is a simple service that does three silly things:
- Get the current number of paperclips
- Increment the number of paperclips
- Calculate how close we are to the death of the universe.

From our protobuf file:
```
service UniversalPaperclips {
  rpc GetPaperclips(Empty) returns (Paperclips);
  rpc IncrementPaperclips(Size) returns (Empty);
  rpc CalculateUniverseLifespan(Empty) returns (Dread);
}
```

## Usage
Running the server
```
go run main.go
```

Running the client
```
go run client/main.go
```

Hitting the server with curl:
```
☁  twirpPaperclips [master] ⚡ curl --request "POST" \
     --location "http://localhost:6666/twirp/twirp.paperclips.UniversalPaperclips/IncrementPaperclips" \
     --header "Content-Type:application/json" \
     --data '{"paperclips":5}'
{}%
☁  twirpPaperclips [master] ⚡ curl --request "POST" \
     --location "http://localhost:6666/twirp/twirp.paperclips.UniversalPaperclips/GetPaperclips" \
     --header "Content-Type:application/json" \
     --data '{}'
{"paperclips":11}%
☁  twirpPaperclips [master] ⚡ curl --request "POST" \
     --location "http://localhost:6666/twirp/twirp.paperclips.UniversalPaperclips/CalculateUniverseLifespan" \
     --header "Content-Type:application/json" \
     --data '{}'
{"paperclips":11,"universeLifespan":"42"}%
```
