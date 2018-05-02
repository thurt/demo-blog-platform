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
	PostBySlugRequest
	CreatePostRequest
	CreatePostWithSlug
	UpdatePostRequest
	UpdatePostWithSlug
	Comment
	CommentRequest
	CreateCommentRequest
	UpdateCommentRequest
	User
	UserRequest
	CreateUserRequest
	CreateUserWithRole
	AuthUserRequest
	AccessToken
	VerifyNewUserRequest
	Email
	UserPassword
*/
package cms

import regexp "regexp"
import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import goregen "github.com/zach-klippenstein/goregen"
import gofuzz "github.com/google/gofuzz"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Post) Validate() error {
	return nil
}
func (this *Post) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}
func (this *PostRequest) Validate() error {
	if !(this.Id > 0) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *PostRequest) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}
func (this *PostBySlugRequest) Validate() error {
	if this.Slug == "" {
		return go_proto_validators.FieldError("Slug", fmt.Errorf(`value '%v' must not be an empty string`, this.Slug))
	}
	if !(len(this.Slug) < 37) {
		return go_proto_validators.FieldError("Slug", fmt.Errorf(`value '%v' must length be less than '37'`, this.Slug))
	}
	return nil
}
func (this *PostBySlugRequest) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}

var Regex_CreatePostRequest_Title = regexp.MustCompile("^.{0,256}$")

func (this *CreatePostRequest) Validate() error {
	if !Regex_CreatePostRequest_Title.MatchString(this.Title) {
		return go_proto_validators.FieldError("Title", fmt.Errorf(`title must be 256 characters or less`))
	}
	if !(len(this.Title) < 257) {
		return go_proto_validators.FieldError("Title", fmt.Errorf(`title must be 256 characters or less`))
	}
	if !(len(this.Content) < 16777216) {
		return go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must length be less than '16777216'`, this.Content))
	}
	return nil
}
func (this *CreatePostRequest) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
	g, _ = goregen.NewGenerator(Regex_CreatePostRequest_Title.String(), &goregen.GeneratorArgs{
		RngSource: c,
	})
	this.Title = g.Generate()
}
func (this *CreatePostWithSlug) Validate() error {
	if this.Post != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Post); err != nil {
			return go_proto_validators.FieldError("Post", err)
		}
	}
	return nil
}
func (this *CreatePostWithSlug) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}

var Regex_UpdatePostRequest_Title = regexp.MustCompile("^.{0,256}$")

func (this *UpdatePostRequest) Validate() error {
	if !(this.Id > 0) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	if !Regex_UpdatePostRequest_Title.MatchString(this.Title) {
		return go_proto_validators.FieldError("Title", fmt.Errorf(`title must be 256 characters or less`))
	}
	if !(len(this.Title) < 257) {
		return go_proto_validators.FieldError("Title", fmt.Errorf(`title must be 256 characters or less`))
	}
	if !(len(this.Content) < 16777216) {
		return go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must length be less than '16777216'`, this.Content))
	}
	return nil
}
func (this *UpdatePostRequest) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
	g, _ = goregen.NewGenerator(Regex_UpdatePostRequest_Title.String(), &goregen.GeneratorArgs{
		RngSource: c,
	})
	this.Title = g.Generate()
}
func (this *UpdatePostWithSlug) Validate() error {
	if this.Post != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Post); err != nil {
			return go_proto_validators.FieldError("Post", err)
		}
	}
	return nil
}
func (this *UpdatePostWithSlug) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}
func (this *Comment) Validate() error {
	return nil
}
func (this *Comment) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}
func (this *CommentRequest) Validate() error {
	if !(this.Id > 0) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *CommentRequest) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}

var Regex_CreateCommentRequest_UserId = regexp.MustCompile("^[[:alpha:]]{3,18}$")

