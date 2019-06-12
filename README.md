# Multi mining demo

**Multi mining** is designed to make mining hashrate act on multi chains(or shards) concurrently to resist 1% attach of the PoW consensus.

**Multi mining demo** shows how it works.

## Install
```
mkdir -p build
cd build
go build ..
```

## Usage

### Args description
```
-mode string
    	mining mode: single / multi
-diff string
    	mining difficulty: a number for single mode , 4 number seperated by ',' for multi mode
```

### Example for single mode
```
./build/multi-mining-demo -logtostderr=true -mode=single -diff=16777216
```

### Example for multi mode
```
./build/multi-mining-demo -logtostderr=true -mode=multi -diff=16777216,8388608,33554432,33554432
```

## Testing
[Test report](./docs/test_report.md)