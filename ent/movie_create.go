// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"interface_project/ent/movie"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MovieCreate is the builder for creating a Movie entity.
type MovieCreate struct {
	config
	mutation *MovieMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (mc *MovieCreate) SetTitle(s string) *MovieCreate {
	mc.mutation.SetTitle(s)
	return mc
}

// SetYear sets the "year" field.
func (mc *MovieCreate) SetYear(s string) *MovieCreate {
	mc.mutation.SetYear(s)
	return mc
}

// SetImageURL sets the "image_url" field.
func (mc *MovieCreate) SetImageURL(s string) *MovieCreate {
	mc.mutation.SetImageURL(s)
	return mc
}

// SetRuntimeStr sets the "runtimeStr" field.
func (mc *MovieCreate) SetRuntimeStr(s string) *MovieCreate {
	mc.mutation.SetRuntimeStr(s)
	return mc
}

// SetGenres sets the "genres" field.
func (mc *MovieCreate) SetGenres(s string) *MovieCreate {
	mc.mutation.SetGenres(s)
	return mc
}

// SetImDbRating sets the "imDbRating" field.
func (mc *MovieCreate) SetImDbRating(s string) *MovieCreate {
	mc.mutation.SetImDbRating(s)
	return mc
}

// SetPlot sets the "plot" field.
func (mc *MovieCreate) SetPlot(s string) *MovieCreate {
	mc.mutation.SetPlot(s)
	return mc
}

// SetStars sets the "stars" field.
func (mc *MovieCreate) SetStars(s string) *MovieCreate {
	mc.mutation.SetStars(s)
	return mc
}

// SetMetacriticRating sets the "metacriticRating" field.
func (mc *MovieCreate) SetMetacriticRating(s string) *MovieCreate {
	mc.mutation.SetMetacriticRating(s)
	return mc
}

// Mutation returns the MovieMutation object of the builder.
func (mc *MovieCreate) Mutation() *MovieMutation {
	return mc.mutation
}

// Save creates the Movie in the database.
func (mc *MovieCreate) Save(ctx context.Context) (*Movie, error) {
	var (
		err  error
		node *Movie
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MovieMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MovieCreate) SaveX(ctx context.Context) *Movie {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MovieCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MovieCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MovieCreate) check() error {
	if _, ok := mc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Movie.title"`)}
	}
	if v, ok := mc.mutation.Title(); ok {
		if err := movie.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Movie.title": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "Movie.year"`)}
	}
	if _, ok := mc.mutation.ImageURL(); !ok {
		return &ValidationError{Name: "image_url", err: errors.New(`ent: missing required field "Movie.image_url"`)}
	}
	if _, ok := mc.mutation.RuntimeStr(); !ok {
		return &ValidationError{Name: "runtimeStr", err: errors.New(`ent: missing required field "Movie.runtimeStr"`)}
	}
	if _, ok := mc.mutation.Genres(); !ok {
		return &ValidationError{Name: "genres", err: errors.New(`ent: missing required field "Movie.genres"`)}
	}
	if _, ok := mc.mutation.ImDbRating(); !ok {
		return &ValidationError{Name: "imDbRating", err: errors.New(`ent: missing required field "Movie.imDbRating"`)}
	}
	if _, ok := mc.mutation.Plot(); !ok {
		return &ValidationError{Name: "plot", err: errors.New(`ent: missing required field "Movie.plot"`)}
	}
	if _, ok := mc.mutation.Stars(); !ok {
		return &ValidationError{Name: "stars", err: errors.New(`ent: missing required field "Movie.stars"`)}
	}
	if _, ok := mc.mutation.MetacriticRating(); !ok {
		return &ValidationError{Name: "metacriticRating", err: errors.New(`ent: missing required field "Movie.metacriticRating"`)}
	}
	return nil
}

func (mc *MovieCreate) sqlSave(ctx context.Context) (*Movie, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MovieCreate) createSpec() (*Movie, *sqlgraph.CreateSpec) {
	var (
		_node = &Movie{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: movie.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: movie.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := mc.mutation.Year(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldYear,
		})
		_node.Year = value
	}
	if value, ok := mc.mutation.ImageURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldImageURL,
		})
		_node.ImageURL = value
	}
	if value, ok := mc.mutation.RuntimeStr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldRuntimeStr,
		})
		_node.RuntimeStr = value
	}
	if value, ok := mc.mutation.Genres(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldGenres,
		})
		_node.Genres = value
	}
	if value, ok := mc.mutation.ImDbRating(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldImDbRating,
		})
		_node.ImDbRating = value
	}
	if value, ok := mc.mutation.Plot(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldPlot,
		})
		_node.Plot = value
	}
	if value, ok := mc.mutation.Stars(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldStars,
		})
		_node.Stars = value
	}
	if value, ok := mc.mutation.MetacriticRating(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldMetacriticRating,
		})
		_node.MetacriticRating = value
	}
	return _node, _spec
}

// MovieCreateBulk is the builder for creating many Movie entities in bulk.
type MovieCreateBulk struct {
	config
	builders []*MovieCreate
}

// Save creates the Movie entities in the database.
func (mcb *MovieCreateBulk) Save(ctx context.Context) ([]*Movie, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Movie, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MovieMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MovieCreateBulk) SaveX(ctx context.Context) []*Movie {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MovieCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MovieCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
