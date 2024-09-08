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
				Columns:    []*schema.Column{EventsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "comment", Type: field.TypeString},
		{Name: "location_x", Type: field.TypeFloat64},
		{Name: "location_y", Type: field.TypeFloat64},
		{Name: "video_url", Type: field.TypeString, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
	}
	// PostImagesColumns holds the columns for the "post_images" table.
	PostImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "image_url", Type: field.TypeString},
		{Name: "post_images", Type: field.TypeInt, Nullable: true},
	}
	// PostImagesTable holds the schema information for the "post_images" table.
	PostImagesTable = &schema.Table{
		Name:       "post_images",
		Columns:    PostImagesColumns,
		PrimaryKey: []*schema.Column{PostImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_images_posts_images",
				Columns:    []*schema.Column{PostImagesColumns[2]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReportsColumns holds the columns for the "reports" table.
	ReportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "reason", Type: field.TypeString, Size: 2147483647},
	}
	// ReportsTable holds the schema information for the "reports" table.
	ReportsTable = &schema.Table{
		Name:       "reports",
		Columns:    ReportsColumns,
		PrimaryKey: []*schema.Column{ReportsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
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
	// PostReportsColumns holds the columns for the "post_reports" table.
	PostReportsColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeInt},
		{Name: "report_id", Type: field.TypeInt},
	}
	// PostReportsTable holds the schema information for the "post_reports" table.
	PostReportsTable = &schema.Table{
		Name:       "post_reports",
		Columns:    PostReportsColumns,
		PrimaryKey: []*schema.Column{PostReportsColumns[0], PostReportsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_reports_post_id",
				Columns:    []*schema.Column{PostReportsColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_reports_report_id",
				Columns:    []*schema.Column{PostReportsColumns[1]},
				RefColumns: []*schema.Column{ReportsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserPostsColumns holds the columns for the "user_posts" table.
	UserPostsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "post_id", Type: field.TypeInt},
	}
	// UserPostsTable holds the schema information for the "user_posts" table.
	UserPostsTable = &schema.Table{
		Name:       "user_posts",
		Columns:    UserPostsColumns,
		PrimaryKey: []*schema.Column{UserPostsColumns[0], UserPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_posts_user_id",
				Columns:    []*schema.Column{UserPostsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_posts_post_id",
				Columns:    []*schema.Column{UserPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserReportsColumns holds the columns for the "user_reports" table.
	UserReportsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "report_id", Type: field.TypeInt},
	}
	// UserReportsTable holds the schema information for the "user_reports" table.
	UserReportsTable = &schema.Table{
		Name:       "user_reports",
		Columns:    UserReportsColumns,
		PrimaryKey: []*schema.Column{UserReportsColumns[0], UserReportsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_reports_user_id",
				Columns:    []*schema.Column{UserReportsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_reports_report_id",
				Columns:    []*schema.Column{UserReportsColumns[1]},
				RefColumns: []*schema.Column{ReportsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		PostsTable,
		PostImagesTable,
		ReportsTable,
		UsersTable,
		EventPostsTable,
		PostReportsTable,
		UserPostsTable,
		UserReportsTable,
	}
)

func init() {
	EventsTable.ForeignKeys[0].RefTable = UsersTable
	PostImagesTable.ForeignKeys[0].RefTable = PostsTable
	EventPostsTable.ForeignKeys[0].RefTable = EventsTable
	EventPostsTable.ForeignKeys[1].RefTable = PostsTable
	PostReportsTable.ForeignKeys[0].RefTable = PostsTable
	PostReportsTable.ForeignKeys[1].RefTable = ReportsTable
	UserPostsTable.ForeignKeys[0].RefTable = UsersTable
	UserPostsTable.ForeignKeys[1].RefTable = PostsTable
	UserReportsTable.ForeignKeys[0].RefTable = UsersTable
	UserReportsTable.ForeignKeys[1].RefTable = ReportsTable
}
