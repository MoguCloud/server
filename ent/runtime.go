// Code generated by entc, DO NOT EDIT.

package ent

import (
	"app/ent/person"
	"app/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	personFields := schema.Person{}.Fields()
	_ = personFields
	// personDescName is the schema descriptor for name field.
	personDescName := personFields[1].Descriptor()
	// person.NameValidator is a validator for the "name" field. It is called by the builders before save.
	person.NameValidator = personDescName.Validators[0].(func(string) error)
	// personDescImg is the schema descriptor for img field.
	personDescImg := personFields[12].Descriptor()
	// person.ImgValidator is a validator for the "img" field. It is called by the builders before save.
	person.ImgValidator = personDescImg.Validators[0].(func(string) error)
	// personDescImgAnidb is the schema descriptor for img_anidb field.
	personDescImgAnidb := personFields[13].Descriptor()
	// person.ImgAnidbValidator is a validator for the "img_anidb" field. It is called by the builders before save.
	person.ImgAnidbValidator = personDescImgAnidb.Validators[0].(func(string) error)
}
