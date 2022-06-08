// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"interface_project/ent/fileentity"
	"interface_project/ent/user"
	"interface_project/ent/wordnode"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FileEntityCreate is the builder for creating a FileEntity entity.
type FileEntityCreate struct {
	config
	mutation *FileEntityMutation
	hooks    []Hook
}

// SetPath sets the "path" field.
func (fec *FileEntityCreate) SetPath(s string) *FileEntityCreate {
	fec.mutation.SetPath(s)
	return fec
}

// SetName sets the "name" field.
func (fec *FileEntityCreate) SetName(s string) *FileEntityCreate {
	fec.mutation.SetName(s)
	return fec
}

// SetSize sets the "size" field.
func (fec *FileEntityCreate) SetSize(i int16) *FileEntityCreate {
	fec.mutation.SetSize(i)
	return fec
}

// SetDeleted sets the "deleted" field.
func (fec *FileEntityCreate) SetDeleted(b bool) *FileEntityCreate {
	fec.mutation.SetDeleted(b)
	return fec
}

// SetCreatedDate sets the "created_date" field.
func (fec *FileEntityCreate) SetCreatedDate(t time.Time) *FileEntityCreate {
	fec.mutation.SetCreatedDate(t)
	return fec
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fec *FileEntityCreate) SetOwnerID(id int) *FileEntityCreate {
	fec.mutation.SetOwnerID(id)
	return fec
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fec *FileEntityCreate) SetNillableOwnerID(id *int) *FileEntityCreate {
	if id != nil {
		fec = fec.SetOwnerID(*id)
	}
	return fec
}

// SetOwner sets the "owner" edge to the User entity.
func (fec *FileEntityCreate) SetOwner(u *User) *FileEntityCreate {
	return fec.SetOwnerID(u.ID)
}

// AddWordnodeIDs adds the "wordnodes" edge to the WordNode entity by IDs.
func (fec *FileEntityCreate) AddWordnodeIDs(ids ...int) *FileEntityCreate {
	fec.mutation.AddWordnodeIDs(ids...)
	return fec
}

// AddWordnodes adds the "wordnodes" edges to the WordNode entity.
func (fec *FileEntityCreate) AddWordnodes(w ...*WordNode) *FileEntityCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return fec.AddWordnodeIDs(ids...)
}

// Mutation returns the FileEntityMutation object of the builder.
func (fec *FileEntityCreate) Mutation() *FileEntityMutation {
	return fec.mutation
}

// Save creates the FileEntity in the database.
func (fec *FileEntityCreate) Save(ctx context.Context) (*FileEntity, error) {
	var (
		err  error
		node *FileEntity
	)
	if len(fec.hooks) == 0 {
		if err = fec.check(); err != nil {
			return nil, err
		}
		node, err = fec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileEntityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fec.check(); err != nil {
				return nil, err
			}
			fec.mutation = mutation
			if node, err = fec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fec.hooks) - 1; i >= 0; i-- {
			if fec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fec *FileEntityCreate) SaveX(ctx context.Context) *FileEntity {
	v, err := fec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fec *FileEntityCreate) Exec(ctx context.Context) error {
	_, err := fec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fec *FileEntityCreate) ExecX(ctx context.Context) {
	if err := fec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fec *FileEntityCreate) check() error {
	if _, ok := fec.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "FileEntity.path"`)}
	}
	if v, ok := fec.mutation.Path(); ok {
		if err := fileentity.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "FileEntity.path": %w`, err)}
		}
	}
	if _, ok := fec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FileEntity.name"`)}
	}
	if v, ok := fec.mutation.Name(); ok {
		if err := fileentity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "FileEntity.name": %w`, err)}
		}
	}
	if _, ok := fec.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "FileEntity.size"`)}
	}
	if _, ok := fec.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "FileEntity.deleted"`)}
	}
	if _, ok := fec.mutation.CreatedDate(); !ok {
		return &ValidationError{Name: "created_date", err: errors.New(`ent: missing required field "FileEntity.created_date"`)}
	}
	return nil
}

func (fec *FileEntityCreate) sqlSave(ctx context.Context) (*FileEntity, error) {
	_node, _spec := fec.createSpec()
	if err := sqlgraph.CreateNode(ctx, fec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fec *FileEntityCreate) createSpec() (*FileEntity, *sqlgraph.CreateSpec) {
	var (
		_node = &FileEntity{config: fec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fileentity.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fileentity.FieldID,
			},
		}
	)
	if value, ok := fec.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileentity.FieldPath,
		})
		_node.Path = value
	}
	if value, ok := fec.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileentity.FieldName,
		})
		_node.Name = value
	}
	if value, ok := fec.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: fileentity.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := fec.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: fileentity.FieldDeleted,
		})
		_node.Deleted = value
	}
	if value, ok := fec.mutation.CreatedDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fileentity.FieldCreatedDate,
		})
		_node.CreatedDate = value
	}
	if nodes := fec.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fileentity.OwnerTable,
			Columns: []string{fileentity.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_files = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fec.mutation.WordnodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fileentity.WordnodesTable,
			Columns: []string{fileentity.WordnodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wordnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FileEntityCreateBulk is the builder for creating many FileEntity entities in bulk.
type FileEntityCreateBulk struct {
	config
	builders []*FileEntityCreate
}

// Save creates the FileEntity entities in the database.
func (fecb *FileEntityCreateBulk) Save(ctx context.Context) ([]*FileEntity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fecb.builders))
	nodes := make([]*FileEntity, len(fecb.builders))
	mutators := make([]Mutator, len(fecb.builders))
	for i := range fecb.builders {
		func(i int, root context.Context) {
			builder := fecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileEntityMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fecb *FileEntityCreateBulk) SaveX(ctx context.Context) []*FileEntity {
	v, err := fecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fecb *FileEntityCreateBulk) Exec(ctx context.Context) error {
	_, err := fecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fecb *FileEntityCreateBulk) ExecX(ctx context.Context) {
	if err := fecb.Exec(ctx); err != nil {
		panic(err)
	}
}