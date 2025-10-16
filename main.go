package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//var nolatex bool

func cube(initial [2][2]int) [4][2][2]int {
	return [4][2][2]int{
		[2][2]int{{initial[0][1], initial[0][0]}, {initial[1][0], initial[1][1]}},
		[2][2]int{{initial[0][0], initial[1][1]}, {initial[1][0], initial[0][1]}},
		[2][2]int{{initial[0][0], initial[0][1]}, {initial[1][1], initial[1][0]}},
		[2][2]int{{initial[1][0], initial[0][1]}, {initial[0][0], initial[1][1]}},
	}
}
func printMatrice(matrice [2][2]int) {
	for row := 0; row < 2; row++ {
		for col := 0; col < 2; col++ {
			fmt.Print(matrice[row][col])
		}
		fmt.Println()
	}
	fmt.Println()
}
func printMatriceLatex(matrice [2][2]int) {
	fmt.Println("\\begin{bmatrix}")
	fmt.Println(matrice[0][0], "&", matrice[0][1], "\\\\")
	fmt.Println(matrice[1][0], "&", matrice[1][1])
	fmt.Println("\\end{bmatrix}")
}
func print4MatriceLatex(cubes [4][2][2]int) {
	for count := 0; count < 3; count++ {
		printMatriceLatex(cubes[count])
		fmt.Println("\\quad")
	}
	printMatriceLatex(cubes[3])
}

func printMatriceLatexDistance(matrice [2][2]int) {
	printMatriceLatex(matrice)
	print4MatriceLatex(cube(matrice))
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
	var matrix [2][2]int
	err := json.Unmarshal([]byte(matrixJSON), &matrix)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing matrix JSON: %v\n", err)
		os.Exit(1)
	}
	printMatriceLatexDistance(matrix)
}
