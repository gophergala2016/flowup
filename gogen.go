package gogen

import "sync"

var (
	// Models is public static set of models exposed
	// by the Gogen, so generators can use it. This
	// set contains every model that was added either
	// manually or by Define function
	Models []*Model

	// Pipes is set of pipelines that should be run when
	// generate is called
	Pipes []Pipeline
)

// Define will store the defined model for the use in
// the generators.
func Define(what interface{}) {
	switch val := what.(type) {
	case *Model:
		Models = append(Models, val)
	default:
		panic("Type passed to define not recognized")
	}
}

// Pipe will register new pipe that will be run
// in parallel
func Pipe(gens ...Generator) {
	pipe := Pipeline{}
	for _, gen := range gens {
		pipe.Add(gen)
	}

	Pipes = append(Pipes, pipe)
}

// Generate will startup a
func Generate() {
	wg := sync.WaitGroup{}

	for _, pipe := range Pipes {
		wg.Add(1)
		go func(pipe Pipeline) {
			for _, gen := range pipe.generators {
				err := gen.Generate()
				// TODO: make this not panic, but return the error
				if err != nil {
					panic(err)
				}
			}

			wg.Done()
		}(pipe)
	}

	wg.Wait()
}