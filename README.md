# NavPool API


The NavPool API is the single actor in the Pool to interact with the Core instance.

It is responsible for setting up Pool Spending and ColdStaking addresses and provides the interface for community fund voting.

## Endpoints

### 1. Create new pool address
```
GET /address/:address/pool
```

Response:
```
{
  "spendingAddress":    "mptAeDyQQgajQoAQNH3NCtgN2bRH6G6iaP",
  "stakingAddress":     "mwsCxYiexGPBBz9LrGNovj9hyfgotyDhyB",
  "coldStakingAddress": "2aR8FeFpSssfU6rLgsL7R7HFLwcPz2VRoKHS5wTEWWAuMQAAAarPrnMWBL333G"
}
```

### 2. Validate address
```
GET /address/:address/validate
```

Response:
```
{
  "isvalid":         true,
  "address":         "2aR8FeFpSssfU6rLgsL7R7HFLwcPz2VRoKHS5wTEWWAuMQAAAarPrnMWBL333G",
  "stakingaddress":  "mwsCxYiexGPBBz9LrGNovj9hyfgotyDhyB",
  "spendingaddress": "mptAeDyQQgajQoAQNH3NCtgN2bRH6G6iaP",
}
```

### 3. Community fund vote
```
POST /community-fund/:type/vote (type: proposal|payment-request)
  
Form Data: {
  "address": "mptAeDyQQgajQoAQNH3NCtgN2bRH6G6iaP",
  "hash":    "987f5e17d99e0ca5f86d8360f5ead5899f7e76351a79171956288d33b5e3b6e0",
  "vote":    "yes|no|remove"
}
```

Response:
```
true|false
```
