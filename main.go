package main

import (
	"fmt"
	"os"
	"strconv"

	pb "github.com/dot96gal/go-cel-sample/proto"
	"github.com/google/cel-go/cel"
)

func main() {
	be, err := os.ReadFile("./data/expression.txt")
	if err != nil {
		panic(err)
	}
	expr := string(be)

	bi, err := os.ReadFile("./data/input.txt")
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(string(bi))
	if err != nil {
		panic(err)
	}
	input := map[string]any{
		"payload": &pb.Payload{
			Num: int32(num),
		},
	}

	env, _ := cel.NewEnv(
		cel.Types(&pb.Payload{}),
		cel.Variable("payload",
			cel.ObjectType("sample.Payload"),
		),
	)

	ast, issues := env.Compile(expr)
	if issues != nil && issues.Err() != nil {
		panic(issues.Err())
	}

	prg, err := env.Program(ast)
	if err != nil {
		panic(err)
	}

	result, _, err := prg.Eval(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("input: %v, result: %v\n", input["payload"], result.Value().(bool))
}
