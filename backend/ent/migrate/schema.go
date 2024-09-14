// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "map_url", Type: field.TypeString},
		{Name: "qr_code_url", Type: field.TypeString, Nullable: true},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_events", Type: field.TypeInt, Nullable: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_users_events",
				Columns:    []*schema.Column{EventsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EventAdminsColumns holds the columns for the "event_admins" table.
	EventAdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "event_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "event_event_admins", Type: field.TypeInt, Nullable: true},
		{Name: "user_event_admins", Type: field.TypeInt, Nullable: true},
	}
	// EventAdminsTable holds the schema information for the "event_admins" table.
	EventAdminsTable = &schema.Table{
		Name:       "event_admins",
		Columns:    EventAdminsColumns,
		PrimaryKey: []*schema.Column{EventAdminsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "event_admins_events_event_admins",
				Columns:    []*schema.Column{EventAdminsColumns[4]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "event_admins_users_event_admins",
				Columns:    []*schema.Column{EventAdminsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "content", Type: field.TypeString},
		{Name: "location_x", Type: field.TypeFloat64},
		{Name: "location_y", Type: field.TypeFloat64},
		{Name: "video_url", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PostImagesColumns holds the columns for the "post_images" table.
	PostImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "post_id", Type: field.TypeInt},
		{Name: "image_url", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "post_images", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PostImagesTable holds the schema information for the "post_images" table.
	PostImagesTable = &schema.Table{
		Name:       "post_images",
		Columns:    PostImagesColumns,
		PrimaryKey: []*schema.Column{PostImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_images_posts_images",
				Columns:    []*schema.Column{PostImagesColumns[4]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt, Unique: true},
		{Name: "username", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "hashed_password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// EventPostsColumns holds the columns for the "event_posts" table.
	EventPostsColumns = []*schema.Column{
		{Name: "event_id", Type: field.TypeInt},
		{Name: "post_id", Type: field.TypeInt},
	}
	// EventPostsTable holds the schema information for the "event_posts" table.
	EventPostsTable = &schema.Table{
		Name:       "event_posts",
		Columns:    EventPostsColumns,
		PrimaryKey: []*schema.Column{EventPostsColumns[0], EventPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "event_posts_event_id",
				Columns:    []*schema.Column{EventPostsColumns[0]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "event_posts_post_id",
				Columns:    []*schema.Column{EventPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		EventAdminsTable,
		PostsTable,
		PostImagesTable,
		UsersTable,
		EventPostsTable,
	}
)

func init() {
	EventsTable.ForeignKeys[0].RefTable = UsersTable
	EventAdminsTable.ForeignKeys[0].RefTable = EventsTable
	EventAdminsTable.ForeignKeys[1].RefTable = UsersTable
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	PostImagesTable.ForeignKeys[0].RefTable = PostsTable
	EventPostsTable.ForeignKeys[0].RefTable = EventsTable
	EventPostsTable.ForeignKeys[1].RefTable = PostsTable
}
