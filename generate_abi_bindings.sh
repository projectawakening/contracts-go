set -o errexit
set -o nounset

worldAbiTemp=$(mktemp)
ERC2771WorldAbiTemp=$(mktemp)
ERC2771ForwarderAbiTemp=$(mktemp)
ERC20AbiTemp=$(mktemp)

# this is the ABI with the prefixed functions so we don't need to use system id
curl -s https://abi-export.projectawakening.io/IWorld-${VERSION}.abi.json > ${worldAbiTemp}

# this is the ABI without the prefix so we can use system id
curl -s https://abi-export.projectawakening.io/ERC2771IWorld-${VERSION}.abi.json > ${ERC2771WorldAbiTemp}

curl -s https://abi-export.projectawakening.io/ERC2771Forwarder-${VERSION}.abi.json > ${ERC2771ForwarderAbiTemp}

curl -s https://abi-export.projectawakening.io/IERC20-${VERSION}.abi.json > ${ERC20AbiTemp}

abigen --abi=${worldAbiTemp} --pkg=world --out=bindings/world/world.go
abigen --abi=${ERC2771WorldAbiTemp} --pkg=ERC2771World --out=bindings/ERC2771World/ERC2771World.go
abigen --abi=${ERC2771ForwarderAbiTemp} --pkg=ERC2771Forwarder --out=bindings/ERC2771Forwarder/ERC2771Forwarder.go
abigen --abi=${ERC20AbiTemp} --pkg=ERC20 --out=bindings/ERC20/ERC20.go
