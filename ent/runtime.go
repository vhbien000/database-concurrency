// Code generated by ent, DO NOT EDIT.

package ent

import (
	"database-concurrency/ent/schema"
	"database-concurrency/ent/serviceprodiver"
	"database-concurrency/ent/ticket"
	"database-concurrency/ent/ticketevent"
	"database-concurrency/ent/transaction"
	"database-concurrency/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	serviceprodiverFields := schema.ServiceProdiver{}.Fields()
	_ = serviceprodiverFields
	// serviceprodiverDescCreatedAt is the schema descriptor for created_at field.
	serviceprodiverDescCreatedAt := serviceprodiverFields[5].Descriptor()
	// serviceprodiver.DefaultCreatedAt holds the default value on creation for the created_at field.
	serviceprodiver.DefaultCreatedAt = serviceprodiverDescCreatedAt.Default.(func() time.Time)
	// serviceprodiverDescUpdatedAt is the schema descriptor for updated_at field.
	serviceprodiverDescUpdatedAt := serviceprodiverFields[6].Descriptor()
	// serviceprodiver.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	serviceprodiver.DefaultUpdatedAt = serviceprodiverDescUpdatedAt.Default.(func() time.Time)
	// serviceprodiver.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	serviceprodiver.UpdateDefaultUpdatedAt = serviceprodiverDescUpdatedAt.UpdateDefault.(func() time.Time)
	// serviceprodiverDescID is the schema descriptor for id field.
	serviceprodiverDescID := serviceprodiverFields[0].Descriptor()
	// serviceprodiver.DefaultID holds the default value on creation for the id field.
	serviceprodiver.DefaultID = serviceprodiverDescID.Default.(func() uuid.UUID)
	ticketFields := schema.Ticket{}.Fields()
	_ = ticketFields
	// ticketDescCreatedAt is the schema descriptor for created_at field.
	ticketDescCreatedAt := ticketFields[6].Descriptor()
	// ticket.DefaultCreatedAt holds the default value on creation for the created_at field.
	ticket.DefaultCreatedAt = ticketDescCreatedAt.Default.(func() time.Time)
	// ticketDescUpdatedAt is the schema descriptor for updated_at field.
	ticketDescUpdatedAt := ticketFields[7].Descriptor()
	// ticket.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ticket.DefaultUpdatedAt = ticketDescUpdatedAt.Default.(func() time.Time)
	// ticket.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ticket.UpdateDefaultUpdatedAt = ticketDescUpdatedAt.UpdateDefault.(func() time.Time)
	// ticketDescID is the schema descriptor for id field.
	ticketDescID := ticketFields[0].Descriptor()
	// ticket.DefaultID holds the default value on creation for the id field.
	ticket.DefaultID = ticketDescID.Default.(func() uuid.UUID)
	ticketeventFields := schema.TicketEvent{}.Fields()
	_ = ticketeventFields
	// ticketeventDescCreatedAt is the schema descriptor for created_at field.
	ticketeventDescCreatedAt := ticketeventFields[6].Descriptor()
	// ticketevent.DefaultCreatedAt holds the default value on creation for the created_at field.
	ticketevent.DefaultCreatedAt = ticketeventDescCreatedAt.Default.(func() time.Time)
	// ticketeventDescUpdatedAt is the schema descriptor for updated_at field.
	ticketeventDescUpdatedAt := ticketeventFields[7].Descriptor()
	// ticketevent.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ticketevent.DefaultUpdatedAt = ticketeventDescUpdatedAt.Default.(func() time.Time)
	// ticketevent.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ticketevent.UpdateDefaultUpdatedAt = ticketeventDescUpdatedAt.UpdateDefault.(func() time.Time)
	// ticketeventDescID is the schema descriptor for id field.
	ticketeventDescID := ticketeventFields[0].Descriptor()
	// ticketevent.DefaultID holds the default value on creation for the id field.
	ticketevent.DefaultID = ticketeventDescID.Default.(func() uuid.UUID)
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescCreatedAt is the schema descriptor for created_at field.
	transactionDescCreatedAt := transactionFields[3].Descriptor()
	// transaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	transaction.DefaultCreatedAt = transactionDescCreatedAt.Default.(func() time.Time)
	// transactionDescID is the schema descriptor for id field.
	transactionDescID := transactionFields[0].Descriptor()
	// transaction.DefaultID holds the default value on creation for the id field.
	transaction.DefaultID = transactionDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
