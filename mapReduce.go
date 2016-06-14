package main

//The code below is from http://marcio.io/

//Channel that collects all the results from mapper task
type Collector chan chan interface{}

// MapperFunc is a function that performs the mapping part of the MapReduce job
type MapperFunc func(interface{}, chan interface{})

// ReducerFunc is a function that performs the reduce part of the MapReduce job
type ReducerFunc func(chan interface{}, chan interface{})

func mapperDispatcher(mapper MapperFunc, input chan interface{}, collector Collector) {
	for item := range input {
		taskOutput := make(chan interface{})
		go mapper(item, taskOutput)
		collector <- taskOutput
	}
	close(collector)
}

func reducerDispatcher(collector Collector, reducerInput chan interface{}) {
	for output := range collector {
		reducerInput <- <-output
	}
	close(reducerInput)
}

func mapReduce(mapper MapperFunc, reducer ReducerFunc, input chan interface{}) chan interface{} {

	reducerInput := make(chan interface{})
	reducerOutput := make(chan interface{})
	collector := make(Collector, MaxWorkers)

	go mapperDispatcher(mapper, input, collector)
	go reducerDispatcher(collector, reducerInput)
	go reducer(reducerInput, reducerOutput)

	return reducerOutput
}
