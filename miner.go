package main

import (
	"github.com/golang/glog"
)

type SingleMiner struct {
	coinbase uint8
	job      *Job
}

func NewSingleMiner(coinbase uint8, job *Job) *SingleMiner {
	miner := new(SingleMiner)
	miner.coinbase = coinbase
	miner.job = job
	return miner
}

func (singleMiner *SingleMiner) mine() {

	go func() {
		merkleRoot := randomMerkleRoot()
		nonce := uint64(0);
		target := diffToTarget(singleMiner.job.difficulty)

		singleHead := new(SingleHead)
		singleHead.coinbase = singleMiner.coinbase
		singleHead.shardNum = singleMiner.job.shardNum
		singleHead.difficulty = singleMiner.job.difficulty

		for {

			singleHead.merkleRoot = merkleRoot
			singleHead.nonce = nonce

			hash := singleHead.hash()
			hashBig := hashToHashBig(hash)
			hashHex := hashToHashHex(hash)

			if hashBig.Cmp(target) < 0 {
				singleMiner.job.blockCount++
				glog.Infof("new block found, miner=%v, hashHex=%v", singleMiner.coinbase, hashHex)
			}

			nonce++

			if nonce >= 1000000 {
				merkleRoot = randomMerkleRoot()
				nonce = 0
			}

		}
	}()
}

type MultiMiner struct {
	coinbase uint8
	jobs     []*Job
}

func NewMultiMiner(coinbase uint8, jobs []*Job) *MultiMiner {
	miner := new(MultiMiner)
	miner.coinbase = coinbase
	miner.jobs = jobs
	return miner
}

func (multiMiner *MultiMiner) makeSingleHead(job *Job, merkleRoot [32]byte) *SingleHead {
	singleHead := new(SingleHead)
	singleHead.coinbase = multiMiner.coinbase
	singleHead.shardNum = job.shardNum
	singleHead.difficulty = job.difficulty
	singleHead.merkleRoot = merkleRoot

	return singleHead
}

func (multiMiner *MultiMiner) makeHeadRoot(hash0, hash1, hash2, hash3 [32]byte) [32]byte {

	hash01 := merkleParent(hash0, hash1)
	hash23 := merkleParent(hash2, hash3)
	hash := merkleParent(hash01, hash23)
	return hash
}

func (multiMiner *MultiMiner) mine() {

	go func() {

		job0 := multiMiner.jobs[0]
		job1 := multiMiner.jobs[1]
		job2 := multiMiner.jobs[2]
		job3 := multiMiner.jobs[3]

		merkleRoot0 := randomMerkleRoot()
		merkleRoot1 := randomMerkleRoot()
		merkleRoot2 := randomMerkleRoot()
		merkleRoot3 := randomMerkleRoot()

		target0 := diffToTarget(job0.difficulty)
		target1 := diffToTarget(job1.difficulty)
		target2 := diffToTarget(job2.difficulty)
		target3 := diffToTarget(job3.difficulty)

		singleHead0 := multiMiner.makeSingleHead(job0, merkleRoot0)
		singleHead1 := multiMiner.makeSingleHead(job1, merkleRoot1)
		singleHead2 := multiMiner.makeSingleHead(job2, merkleRoot2)
		singleHead3 := multiMiner.makeSingleHead(job3, merkleRoot3)

		headRoot := multiMiner.makeHeadRoot(singleHead0.hash(), singleHead1.hash(), singleHead2.hash(), singleHead3.hash())

		nonce := uint64(0)
		multiHead := new(MultiHead)

		for {

			multiHead.headRoot = headRoot
			multiHead.shardCount = uint32(len(multiMiner.jobs))
			multiHead.nonce = nonce

			hash := multiHead.hash()
			hashBig := hashToHashBig(hash)
			hashHex := hashToHashHex(hash)

			found := false
			if hashBig.Cmp(target0) < 0 {
				job0.blockCount++
				glog.Infof("new block found, shardNum=%v, miner=%v, hashHex=%v", job0.shardNum, multiMiner.coinbase, hashHex)

				found = true
				merkleRoot0 = randomMerkleRoot()
				singleHead0.merkleRoot = merkleRoot0
			}
			if hashBig.Cmp(target1) < 0 {
				job1.blockCount++
				glog.Infof("new block found, shardNum=%v, miner=%v, hashHex=%v", job1.shardNum, multiMiner.coinbase, hashHex)

				found = true
				merkleRoot1 = randomMerkleRoot()
				singleHead1.merkleRoot = merkleRoot1
			}
			if hashBig.Cmp(target2) < 0 {
				job2.blockCount++
				glog.Infof("new block found, shardNum=%v, miner=%v, hashHex=%v", job2.shardNum, multiMiner.coinbase, hashHex)

				found = true
				merkleRoot2 = randomMerkleRoot()
				singleHead2.merkleRoot = merkleRoot2
			}
			if hashBig.Cmp(target3) < 0 {
				job3.blockCount++
				glog.Infof("new block found, shardNum=%v, miner=%v, hashHex=%v", job3.shardNum, multiMiner.coinbase, hashHex)

				found = true
				merkleRoot3 = randomMerkleRoot()
				singleHead3.merkleRoot = merkleRoot3
			}

			if found {
				nonce = 0
				headRoot = multiMiner.makeHeadRoot(singleHead0.hash(), singleHead1.hash(), singleHead2.hash(), singleHead3.hash())
			}

			nonce++

			if nonce >= 1000000 {
				merkleRoot0 = randomMerkleRoot()
				merkleRoot1 = randomMerkleRoot()
				merkleRoot2 = randomMerkleRoot()
				merkleRoot3 = randomMerkleRoot()

				singleHead0.merkleRoot = merkleRoot0
				singleHead1.merkleRoot = merkleRoot1
				singleHead2.merkleRoot = merkleRoot2
				singleHead3.merkleRoot = merkleRoot3

				nonce = 0
				headRoot = multiMiner.makeHeadRoot(singleHead0.hash(), singleHead1.hash(), singleHead2.hash(), singleHead3.hash())
			}

		}
	}()
}
