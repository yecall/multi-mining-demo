# Test report of multi mining

## Single mining

Environment: MacBook Pro 2.8 GHz Intel Core i7

Miner count：2 (goroutine)

Run duration：1319s

Difficulty：33554432

Avg block interval：12.816s

Command:
```
./build/multi-mining-demo -logtostderr=true -mode=single -diff=33554432
```

Original output:
```
job report:
	shard num: 0
	difficulty: 33554432
	start time: 2019-06-12 19:01:15
	duration: 1319
	block count: 103
	avg block interval: 12.816
```

## Multi mining

Environment: MacBook Pro 2.8 GHz Intel Core i7

Miner count：2 (goroutine)

Run duration：1319s

Difficulties of 4 shards：33554432,8388608,16777216,67108864

Avg block interval of shard#0：12.815s

Command:
```
./build/multi-mining-demo -logtostderr=true -mode=multi -diff=33554432,8388608,16777216,67108864
```

Original output:

```
job report:
	shard num: 2
	difficulty: 16777216
	start time: 2019-06-12 21:12:47
	duration: 1319
	block count: 232
	avg block interval: 5.690
job report:
	shard num: 1
	difficulty: 8388608
	start time: 2019-06-12 21:12:47
	duration: 1319
	block count: 424
	avg block interval: 3.113
job report:
	shard num: 0
	difficulty: 33554432
	start time: 2019-06-12 21:12:47
	duration: 1319
	block count: 103
	avg block interval: 12.815
job report:
	shard num: 3
	difficulty: 67108864
	start time: 2019-06-12 21:12:47
	duration: 1319
	block count: 50
	avg block interval: 26.400
```

## Conclusion
With the same pysical hashrate, 
the avg block interval of one shard in multi mining mode is almost the same with that in single mining mode 