// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cms.proto

/*
Package cms is a generated protocol buffer package.

CMS

CMS Service API provides access to CMS entities and supports CMS use-cases

It is generated from these files:
	cms.proto

It has these top-level messages:
	Post
	PostRequest
	CreatePostRequest
	UpdatePostRequest
	Comment
	CommentRequest
	CreateCommentRequest
	UpdateCommentRequest
	User
	UserRequest
	CreateUserRequest
*/
package cms

import regexp "regexp"
import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Post) Validate() error {
	return nil
}
func (this *PostRequest) Validate() error {
	if !(this.Id > 0) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *CreatePostRequest) Validate() error {
	if !(len(this.Title) < 256) {
		return go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must length be less than '256'`, this.Title))
	}
	if !(len(this.Content) < 16777216) {
		return go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must length be less than '16777216'`, this.Content))
	}
	if !(len(this.Slug) == 0) {
		return go_proto_validators.FieldError("Slug", fmt.Errorf(`Do not include a slug value in your request. This field is only used internally by the server.`))
	}
	return nil
}
func (this *UpdatePostRequest) Validate() error {
	if !(this.Id > 0) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	if !(len(this.Title) < 256) {
		return go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must length be less than '256'`, this.Title))
	}
	if !(len(this.Content) < 16777216) {
		return go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must length be less than '16777216'`, this.Content))
	}
	if !(len(this.Slug) == 0) {
		return go_proto_validators.FieldError("Slug", fmt.Errorf(`Do not include a slug value in your request. This field is only used internally by the server.`))
	}
	return nil
}
func (this *Comment) Validate() error {
	return nil
}
func (this *CommentRequest) Validate() error {
	if !(this.Id > 0) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}

var _regex_CreateCommentRequest_UserId = regexp.MustCompile("^[a-zA-Z]{3,18}$")

func (this *CreateCommentRequest) Validate() error {
	if this.Content == "" {
		return go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must not be an empty string`, this.Content))
	}
	if !(len(this.Content) < 65536) {
		return go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must length be less than '65536'`, this.Content))
	}
	if !_regex_CreateCommentRequest_UserId.MatchString(this.UserId) {
		return go_proto_validators.FieldError("UserId", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if this.UserId == "" {
		return go_proto_validators.FieldError("UserId", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if this.PostId == "" {
		return go_proto_validators.FieldError("PostId", fmt.Errorf(`value '%v' must not be an empty string`, this.PostId))
	}
	if !(len(this.PostId) < 37) {
		return go_proto_validators.FieldError("PostId", fmt.Errorf(`value '%v' must length be less than '37'`, this.PostId))
	}
	return nil
}
func (this *UpdateCommentRequest) Validate() error {
	if !(this.Id > 0) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	if this.Content == "" {
		return go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must not be an empty string`, this.Content))
	}
	if !(len(this.Content) < 65536) {
		return go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must length be less than '65536'`, this.Content))
	}
	return nil
}
func (this *User) Validate() error {
	return nil
}

var _regex_UserRequest_Id = regexp.MustCompile("^[a-zA-Z]{3,18}$")

func (this *UserRequest) Validate() error {
	if !_regex_UserRequest_Id.MatchString(this.Id) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if this.Id == "" {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	return nil
}

var _regex_CreateUserRequest_Id = regexp.MustCompile("^[a-zA-Z]{3,18}$")

func (this *CreateUserRequest) Validate() error {
	if !_regex_CreateUserRequest_Id.MatchString(this.Id) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if this.Id == "" {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if this.Email == "" {
		return go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must not be an empty string`, this.Email))
	}
	if !(len(this.Email) < 256) {
		return go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must length be less than '256'`, this.Email))
	}
	if this.Password == "" {
		return go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	if !(len(this.Password) < 51) {
		return go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must length be less than '51'`, this.Password))
	}
	return nil
}
