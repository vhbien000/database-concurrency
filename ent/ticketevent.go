// Code generated by ent, DO NOT EDIT.

package ent

import (
	"database-concurrency/ent/ticket"
	"database-concurrency/ent/ticketevent"
	"database-concurrency/ent/user"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// TicketEvent is the model entity for the TicketEvent schema.
type TicketEvent struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// TicketID holds the value of the "ticket_id" field.
	TicketID uuid.UUID `json:"ticket_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Metadada holds the value of the "metadada" field.
	Metadada map[string]interface{} `json:"metadada,omitempty"`
	// Versions holds the value of the "versions" field.
	Versions string `json:"versions,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TicketEventQuery when eager-loading is set.
	Edges TicketEventEdges `json:"edges"`
}

// TicketEventEdges holds the relations/edges for other nodes in the graph.
type TicketEventEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Ticket holds the value of the ticket edge.
	Ticket *Ticket `json:"ticket,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TicketEventEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// TicketOrErr returns the Ticket value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TicketEventEdges) TicketOrErr() (*Ticket, error) {
	if e.loadedTypes[1] {
		if e.Ticket == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: ticket.Label}
		}
		return e.Ticket, nil
	}
	return nil, &NotLoadedError{edge: "ticket"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TicketEvent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ticketevent.FieldMetadada:
			values[i] = new([]byte)
		case ticketevent.FieldType, ticketevent.FieldVersions:
			values[i] = new(sql.NullString)
		case ticketevent.FieldCreatedAt, ticketevent.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case ticketevent.FieldID, ticketevent.FieldTicketID, ticketevent.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TicketEvent", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TicketEvent fields.
func (te *TicketEvent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ticketevent.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				te.ID = *value
			}
		case ticketevent.FieldTicketID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ticket_id", values[i])
			} else if value != nil {
				te.TicketID = *value
			}
		case ticketevent.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				te.UserID = *value
			}
		case ticketevent.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				te.Type = value.String
			}
		case ticketevent.FieldMetadada:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadada", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &te.Metadada); err != nil {
					return fmt.Errorf("unmarshal field metadada: %w", err)
				}
			}
		case ticketevent.FieldVersions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field versions", values[i])
			} else if value.Valid {
				te.Versions = value.String
			}
		case ticketevent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				te.CreatedAt = value.Time
			}
		case ticketevent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				te.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the TicketEvent entity.
func (te *TicketEvent) QueryUser() *UserQuery {
	return (&TicketEventClient{config: te.config}).QueryUser(te)
}

// QueryTicket queries the "ticket" edge of the TicketEvent entity.
func (te *TicketEvent) QueryTicket() *TicketQuery {
	return (&TicketEventClient{config: te.config}).QueryTicket(te)
}

// Update returns a builder for updating this TicketEvent.
// Note that you need to call TicketEvent.Unwrap() before calling this method if this TicketEvent
// was returned from a transaction, and the transaction was committed or rolled back.
func (te *TicketEvent) Update() *TicketEventUpdateOne {
	return (&TicketEventClient{config: te.config}).UpdateOne(te)
}

// Unwrap unwraps the TicketEvent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (te *TicketEvent) Unwrap() *TicketEvent {
	_tx, ok := te.config.driver.(*txDriver)
	if !ok {
		panic("ent: TicketEvent is not a transactional entity")
	}
	te.config.driver = _tx.drv
	return te
}

// String implements the fmt.Stringer.
func (te *TicketEvent) String() string {
	var builder strings.Builder
	builder.WriteString("TicketEvent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", te.ID))
	builder.WriteString("ticket_id=")
	builder.WriteString(fmt.Sprintf("%v", te.TicketID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", te.UserID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(te.Type)
	builder.WriteString(", ")
	builder.WriteString("metadada=")
	builder.WriteString(fmt.Sprintf("%v", te.Metadada))
	builder.WriteString(", ")
	builder.WriteString("versions=")
	builder.WriteString(te.Versions)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(te.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(te.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TicketEvents is a parsable slice of TicketEvent.
type TicketEvents []*TicketEvent

func (te TicketEvents) config(cfg config) {
	for _i := range te {
		te[_i].config = cfg
	}
}
