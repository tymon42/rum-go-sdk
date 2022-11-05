# rum-go-sdk  
> ‼️ Because of the high speed of quorum evolution, this SDK may not be 100% suitable with latest quorum version.  

# What is it?  
This is a third party quorum SDK in golang. 

## What is rum?  
Rum System is flexable communicating protocol. You can treat it like well-backuped DB for the app that you want to build. [More about Rum](https://github.com/rumsystem/quorum)  

# Usage 
install: `go get github.com/tymon42/rum-go-sdk`

```
import (
    "github.com/tymon42/rum-go-sdk"
)

rum_client := rumgosdk.Connect('http://127.0.0.1:8002')

rum_client.Node()
```
