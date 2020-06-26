package learnings

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	jobId int
	randomNumber int
}

type Result struct {
	job Job
	sum int
}

//at a time only 10 jobs will be held and also 10 results willbe held, until a worker takes the jobs and process
var jobs = make(chan Job, 5)
var results = make(chan Result, 5)

//This is the main function which finds sum
func findSum(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}


//each worker is created here and while the worker is created, he snatches the 10 jobs which are in hold by the jobs channel
func createWorker(wg *sync.WaitGroup){
	for job := range jobs {
		output := Result{job, findSum(job.randomNumber)}
		results <- output
	}
	wg.Done()
}


//creates a pool of workers mentioned above this is just for waiting for them in highlevel
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i< noOfWorkers; i++ {
		wg.Add(1)
		go createWorker(&wg)
	}
	wg.Wait()
	close(results)
}

//This is where we create jobs, with specifying job Id and random number
func allocate(noofjobs int){
	for i:= 0;i<noofjobs; i++ {
		randomNumber := rand.Intn(999)
		job := Job{
			jobId: i,
			randomNumber: randomNumber,
		}

		jobs <- job
	}
	fmt.Println("created jobs")
	close(jobs)
}

//This is where the result prints, this also holds 10 at a time
func result(done chan bool) {
	for result := range results {
		fmt.Printf("\njob id %d with random number %d is done and result is %d", result.job.jobId, result.job.randomNumber, result.sum)
	}
	done <- true
}

//Here we first create jobs based on the number, remember only 10 will be created and pushed and waits until a worker reads from the job pool
//then we use done channel just so the result printed is waited here
//Worker pool is created next and each worker is assigned 10 jobs(as job holds only 10 of them)
//the done channel is read and once finished the end time is printed
func LearnWorkerPools() {
	startTime := time.Now()
	noofJobs := 50
	fmt.Println("calling allocate")
	go allocate(noofJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<- done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("\ntotal time take is ", diff.Seconds(), "seconds")
}