package main

import (
	"math/rand"
)

func utils() ResponseType {
    days := generateDays(count)
    prices := generatePrices(days)
    ticketData := TicketData{
        SpaceLines: spaceLines,
        TripTypes:  tripTypes,
        Days:       days,
        Prices:     prices,
    }
    tickets := generateTickets(ticketData)

	return ResponseType{
        BaseResponse: BaseResponse{
            Count:   count,
            Columns: columns,
        },
        Tickets: tickets,
    }
}

func generateDays(count int) []int {
    days := make([]int, count)
    for i := 0; i < count; i++ {
        days[i] = rand.Intn(30)
    }
    return days
}

func generatePrices(days []int) []int {
    prices := make([]int, len(days))
    for i, day := range days {
        prices[i] = priceMultip * day
    }
    return prices
}

func generateTickets(data TicketData) []Ticket {
    tickets := make([]Ticket, count)
    for i := 0; i < count; i++ {
        tickets[i] = Ticket{
            SpaceLine: data.SpaceLines[i%len(data.SpaceLines)],
            Days:      data.Days[i],
            TripType:  data.TripTypes[i%len(data.TripTypes)],
            Price:     data.Prices[i],
        }
    }
    return tickets
}