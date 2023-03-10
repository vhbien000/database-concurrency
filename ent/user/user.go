// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTickets holds the string denoting the tickets edge name in mutations.
	EdgeTickets = "tickets"
	// EdgeTicketEvents holds the string denoting the ticket_events edge name in mutations.
	EdgeTicketEvents = "ticket_events"
	// Table holds the table name of the user in the database.
	Table = "users"
	// TicketsTable is the table that holds the tickets relation/edge.
	TicketsTable = "tickets"
	// TicketsInverseTable is the table name for the Ticket entity.
	// It exists in this package in order to avoid circular dependency with the "ticket" package.
	TicketsInverseTable = "tickets"
	// TicketsColumn is the table column denoting the tickets relation/edge.
	TicketsColumn = "user_id"
	// TicketEventsTable is the table that holds the ticket_events relation/edge.
	TicketEventsTable = "ticket_events"
	// TicketEventsInverseTable is the table name for the TicketEvent entity.
	// It exists in this package in order to avoid circular dependency with the "ticketevent" package.
	TicketEventsInverseTable = "ticket_events"
	// TicketEventsColumn is the table column denoting the ticket_events relation/edge.
	TicketEventsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPhone,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
