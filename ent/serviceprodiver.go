// Code generated by ent, DO NOT EDIT.

package ent

import (
	"database-concurrency/ent/serviceprodiver"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ServiceProdiver is the model entity for the ServiceProdiver schema.
type ServiceProdiver struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// VerdorRef holds the value of the "verdor_ref" field.
	VerdorRef string `json:"verdor_ref,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServiceProdiver) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case serviceprodiver.FieldName, serviceprodiver.FieldEmail, serviceprodiver.FieldPhone, serviceprodiver.FieldVerdorRef:
			values[i] = new(sql.NullString)
		case serviceprodiver.FieldCreatedAt, serviceprodiver.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case serviceprodiver.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ServiceProdiver", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServiceProdiver fields.
func (sp *ServiceProdiver) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case serviceprodiver.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sp.ID = *value
			}
		case serviceprodiver.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sp.Name = value.String
			}
		case serviceprodiver.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				sp.Email = value.String
			}
		case serviceprodiver.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				sp.Phone = value.String
			}
		case serviceprodiver.FieldVerdorRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field verdor_ref", values[i])
			} else if value.Valid {
				sp.VerdorRef = value.String
			}
		case serviceprodiver.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sp.CreatedAt = value.Time
			}
		case serviceprodiver.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sp.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ServiceProdiver.
// Note that you need to call ServiceProdiver.Unwrap() before calling this method if this ServiceProdiver
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *ServiceProdiver) Update() *ServiceProdiverUpdateOne {
	return (&ServiceProdiverClient{config: sp.config}).UpdateOne(sp)
}

// Unwrap unwraps the ServiceProdiver entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *ServiceProdiver) Unwrap() *ServiceProdiver {
	_tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ServiceProdiver is not a transactional entity")
	}
	sp.config.driver = _tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *ServiceProdiver) String() string {
	var builder strings.Builder
	builder.WriteString("ServiceProdiver(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sp.ID))
	builder.WriteString("name=")
	builder.WriteString(sp.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(sp.Email)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(sp.Phone)
	builder.WriteString(", ")
	builder.WriteString("verdor_ref=")
	builder.WriteString(sp.VerdorRef)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sp.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ServiceProdivers is a parsable slice of ServiceProdiver.
type ServiceProdivers []*ServiceProdiver

func (sp ServiceProdivers) config(cfg config) {
	for _i := range sp {
		sp[_i].config = cfg
	}
}
