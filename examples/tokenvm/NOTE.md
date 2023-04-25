

./build/token-cli 



# getNodeID

```
curl --location --request POST 'http://127.0.0.1:43197/ext/info' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"info.getNodeID"
}'
```

# isBootstrapped
```
curl --location --request POST 'http://127.0.0.1:43197/ext/info' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"info.isBootstrapped"
}'
```
# getNetworkID

```
curl --location --request POST 'http://127.0.0.1:43197/ext/info' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"info.getNetworkID"
}'
```



# import ANR 

To connects to the Avalanche Network Runner server running in
the background and pulls the URIs of all nodes tracking each chain you
created

```
./build/token-cli chain import-anr
```

# create asset 

```
./build/token-cli action create-asset
```

# mint asset 

```
./build/token-cli action mint-asset
```

# balance 
```
./build/token-cli key balance
```



txID: d6SdaAszimoW1qoLUVJjbhFvVKVUGQqWUk3HapwdwFCY9SBJo

recipient: token1rvzhmceq997zntgvravfagsks6w0ryud3rylh4cdvayry0dl97nsjzf3yp

