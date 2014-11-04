package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*===============================================================
 * Functions to manipulate a "field" of cells --- the main data
 * that must be managed by this program.
 *==============================================================*/

// The data stored in a single cell of a field
type Cell struct {
	kind  string
	score float64
	prekind string
}

// createField should create a new field of the ysize rows and xsize columns,
// so that field[r][c] gives the Cell at position (r,c).
func createField(rsize, csize int) [][]Cell {
	f := make([][]Cell, rsize)
	for i := range f {
		f[i] = make([]Cell, csize)
	}
	return f
}

// inField returns true iff (row,col) is a valid cell in the field
func inField(field [][]Cell, row, col int) bool {
	return row >= 0 && row < len(field) && col >= 0 && col < len(field[0])
}

// readFieldFromFile should open the given file and read the initial
// values for the field. The first line of the file will contain
// two space-separated integers saying how many rows and columns
// the field should have:
//    10 15
// each subsequent line will consist of a string of Cs and Ds, which
// are the initial strategies for the cells:
//    CCCCCCDDDCCCCCC
//
// If there is ever an error reading, this function should cause the
// program to quit immediately.

/*
 *First, reads the contents from the file to a string array lines using append method.
 *Then,loop from the second line of the lines(the first line is the size of content)
 *and put the value of lines to mycell(which type is [][]Cell) 
 */
func readFieldFromFile(filename string) [][]Cell {
	in,err := os.Open(filename)
	if err != nil{
		fmt.Println("Error: something went wrong opening the file")
		fmt.Println("Probably you gave the wrong filename.")
		os.Exit(3)
	}

    
    var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(in)
	
	for scanner.Scan(){
	 	lines = append(lines, scanner.Text())
	 }

    
	var mycell [][]Cell = createField(len(lines)-1,len(lines[1]))
	// c[0][0].kind = "a"
	// fmt.Println(c[0][0].kind())
    for i:=1; i<len(lines); i++ {
         
    	 for j:=0; j<len(lines[1]); j++ {
    	 	//fmt.Println(len(lines[1]))
    	 	// fmt.Println(string(lines[i][j]))
            mycell[i-1][j].kind = string(lines[i][j])
            // fmt.Println(mycell[i-1][j].kind)
    	}	
    }

	
	if scanner.Err() != nil {
		fmt.Println("Sorry: there was some kind of error during the file reading")
		os.Exit(3)
	}
    
    in.Close()
    return mycell
    // WRITE YOUR CODE HERE
}

// drawField should draw a representation of the field on a canvas and save the
// canvas to a PNG file with a name given by the parameter filename.  Each cell
// in the field should be a 5-by-5 square, and cells of the "D" kind should be
// drawn red and cells of the "C" kind should be drawn blue.

/*
*According to the string in field [][]Cell
*Draw the picture "filename" in different colors
*/
func drawField(field [][]Cell, filename string) {
    
	cellpic := CreateNewCanvas(5*len(field),5*len(field[0]))
	var x1, x2, y1, y2 float64

	for i:=0; i<len(field); i++{
        for j:=0; j<len(field[0]); j++{
        	x1, y1 = float64(5*i), float64(5*j)
            x2, y2 = float64(5*i+5), float64(5*j+5)
        	if field[i][j].kind=="C" && field[i][j].prekind=="C"{
				cellpic.SetStrokeColor(MakeColor(0,0,255))
                cellpic.SetFillColor(MakeColor(0,0,255))
			}
			if field[i][j].kind=="D" && field[i][j].prekind=="C"{
				cellpic.SetStrokeColor(MakeColor(255,255,0))
                cellpic.SetFillColor(MakeColor(255,255,0))
			}
			if field[i][j].kind=="D" && field[i][j].prekind=="D"{
				cellpic.SetStrokeColor(MakeColor(255,0,0))
                cellpic.SetFillColor(MakeColor(255,0,0))
			}
			if field[i][j].kind=="C" && field[i][j].prekind=="D"{
				cellpic.SetStrokeColor(MakeColor(0,255,0))
                cellpic.SetFillColor(MakeColor(0,255,0))
			}
			cellpic.MoveTo(x1,y1)
            cellpic.LineTo(x2,y1)
            cellpic.LineTo(x2,y2)
            cellpic.LineTo(x1,y2)
            cellpic.LineTo(x1,y1)
            cellpic.FillStroke()
        }		
				
		}
		cellpic.SaveToPNG(filename)
    // WRITE YOUR CODE HERE
}

/*===============================================================
 * Functions to simulate the spatial games
 *==============================================================*/

// play a game between a cell of type "me" and a cell of type "them" (both me
// and them should be either "C" or "D"). This returns the reward that "me"
// gets when playing against them.
func gameBetween(me, them string, b float64) float64 {
	if me == "C" && them == "C" {
		return 1
	} else if me == "C" && them == "D" {
		return 0
	} else if me == "D" && them == "C" {
		return b
	} else if me == "D" && them == "D" {
		return 0
	} else {
		fmt.Println("type ==", me, them)
		panic("This shouldn't happen")
		return 0
	}
}

// updateScores goes through every cell, and plays the Prisoner's dilema game
// with each of it's in-field nieghbors (including itself). It updates the
// score of each cell to be the sum of that cell's winnings from the game.

