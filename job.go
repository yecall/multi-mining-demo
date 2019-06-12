package main

import (
	"fmt"
	"github.com/golang/glog"
	"time"
)

type Job struct {
	shardNum   uint32
	difficulty uint64
	blockCount int64
	startTsN   int64
}

func NewJob(shardNum uint32, difficulty uint64) *Job {
	job := new(Job)
	job.shardNum = shardNum
	job.difficulty = difficulty
	job.blockCount = 0
	job.startTsN = time.Now().UnixNano()

	return job
}

func (job *Job) report() {
	ticker := time.NewTicker(time.Second * 60)
	go func() {
		for _ = range ticker.C {
			job.reportOnce()
		}
	}()
}

func (job *Job) reportOnce() {

	now := time.Now().UnixNano()
	durationN := now - job.startTsN
	avgBlockIntervalN := int64(0)
	if job.blockCount > 0 {
		avgBlockIntervalN = durationN / job.blockCount
	}

	durationS := fmt.Sprintf("%d", durationN/1000000000)
	avgBlockIntervalS := fmt.Sprintf("%.3f", float64(avgBlockIntervalN)/1000000000)

	reportStr := fmt.Sprintf(
		"\njob report:\n\tshard num: %v\n"+
			"\tdifficulty: %v\n"+
			"\tstart time: %v\n"+
			"\tduration: %v\n"+
			"\tblock count: %v\n"+
			"\tavg block interval: %v",
		job.shardNum, job.difficulty, time.Unix(job.startTsN/1000000000, 0).Format("2006-01-02 15:04:05"),
		durationS, job.blockCount, avgBlockIntervalS)

	glog.Info(reportStr)

}
