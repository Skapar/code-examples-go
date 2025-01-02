package main

type BaseResponse struct {
    Count   int
    Columns []string
}

type ResponseType struct {
    BaseResponse
    Tickets []Ticket
}

type Ticket struct {
    SpaceLine string
    Days      int
    TripType  string
    Price     int
}

type TicketData struct {
    SpaceLines []string
    TripTypes  []string
    Days       []int
    Prices     []int
}