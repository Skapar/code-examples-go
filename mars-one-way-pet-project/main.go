package main

func main() {
	response := utils()

    printHeader(response.Columns)
    printTickets(response.Tickets)
}