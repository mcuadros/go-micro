// Package model is a data modelling interface
package model

import (
	"github.com/micro/go-micro/v2/codec"
	"github.com/micro/go-micro/v2/store"
)

// Model represents a data model interface for data access
type Model interface {
	// Initialise options
	Init(...Option) error
	// Retrieve the options
	Options() Options
	// New entity returns a new entity
	NewEntity(id string, value interface{})
	// Register a data type
	Register(v interface{}) error
	// Create an entity
	Create(Entity, ...CreateOption) error
	// Read an into the value provided
	Read(v interface{}, opts ...ReadOption) error
	// Update an enity
	Update(Entity, ...UpdateOption) error
	// Delete an entity
	Delete(Entity, ...DeleteOption) error
	// The implementation e.g crud
	String() string
}

type Entity interface {
	// Id of the entity
	ID() string
	// Associated value
	Value() interface{}
}

// Options to the model
type Options struct {
	// The store for data storage
	Store store.Store
	// The codec for marshalling
	Codec codec.Marshaler
}

type Option func(o *Options) error

// CreateOptions for writing entities
type CreateOptions struct{}

type CreateOption func(o *CreateOptions) error

// ReadOptions for reading entities
type ReadOptions struct{}

type ReadOption func(o *ReadOptions) error

// UpdateOptions for updating entities
type UpdateOptions struct{}

type UpdateOption func(o *UpdateOptions) error

// DeleteOptions for deleting entities
type DeleteOptions struct{}

type DeleteOption func(o *DeleteOptions) error