func (this *CreateCommentRequest) Validate() error {
	if this.Content == "" {
		return go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must not be an empty string`, this.Content))
	}
	if !(len(this.Content) < 65536) {
		return go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must length be less than '65536'`, this.Content))
	}
	if !Regex_CreateCommentRequest_UserId.MatchString(this.UserId) {
		return go_proto_validators.FieldError("UserId", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if this.UserId == "" {
		return go_proto_validators.FieldError("UserId", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if !(this.PostId > 0) {
		return go_proto_validators.FieldError("PostId", fmt.Errorf(`value '%v' must be greater than '0'`, this.PostId))
	}
	return nil
}
func (this *CreateCommentRequest) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
	g, _ = goregen.NewGenerator(Regex_CreateCommentRequest_UserId.String(), &goregen.GeneratorArgs{
		RngSource: c,
	})
	this.UserId = g.Generate()
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
func (this *UpdateCommentRequest) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}
func (this *User) Validate() error {
	return nil
}
func (this *User) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}

var Regex_UserRequest_Id = regexp.MustCompile("^[[:alpha:]]{3,18}$")

func (this *UserRequest) Validate() error {
	if !Regex_UserRequest_Id.MatchString(this.Id) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if this.Id == "" {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	return nil
}
func (this *UserRequest) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
	g, _ = goregen.NewGenerator(Regex_UserRequest_Id.String(), &goregen.GeneratorArgs{
		RngSource: c,
	})
	this.Id = g.Generate()
}

var Regex_CreateUserRequest_Id = regexp.MustCompile("^[[:alpha:]]{3,18}$")
var Regex_CreateUserRequest_Email = regexp.MustCompile("^[[:alnum:]][-_.a-zA-Z0-9]{0,63}@[-_.a-zA-Z0-9]{1,190}$")
var Regex_CreateUserRequest_Password = regexp.MustCompile("^[[:graph:]]{6,50}$")

func (this *CreateUserRequest) Validate() error {
	if !Regex_CreateUserRequest_Id.MatchString(this.Id) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if this.Id == "" {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if !Regex_CreateUserRequest_Email.MatchString(this.Email) {
		return go_proto_validators.FieldError("Email", fmt.Errorf(`Email must be valid`))
	}
	if this.Email == "" {
		return go_proto_validators.FieldError("Email", fmt.Errorf(`Email must be valid`))
	}
	if !(len(this.Email) < 256) {
		return go_proto_validators.FieldError("Email", fmt.Errorf(`Email must be valid`))
	}
	if !Regex_CreateUserRequest_Password.MatchString(this.Password) {
		return go_proto_validators.FieldError("Password", fmt.Errorf(`Password must be at least 6 printable characters without whitespaces or control characters`))
	}
	if this.Password == "" {
		return go_proto_validators.FieldError("Password", fmt.Errorf(`Password must be at least 6 printable characters without whitespaces or control characters`))
	}
	if !(len(this.Password) < 51) {
		return go_proto_validators.FieldError("Password", fmt.Errorf(`Password must be at least 6 printable characters without whitespaces or control characters`))
	}
	return nil
}
func (this *CreateUserRequest) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
	g, _ = goregen.NewGenerator(Regex_CreateUserRequest_Id.String(), &goregen.GeneratorArgs{
		RngSource: c,
	})
	this.Id = g.Generate()
	g, _ = goregen.NewGenerator(Regex_CreateUserRequest_Email.String(), &goregen.GeneratorArgs{
		RngSource: c,
	})
	this.Email = g.Generate()
	g, _ = goregen.NewGenerator(Regex_CreateUserRequest_Password.String(), &goregen.GeneratorArgs{
		RngSource: c,
	})
	this.Password = g.Generate()
}
func (this *CreateUserWithRole) Validate() error {
	if this.User != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *CreateUserWithRole) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}

var Regex_AuthUserRequest_Id = regexp.MustCompile("^[[:alpha:]]{3,18}$")
var Regex_AuthUserRequest_Password = regexp.MustCompile("^[[:graph:]]{6,50}$")

func (this *AuthUserRequest) Validate() error {
	if !Regex_AuthUserRequest_Id.MatchString(this.Id) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if this.Id == "" {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`User Id must be 3-18 characters and can only include letters`))
	}
	if !Regex_AuthUserRequest_Password.MatchString(this.Password) {
		return go_proto_validators.FieldError("Password", fmt.Errorf(`Password must be at least 6 printable characters without whitespaces or control characters`))
	}
	if this.Password == "" {
		return go_proto_validators.FieldError("Password", fmt.Errorf(`Password must be at least 6 printable characters without whitespaces or control characters`))
	}
	if !(len(this.Password) < 51) {
		return go_proto_validators.FieldError("Password", fmt.Errorf(`Password must be at least 6 printable characters without whitespaces or control characters`))
	}
	return nil
}
func (this *AuthUserRequest) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
	g, _ = goregen.NewGenerator(Regex_AuthUserRequest_Id.String(), &goregen.GeneratorArgs{
		RngSource: c,
	})
	this.Id = g.Generate()
	g, _ = goregen.NewGenerator(Regex_AuthUserRequest_Password.String(), &goregen.GeneratorArgs{
		RngSource: c,
	})
	this.Password = g.Generate()
}
func (this *AccessToken) Validate() error {
	return nil
}
func (this *AccessToken) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}

var _regex_VerifyNewUserRequest_Token = regexp.MustCompile("^[[:xdigit:]]{6}$")

func (this *VerifyNewUserRequest) Validate() error {
	if !_regex_VerifyNewUserRequest_Token.MatchString(this.Token) {
		return go_proto_validators.FieldError("Token", fmt.Errorf(`Sorry, that token does not appear valid. Please try entering the token again`))
	}
	if this.Token == "" {
		return go_proto_validators.FieldError("Token", fmt.Errorf(`Sorry, that token does not appear valid. Please try entering the token again`))
	}
	return nil
}
func (this *Email) Validate() error {
	return nil
}
func (this *UserPassword) Fuzz(c gofuzz.Continue) {
	c.FuzzNoCustom(this)
	var g goregen.Generator
	var _ = g // Reference g to suppress errors if it is not otherwise used.
}
