package main

import (
	"flag"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {

	mode := flag.String("mode", "", "mining mode: single / multi")
	diff := flag.String("diff", "", "mining difficulty: 1 number for single mode, 4 numbers seperated by ',' for multi mode")
	flag.Parse()

	switch *mode {
	case "single":
		singleMineTest(*diff)
	case "multi":
		multiMineTest(*diff)
	default:
		glog.Fatalln("unsupported mode, exit")
	}

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGUSR2)
	select {
	case sig := <-signalChan:

		switch sig {
		default:
			glog.Infof("now exiting, sig=%v", sig.String())
		}
		break
	}
}

func singleMineTest(diff string) {

	difficulty, err := strconv.Atoi(diff)

	if err != nil {
		glog.Fatalln("invalid diff, exit")
	}

	job := NewJob(0, uint64(difficulty))

	job.report()

	singleMiner0 := NewSingleMiner(0, job)
	singleMiner0.mine()

	singleMiner1 := NewSingleMiner(1, job)
	singleMiner1.mine()
}

func multiMineTest(diff string) {

	ret := strings.Split(diff, ",")

	var difficulties [4]uint64
	for i,one := range ret{
		difficulty, err := strconv.Atoi(one)
		if err != nil {
			glog.Fatalln("invalid diff, exit")
		}
		difficulties[i] = uint64(difficulty)
	}

	job0 := NewJob(0, difficulties[0])
	job1 := NewJob(1, difficulties[1])
	job2 := NewJob(2, difficulties[2])
	job3 := NewJob(3, difficulties[3])

	job0.report()
	job1.report()
	job2.report()
	job3.report()

	jobs := []*Job{job0, job1, job2, job3}

	multiMiner0 := NewMultiMiner(0, jobs)
	multiMiner0.mine()

	multiMiner1 := NewMultiMiner(1, jobs)
	multiMiner1.mine()
}
