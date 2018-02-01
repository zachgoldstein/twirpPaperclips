curl --request "POST" \
     --location "http://localhost:6666/twirp/twirp.paperclips.UniversalPaperclips/GetPaperclips" \
     --header "Content-Type:application/json" \
     --data '{}' \
     --verbose

{"paperclips":6}

curl --request "POST" \
    --location "http://localhost:6666/twirp/twirp.paperclips.UniversalPaperclips/IncrementPaperclips" \
    --header "Content-Type:application/json" \
    --data '{"paperclips":5}' \
    --verbose

curl --request "POST" \
    --location "http://localhost:6666/twirp/twirp.paperclips.UniversalPaperclips/CalculateUniverseLifespan" \                                      
    --header "Content-Type:application/json" \
    --data '{}' \
    --verbose

{"paperclips":6,"universeLifespan":"42"}
