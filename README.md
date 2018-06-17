# BalanceBadge.io [![Build Status](https://travis-ci.com/hunterlong/balancebadge.svg?branch=master)](https://travis-ci.com/hunterlong/balancebadge)
There's a badge for everything else, now you can have a badge for your cryptocurrency balances. 

## Badge URL
You can easily format a badge by inserting your Coin type and the address.
```
https://img.balancebadge.io/<COIN>/<ADDRESS>.svg
```

### Accepted Cryptocurrencies
- [x] Bitcoin `btc` [![Balance](https://img.balancebadge.io/btc/1LhWMukxP6QGhW6TMEZRcqEUW2bFMA4Rwx.svg)](https://blockchain.info/address/1LhWMukxP6QGhW6TMEZRcqEUW2bFMA4Rwx)
- [x] Ethereum `eth` [![Balance](https://img.balancebadge.io/eth/0x004f3e7ffa2f06ea78e14ed2b13e87d710e8013f.svg)](https://etherscan.io/address/0x004f3e7ffa2f06ea78e14ed2b13e87d710e8013f)
- [ ] Litecoin `ltc`
- [ ] ERC20 Tokens `token`

## Customize Badge
You can send parameters with the SVG request to customize your badge on the fly. 
- `label` Text for the left side of the badge [![Balance](https://img.balancebadge.io/btc/1LhWMukxP6QGhW6TMEZRcqEUW2bFMA4Rwx.svg?label=MtGOX)](https://blockchain.info/address/1LhWMukxP6QGhW6TMEZRcqEUW2bFMA4Rwx)
- `color` Hex color for the right side of the badge [![Balance](https://img.balancebadge.io/eth/0x004f3e7ffa2f06ea78e14ed2b13e87d710e8013f.svg?color=ffb121)](https://etherscan.io/address/0x004f3e7ffa2f06ea78e14ed2b13e87d710e8013f)
