pair:
	solc --abi pancake/IPancakePair.sol -o pancake --overwrite
	solc --bin pancake/IPancakePair.sol -o pancake --overwrite
	abigen --bin=pancake/IPancakePair.bin --abi=pancake/IPancakePair.abi --pkg=pair --out=IPancakePair.go
factory:
	solc --abi pancake/IpancakeFactory.sol -o pancake --overwrite
	solc --bin pancake/IpancakeFactory.sol -o pancake --overwrite
	abigen --bin=pancake/IpancakeFactory.bin --abi=pancake/IpancakeFactory.abi --pkg=factory --out=IpancakeFactory.go
run:
	go run main.go