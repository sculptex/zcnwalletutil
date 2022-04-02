# zcnwalletutil
ZCN wallet utility for offline wallet restoration 

## `./zcnwalletutil`
Create New Wallet and output as list
- Example Output:-
```
client_id:  2047b5703d9f372555d282bee5f7cd2f4ac27614d5a719c8271abcc2a9006617
mnemonic:  congress solar expect duck leopard march agent tray gym price zone pull shop melt speed fog fly alley annual throw vast meadow address wet
public_key:  5fecead94488990c2aa9ece3d84aef9afbe22e14cb9113135b63843033503818fb96380c2a3d1dba9ad9b76d914a401746dadee993c6c35b1e76b2dff9b94a0a
private_key:  395d885705b549041e110696dbc5749d90c7062d34030ab39dd62bd44d183f0b
```

## `./zcnwalletutil --json`
Create New Wallet and output as JSON
- Example Output:-
```
{"client_id":"2047b5703d9f372555d282bee5f7cd2f4ac27614d5a719c8271abcc2a9006617","mnemonic":"congress solar expect duck leopard march agent tray gym price zone pull shop melt speed fog fly alley annual throw vast meadow address wet","private_key":"395d885705b549041e110696dbc5749d90c7062d34030ab39dd62bd44d183f0b","public_key":"5fecead94488990c2aa9ece3d84aef9afbe22e14cb9113135b63843033503818fb96380c2a3d1dba9ad9b76d914a401746dadee993c6c35b1e76b2dff9b94a0a"}
```

## `./zcnwalletutil --mnemonic "xxxxx"`
Restore Wallet using mnemonic phrase and output as list
- Example Output:-
```
client_id:  2047b5703d9f372555d282bee5f7cd2f4ac27614d5a719c8271abcc2a9006617
mnemonic:  congress solar expect duck leopard march agent tray gym price zone pull shop melt speed fog fly alley annual throw vast meadow address wet
public_key:  5fecead94488990c2aa9ece3d84aef9afbe22e14cb9113135b63843033503818fb96380c2a3d1dba9ad9b76d914a401746dadee993c6c35b1e76b2dff9b94a0a
private_key:  395d885705b549041e110696dbc5749d90c7062d34030ab39dd62bd44d183f0b
```

## `./zcnwalletutil --mnemonic "xxxxx" --json`
Restore Wallet using mnemonic phrase and output as JSON
- Example Output:-
```
{"client_id":"e10d6140741e21db2433046f15a0f1f0f5421aa9489307e2b1621c001446ae62","mnemonic":"indoor boost improve become trumpet wear fatigue squeeze host object also beef tube remember extra bubble such cigar crouch super palm draw page reward","private_key":"ab15cd5c45ef4dfdc904a83c7676dd7b4dc6a9590fd148e6e343f6ff91e81b05","public_key":"1bfb722c79334ffaa5dc72695e926818c9721ae548b44148ff524d814197a4239e99d36f32172ac6b6d5b256c604650f9ad9fa60487c81062af6687947752591"}
```
