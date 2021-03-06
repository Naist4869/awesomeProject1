// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Naist4869/awesomeProject1/ent/issuex1355"
	"github.com/Naist4869/awesomeProject1/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeIssueX1355 = "IssueX1355"
)

// IssueX1355Mutation represents an operation that mutates the IssueX1355 nodes in the graph.
type IssueX1355Mutation struct {
	config
	op            Op
	typ           string
	id            *int
	time          *time.Time
	int32         *int32
	addint32      *int32
	str           *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*IssueX1355, error)
	predicates    []predicate.IssueX1355
}

var _ ent.Mutation = (*IssueX1355Mutation)(nil)

// issuex1355Option allows management of the mutation configuration using functional options.
type issuex1355Option func(*IssueX1355Mutation)

// newIssueX1355Mutation creates new mutation for the IssueX1355 entity.
func newIssueX1355Mutation(c config, op Op, opts ...issuex1355Option) *IssueX1355Mutation {
	m := &IssueX1355Mutation{
		config:        c,
		op:            op,
		typ:           TypeIssueX1355,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withIssueX1355ID sets the ID field of the mutation.
func withIssueX1355ID(id int) issuex1355Option {
	return func(m *IssueX1355Mutation) {
		var (
			err   error
			once  sync.Once
			value *IssueX1355
		)
		m.oldValue = func(ctx context.Context) (*IssueX1355, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().IssueX1355.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withIssueX1355 sets the old IssueX1355 of the mutation.
func withIssueX1355(node *IssueX1355) issuex1355Option {
	return func(m *IssueX1355Mutation) {
		m.oldValue = func(context.Context) (*IssueX1355, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m IssueX1355Mutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m IssueX1355Mutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *IssueX1355Mutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetTime sets the "time" field.
func (m *IssueX1355Mutation) SetTime(t time.Time) {
	m.time = &t
}

// Time returns the value of the "time" field in the mutation.
func (m *IssueX1355Mutation) Time() (r time.Time, exists bool) {
	v := m.time
	if v == nil {
		return
	}
	return *v, true
}

// OldTime returns the old "time" field's value of the IssueX1355 entity.
// If the IssueX1355 object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IssueX1355Mutation) OldTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTime: %w", err)
	}
	return oldValue.Time, nil
}

// ClearTime clears the value of the "time" field.
func (m *IssueX1355Mutation) ClearTime() {
	m.time = nil
	m.clearedFields[issuex1355.FieldTime] = struct{}{}
}

// TimeCleared returns if the "time" field was cleared in this mutation.
func (m *IssueX1355Mutation) TimeCleared() bool {
	_, ok := m.clearedFields[issuex1355.FieldTime]
	return ok
}

// ResetTime resets all changes to the "time" field.
func (m *IssueX1355Mutation) ResetTime() {
	m.time = nil
	delete(m.clearedFields, issuex1355.FieldTime)
}

// SetInt32 sets the "int32" field.
func (m *IssueX1355Mutation) SetInt32(i int32) {
	m.int32 = &i
	m.addint32 = nil
}

// Int32 returns the value of the "int32" field in the mutation.
func (m *IssueX1355Mutation) Int32() (r int32, exists bool) {
	v := m.int32
	if v == nil {
		return
	}
	return *v, true
}

// OldInt32 returns the old "int32" field's value of the IssueX1355 entity.
// If the IssueX1355 object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IssueX1355Mutation) OldInt32(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldInt32 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldInt32 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInt32: %w", err)
	}
	return oldValue.Int32, nil
}

// AddInt32 adds i to the "int32" field.
func (m *IssueX1355Mutation) AddInt32(i int32) {
	if m.addint32 != nil {
		*m.addint32 += i
	} else {
		m.addint32 = &i
	}
}

// AddedInt32 returns the value that was added to the "int32" field in this mutation.
func (m *IssueX1355Mutation) AddedInt32() (r int32, exists bool) {
	v := m.addint32
	if v == nil {
		return
	}
	return *v, true
}

// ClearInt32 clears the value of the "int32" field.
func (m *IssueX1355Mutation) ClearInt32() {
	m.int32 = nil
	m.addint32 = nil
	m.clearedFields[issuex1355.FieldInt32] = struct{}{}
}

// Int32Cleared returns if the "int32" field was cleared in this mutation.
func (m *IssueX1355Mutation) Int32Cleared() bool {
	_, ok := m.clearedFields[issuex1355.FieldInt32]
	return ok
}

// ResetInt32 resets all changes to the "int32" field.
func (m *IssueX1355Mutation) ResetInt32() {
	m.int32 = nil
	m.addint32 = nil
	delete(m.clearedFields, issuex1355.FieldInt32)
}

// SetStr sets the "str" field.
func (m *IssueX1355Mutation) SetStr(s string) {
	m.str = &s
}

// Str returns the value of the "str" field in the mutation.
func (m *IssueX1355Mutation) Str() (r string, exists bool) {
	v := m.str
	if v == nil {
		return
	}
	return *v, true
}

// OldStr returns the old "str" field's value of the IssueX1355 entity.
// If the IssueX1355 object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IssueX1355Mutation) OldStr(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStr is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStr requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStr: %w", err)
	}
	return oldValue.Str, nil
}

// ResetStr resets all changes to the "str" field.
func (m *IssueX1355Mutation) ResetStr() {
	m.str = nil
}

// Op returns the operation name.
func (m *IssueX1355Mutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (IssueX1355).
func (m *IssueX1355Mutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *IssueX1355Mutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.time != nil {
		fields = append(fields, issuex1355.FieldTime)
	}
	if m.int32 != nil {
		fields = append(fields, issuex1355.FieldInt32)
	}
	if m.str != nil {
		fields = append(fields, issuex1355.FieldStr)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *IssueX1355Mutation) Field(name string) (ent.Value, bool) {
	switch name {
	case issuex1355.FieldTime:
		return m.Time()
	case issuex1355.FieldInt32:
		return m.Int32()
	case issuex1355.FieldStr:
		return m.Str()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *IssueX1355Mutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case issuex1355.FieldTime:
		return m.OldTime(ctx)
	case issuex1355.FieldInt32:
		return m.OldInt32(ctx)
	case issuex1355.FieldStr:
		return m.OldStr(ctx)
	}
	return nil, fmt.Errorf("unknown IssueX1355 field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *IssueX1355Mutation) SetField(name string, value ent.Value) error {
	switch name {
	case issuex1355.FieldTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTime(v)
		return nil
	case issuex1355.FieldInt32:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInt32(v)
		return nil
	case issuex1355.FieldStr:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStr(v)
		return nil
	}
	return fmt.Errorf("unknown IssueX1355 field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *IssueX1355Mutation) AddedFields() []string {
	var fields []string
	if m.addint32 != nil {
		fields = append(fields, issuex1355.FieldInt32)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *IssueX1355Mutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case issuex1355.FieldInt32:
		return m.AddedInt32()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *IssueX1355Mutation) AddField(name string, value ent.Value) error {
	switch name {
	case issuex1355.FieldInt32:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddInt32(v)
		return nil
	}
	return fmt.Errorf("unknown IssueX1355 numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *IssueX1355Mutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(issuex1355.FieldTime) {
		fields = append(fields, issuex1355.FieldTime)
	}
	if m.FieldCleared(issuex1355.FieldInt32) {
		fields = append(fields, issuex1355.FieldInt32)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *IssueX1355Mutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *IssueX1355Mutation) ClearField(name string) error {
	switch name {
	case issuex1355.FieldTime:
		m.ClearTime()
		return nil
	case issuex1355.FieldInt32:
		m.ClearInt32()
		return nil
	}
	return fmt.Errorf("unknown IssueX1355 nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *IssueX1355Mutation) ResetField(name string) error {
	switch name {
	case issuex1355.FieldTime:
		m.ResetTime()
		return nil
	case issuex1355.FieldInt32:
		m.ResetInt32()
		return nil
	case issuex1355.FieldStr:
		m.ResetStr()
		return nil
	}
	return fmt.Errorf("unknown IssueX1355 field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *IssueX1355Mutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *IssueX1355Mutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *IssueX1355Mutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *IssueX1355Mutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *IssueX1355Mutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *IssueX1355Mutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *IssueX1355Mutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown IssueX1355 unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *IssueX1355Mutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown IssueX1355 edge %s", name)
}
