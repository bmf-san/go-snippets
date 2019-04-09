// see: https://github.com/Lebonesco/go_worker_pool
package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes creates a random string.
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// CreateJobs creates a list of jobs.
func CreateJobs(amount int) []string {
	var jobs []string

	for i := 0; i < amount; i++ {
		jobs = append(jobs, RandStringRunes(8))
	}
	return jobs
}

// DoWork mimics any types of job that can be run concurrently.
func DoWork(word string, id int) {
	h := fnv.New32a()
	h.Write([]byte(word))
	time.Sleep(time.Second)
	log.Printf("worker [%d] - created hash [%d] from word [%s]\n", id, h.Sum32(), word)
}

// Work is a struct for Worker
type Work struct {
	ID  int
	Job string
}

// Worker is a worker.
type Worker struct {
	ID            int
	WorkerChannel chan chan Work
	Channel       chan Work
	End           chan bool
}

// Start starts a worker.
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			select {
			case job := <-w.Channel:
				DoWork(job.Job, w.ID)
			case <-w.End:
				return
			}
		}
	}()
}

// Stop stops a worker.
func (w *Worker) Stop() {
	log.Printf("worker [%d] is stopping", w.ID)
	w.End <- true
}

var WorkerChannel = make(chan chan Work)

// Collector is a collector
type Collector struct {
	Work chan Work
	End  chan bool
}

// StartDispatcher starts a dispatcher.
func StartDispatcher(workerCount int) Collector {
	var i int
	var workers []Worker
	input := make(chan Work)
	end := make(chan bool)
	collector := Collector{Work: input, End: end}

	for i < workerCount {
		i++
		log.Println("stating worker: ", i)
		worker := Worker{
			ID:            i,
			Channel:       make(chan Work),
			WorkerChannel: WorkerChannel,
			End:           make(chan bool),
		}

		worker.Start()
		workers = append(workers, worker)
	}

	go func() {
		for {
			select {
			case <-end:
				for _, w := range workers {
					w.Stop()
				}
				return
			case work := <-input:
				worker := <-WorkerChannel
				worker <- work
			}
		}
	}()

	return collector
}

const WORKER_COUNT = 5
const JOB_COUNT = 100

func main() {
	log.Println("starting application...")
	collector := StartDispatcher(WORKER_COUNT)

	for i, job := range CreateJobs(JOB_COUNT) {
		collector.Work <- Work{Job: job, ID: i}
	}
}
