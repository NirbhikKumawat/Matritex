package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func cube(initial [][]int) [][][]int {
	return [][][]int{
		{{initial[0][1], initial[0][0]}, {initial[1][0], initial[1][1]}},
		{{initial[0][0], initial[1][1]}, {initial[1][0], initial[0][1]}},
		{{initial[0][0], initial[0][1]}, {initial[1][1], initial[1][0]}},
		{{initial[1][0], initial[0][1]}, {initial[0][0], initial[1][1]}},
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
func printMatrixLatex(matrix [][]int, matrixType string) {
	fmt.Printf("\\begin{%s}", matrixType)
	fmt.Println()
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
			if j != len(matrix[i])-1 {
				fmt.Print("&")
			}
		}
		if i != len(matrix)-1 {
			fmt.Print("\\\\")
		}
		fmt.Println()
	}
	fmt.Printf("\\end{%s}", matrixType)
	fmt.Println()
}
func printNMatrix(cubes [][][]int) {
	for _, row := range cubes {
		printMatrix(row)
	}
}
func printNMatrixLatexSingleRow(cubes [][][]int, matrixType string) {
	for count := 0; count < len(cubes)-1; count++ {
		printMatrixLatex(cubes[count], matrixType)
		fmt.Println("\\quad")
	}
	printMatrixLatex(cubes[len(cubes)-1], matrixType)
}
func printMatrixDistanceSingleRow(matrix [][]int, matrixType string) {
	printMatrixLatex(matrix, matrixType)
	printNMatrixLatexSingleRow(cube(matrix), matrixType)
}
func main() {
	var rootCmd = &cobra.Command{
		Use:   "matritex",
		Short: "Output LaTeX Matrix",
		Long:  "A tool useful for converting a normal matrix into LaTeX format",
		Args:  cobra.MinimumNArgs(1),
		Run:   runMatrice,
	}
	//rootCmd.Flags().BoolVarP(&nolatex, "no-latex", "l", false, "no Latex")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func runMatrice(cmd *cobra.Command, args []string) {
	matrixJSON := args[0]
	var matrix [][]int
	err := json.Unmarshal([]byte(matrixJSON), &matrix)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing matrix JSON: %v\n", err)
		os.Exit(1)
	}
	printMatrixDistanceSingleRow(matrix, "pmatrix")
}
