// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CreateTicketLogsColumns holds the columns for the "create_ticket_logs" table.
	CreateTicketLogsColumns = []*schema.Column{
		{Name: "ticket_id", Type: field.TypeUUID},
		{Name: "unique_id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CreateTicketLogsTable holds the schema information for the "create_ticket_logs" table.
	CreateTicketLogsTable = &schema.Table{
		Name:       "create_ticket_logs",
		Columns:    CreateTicketLogsColumns,
		PrimaryKey: []*schema.Column{CreateTicketLogsColumns[0]},
	}
	// ServiceProdiversColumns holds the columns for the "service_prodivers" table.
	ServiceProdiversColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "verdor_ref", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ServiceProdiversTable holds the schema information for the "service_prodivers" table.
	ServiceProdiversTable = &schema.Table{
		Name:       "service_prodivers",
		Columns:    ServiceProdiversColumns,
		PrimaryKey: []*schema.Column{ServiceProdiversColumns[0]},
	}
	// TicketsColumns holds the columns for the "tickets" table.
	TicketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "status", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "versions", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_event_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// TicketsTable holds the schema information for the "tickets" table.
	TicketsTable = &schema.Table{
		Name:       "tickets",
		Columns:    TicketsColumns,
		PrimaryKey: []*schema.Column{TicketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tickets_ticket_events_last_event",
				Columns:    []*schema.Column{TicketsColumns[6]},
				RefColumns: []*schema.Column{TicketEventsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "tickets_users_tickets",
				Columns:    []*schema.Column{TicketsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TicketEventsColumns holds the columns for the "ticket_events" table.
	TicketEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "versions", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "ticket_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// TicketEventsTable holds the schema information for the "ticket_events" table.
	TicketEventsTable = &schema.Table{
		Name:       "ticket_events",
		Columns:    TicketEventsColumns,
		PrimaryKey: []*schema.Column{TicketEventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ticket_events_tickets_ticket_events",
				Columns:    []*schema.Column{TicketEventsColumns[6]},
				RefColumns: []*schema.Column{TicketsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "ticket_events_users_ticket_events",
				Columns:    []*schema.Column{TicketEventsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "hash", Type: field.TypeString, Unique: true},
		{Name: "time", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CreateTicketLogsTable,
		ServiceProdiversTable,
		TicketsTable,
		TicketEventsTable,
		TransactionsTable,
		UsersTable,
	}
)

func init() {
	TicketsTable.ForeignKeys[0].RefTable = TicketEventsTable
	TicketsTable.ForeignKeys[1].RefTable = UsersTable
	TicketEventsTable.ForeignKeys[0].RefTable = TicketsTable
	TicketEventsTable.ForeignKeys[1].RefTable = UsersTable
}
