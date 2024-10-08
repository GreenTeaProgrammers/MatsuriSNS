// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/event"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/eventadmin"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/post"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/postimage"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/schema"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescTitle is the schema descriptor for title field.
	eventDescTitle := eventFields[0].Descriptor()
	// event.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	event.TitleValidator = eventDescTitle.Validators[0].(func(string) error)
	// eventDescMapURL is the schema descriptor for map_url field.
	eventDescMapURL := eventFields[2].Descriptor()
	// event.MapURLValidator is a validator for the "map_url" field. It is called by the builders before save.
	event.MapURLValidator = eventDescMapURL.Validators[0].(func(string) error)
	// eventDescCreatedAt is the schema descriptor for created_at field.
	eventDescCreatedAt := eventFields[6].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the created_at field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescUpdatedAt is the schema descriptor for updated_at field.
	eventDescUpdatedAt := eventFields[7].Descriptor()
	// event.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	event.DefaultUpdatedAt = eventDescUpdatedAt.Default.(func() time.Time)
	// event.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	event.UpdateDefaultUpdatedAt = eventDescUpdatedAt.UpdateDefault.(func() time.Time)
	eventadminFields := schema.EventAdmin{}.Fields()
	_ = eventadminFields
	// eventadminDescCreatedAt is the schema descriptor for created_at field.
	eventadminDescCreatedAt := eventadminFields[2].Descriptor()
	// eventadmin.DefaultCreatedAt holds the default value on creation for the created_at field.
	eventadmin.DefaultCreatedAt = eventadminDescCreatedAt.Default.(func() time.Time)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescContent is the schema descriptor for content field.
	postDescContent := postFields[2].Descriptor()
	// post.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	post.ContentValidator = postDescContent.Validators[0].(func(string) error)
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[6].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updated_at field.
	postDescUpdatedAt := postFields[7].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
	postimageFields := schema.PostImage{}.Fields()
	_ = postimageFields
	// postimageDescImageURL is the schema descriptor for image_url field.
	postimageDescImageURL := postimageFields[1].Descriptor()
	// postimage.ImageURLValidator is a validator for the "image_url" field. It is called by the builders before save.
	postimage.ImageURLValidator = postimageDescImageURL.Validators[0].(func(string) error)
	// postimageDescCreatedAt is the schema descriptor for created_at field.
	postimageDescCreatedAt := postimageFields[2].Descriptor()
	// postimage.DefaultCreatedAt holds the default value on creation for the created_at field.
	postimage.DefaultCreatedAt = postimageDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescHashedPassword is the schema descriptor for hashed_password field.
	userDescHashedPassword := userFields[1].Descriptor()
	// user.HashedPasswordValidator is a validator for the "hashed_password" field. It is called by the builders before save.
	user.HashedPasswordValidator = userDescHashedPassword.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[3].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
