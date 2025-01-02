package main

import "fmt"

func printHeader(columns []string) {
    fmt.Printf("%-20s %-10s %-10s %10s\n", columns[0], columns[1], columns[2], columns[3])
    fmt.Println("=====================================================")
}

func printTickets(tickets []Ticket) {
    for _, ticket := range tickets {
        fmt.Printf("%-20v %-10v %-10v $%10d\n", ticket.SpaceLine, ticket.Days, ticket.TripType, ticket.Price)
    }
}