/*
*Sum the scores of eight neighbors and the cell itself(9 in total).
*Define the score of the cell to be this sum.
*/
func updateScores(field [][]Cell, b float64) {
	var sum float64
	for i:=0; i<len(field); i++ {
		for j:=0; j<len(field[0]); j++ {
			sum=0
		    for k:=i-1; k<=i+1; k=k+1 {
		    	for g:=j-1; g<=j+1; g=g+1{
                    if inField(field, k, g){
                    	sum = sum + gameBetween(field[i][j].kind, field[k][g].kind, b)
                    }
		    	}
		    }
		    field[i][j].score = sum
		    }
	}
    // WRITE YOUR CODE HERE
}		
		 //    sum=0
			// if inField(field, i-1, j+1){
			// 	sum = sum + gameBetween(field[i][j].kind, field[i-1][j+1].kind, b)
			// }
			// if inField(field, i-1, j){
			// 	sum = sum + gameBetween(field[i][j].kind, field[i-1][j].kind, b)
			// }
			// if inField(field, i-1, j-1){
			// 	sum = sum + gameBetween(field[i][j].kind, field[i-1][j-1].kind, b)
			// }
			// if inField(field, i, j+1){
			// 	sum = sum + gameBetween(field[i][j].kind, field[i][j+1].kind, b)
			// }
   //          if inField(field, i, j){
			// 	sum = sum + gameBetween(field[i][j].kind, field[i][j].kind, b)
			// }
			// if inField(field, i, j-1){
			// 	sum = sum + gameBetween(field[i][j].kind, field[i][j-1].kind, b)
			// }
			// if inField(field, i+1, j+1){
			// 	sum = sum + gameBetween(field[i][j].kind, field[i+1][j+1].kind, b)
			// }
			// if inField(field, i+1, j){
			// 	sum = sum + gameBetween(field[i][j].kind, field[i+1][j].kind, b)
			// }
			// if inField(field, i+1, j-1){
			// 	sum = sum + gameBetween(field[i][j].kind, field[i+1][j-1].kind, b)
			// }
			
		

// updateStrategies create a new field by going through every cell (r,c), and
// looking at each of the cells in its neighborhood (including itself) and the
// setting the kind of cell (r,c) in the new field to be the kind of the
// neighbor with the largest score

/*
*Compare the scores of eight neighbors and itself.
*Change strategy to the strategy of the biggest one.
*/

func updateStrategies(field [][]Cell) [][]Cell {
	var max float64
	var mark string
	var newcell [][]Cell = createField(len(field),len(field[0]))
	for i:=0; i<len(field); i++ {
		for j:=0; j<len(field[0]); j++ {
			max=-1
			for k:=i-1; k<=i+1; k=k+1 {
		    	for g:=j-1; g<=j+1; g=g+1{
                    if (inField(field, k, g)) && (max<field[k][g].score) {
				       max=field[k][g].score
				       mark=field[k][g].kind
			        }
		    	}
		    }
		    newcell[i][j].prekind=field[i][j].kind
		    newcell[i][j].kind=mark
		    }  
	}
    // WRITE YOUR CODE HERE
	return newcell // This is included only so this template will compile
}
            
			// if (inField(field, i-1, j)) && (max<field[i-1][j].score) {
			// 	max=field[i-1][j].score
			// 	mark=field[i-1][j].kind
			// }
			// if (inField(field, i-1, j+1)) && (max<field[i-1][j+1].score) {
			// 	max=field[i-1][j+1].score
			// 	mark=field[i-1][j+1].kind
			// }
			// if (inField(field, i, j-1)) && (max<field[i][j-1].score) {
			// 	max=field[i][j-1].score
			// 	mark=field[i][j-1].kind
			// }
			// if (inField(field, i, j)) && (max<field[i][j].score) {
			// 	max=field[i][j].score
			// 	mark=field[i][j].kind
			// }
			// if (inField(field, i, j+1)) && (max<field[i][j+1].score) {
			// 	max=field[i][j+1].score
			// 	mark=field[i][j+1].kind
			// }
			// if (inField(field, i+1, j-1)) && (max<field[i+1][j-1].score) {
			// 	max=field[i+1][j-1].score
			// 	mark=field[i+1][j-1].kind
			// }
			// if (inField(field, i+1, j)) && (max<field[i+1][j].score) {
			// 	max=field[i+1][j].score
			// 	mark=field[i+1][j].kind
			// }
			// if (inField(field, i+1, j+1)) && (max<field[i+1][j+1].score) {
			// 	max=field[i+1][j+1].score
			// 	mark=field[i+1][j+1].kind
			// }
			
			// fmt.Print(mark)
			
			
		

// evolve takes an intial field and evolves it for nsteps according to the game
// rule. At each step, it should call "updateScores()" and the updateStrategies
func evolve(field [][]Cell, nsteps int, b float64) [][]Cell {
	for i := 0; i < nsteps; i++ {
		updateScores(field, b)
		field = updateStrategies(field)
	}
	return field
}

// Implements a Spatial Games version of prisoner's dilemma. The command-line
// usage is:
//     ./spatial field_file b nsteps
// where 'field_file' is the file continaing the initial arrangment of cells, b
// is the reward for defecting against a cooperator, and nsteps is the number
// of rounds to update stategies.
//
func main() {
	// parse the command line
	if len(os.Args) != 4 {
		fmt.Println("Error: should spatial field_file b nsteps")
		return
	}

	fieldFile := os.Args[1]

	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil || b <= 0 {
		fmt.Println("Error: bad b parameter.")
		return
	}

	nsteps, err := strconv.Atoi(os.Args[3])
	if err != nil || nsteps < 0 {
		fmt.Println("Error: bad number of steps.")
		return
	}

    // read the field
	field := readFieldFromFile(fieldFile)
    fmt.Println("Field dimensions are:", len(field), "by", len(field[0]))

    // evolve the field for nsteps and write it as a PNG
	field = evolve(field, nsteps, b)
	drawField(field, "Prisoners.png")
}
