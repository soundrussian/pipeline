package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {
	fmt.Println("Hello, Go in Clouds! And CI, too!")
	fmt.Println("Deploying changes is a breathe!")

	data := [][]string{
		[]string{"1/1/2014", "Domain name", "1234", "$10.98"},
		[]string{"1/1/2014", "January Hosting", "2345", "$54.95"},
		[]string{"1/4/2014", "February Hosting", "3456", "$51.00"},
		[]string{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	table.SetFooter([]string{"", "", "Total", "$146.93"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}
