# [WIP]rum-go-sdk  

### WIP
> ‼️ Cause the hight speed of quorum evolution, this SDK may not be 100% suitable with latest quorum version.  
- [ ] Management  
- [ ] Group  
- [ ] Keystore  
- [ ] Tools  
- [ ] Apps


# ❓ What is it?  
This is a third party quorum SDK in golang. 

## What is rum?  
Rum System is flexable de-communicating protocol. You can treat it like well-backuped DB for the app that you want to build. [More able Rum](https://github.com/rumsystem/quorum)  

# How to start?  
```
rum_client := rumgosdk.Connect('http://127.0.0.1:8002')
rum_client.Node()
