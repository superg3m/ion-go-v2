package main

import (
	"fmt"
	"io/fs"
	"ion/go/v2/AssemblyAST"
	"ion/go/v2/AssemblyEmitter"
	"ion/go/v2/Codegen"
	"ion/go/v2/IR"
	"ion/go/v2/Lexer"
	"ion/go/v2/Parser"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func compileFile(inputPath string, outputAsm string) error {
	tokenStream := Lexer.GenerateTokenStream(inputPath)
	program := Parser.ParseProgram(tokenStream)
	ir := IR.GenerateIntermediateRepresentation(program)
	assembly := Codegen.GenerateAssemblyProgram(ir)
	finalStackOffset := 0
	assembly, finalStackOffset = Codegen.ReplacePseudoRegisters(assembly)
	assembly.FunctionDefinition.Instructions = slices.Insert(assembly.FunctionDefinition.Instructions, 0, AssemblyAST.NewStackAllocateInstruction(finalStackOffset))
	assembly = Codegen.ReplaceInvalidInstructions(assembly)

	return emitInstructions(outputAsm, assembly)
}

func emitInstructions(output string, assembly AssemblyAST.Program) error {
	emitter := &AssemblyEmitter.ATTx64Emitter{}
	instructions := emitter.EmitAssemblyProgram(assembly)

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range instructions {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func buildExecutable(asmPath string, exePath string) error {
	out, err := exec.Command(
		"gcc",
		asmPath,
		"-o",
		exePath,
	).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%v\n%s", err, string(out))
	}

	return nil
}

func runExecutable(path string) (int32, string, error) {
	cmd := exec.Command(path)

	out, err := cmd.CombinedOutput()

	exitCode := int32(0)

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = int32(exitErr.ExitCode())
		} else {
			return 0, "", err
		}
	}

	return exitCode, string(out), nil
}

func parseExpected(path string) int32 {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "// EXPECT:") {
			value := strings.TrimSpace(
				strings.TrimPrefix(line, "// EXPECT:"),
			)

			n, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			return int32(n)
		}
	}

	panic("missing // EXPECT:")
}

func runTest(path string) bool {
	name := filepath.Base(path)

	fmt.Println("====================================")
	fmt.Println("Running:", name)

	asmPath := "./Test/test.s"
	exePath := "./Test/test.exe"

	err := compileFile(path, asmPath)
	if err != nil {
		fmt.Println("COMPILER ERROR")
		fmt.Println(err)
		return false
	}

	err = buildExecutable(asmPath, exePath)
	if err != nil {
		fmt.Println("GCC ERROR")
		fmt.Println(err)
		return false
	}

	exitCode, stdout, err := runExecutable(exePath)
	if err != nil {
		fmt.Println("RUNTIME ERROR")
		fmt.Println(err)
		return false
	}

	expected := parseExpected(path)
	if exitCode != expected {
		fmt.Println("FAIL")
		fmt.Println("Expected:", expected)
		fmt.Println("Got:", exitCode)

		if stdout != "" {
			fmt.Println("Program Output:")
			fmt.Println(stdout)
		}

		return false
	}

	fmt.Println("PASS")

	if stdout != "" {
		fmt.Println("Program Output:")
		fmt.Println(stdout)
	}

	return true
}

func main() {
	dir := "./Test"

	total := 0
	passed := 0

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".c" {
			return nil
		}

		total++

		if runTest(path) {
			passed++
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("====================================")
	fmt.Printf("RESULT: %d/%d PASSED\n", passed, total)
}
