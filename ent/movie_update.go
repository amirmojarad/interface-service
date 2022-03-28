// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"interface_project/ent/movie"
	"interface_project/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MovieUpdate is the builder for updating Movie entities.
type MovieUpdate struct {
	config
	hooks    []Hook
	mutation *MovieMutation
}

// Where appends a list predicates to the MovieUpdate builder.
func (mu *MovieUpdate) Where(ps ...predicate.Movie) *MovieUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetTitle sets the "title" field.
func (mu *MovieUpdate) SetTitle(s string) *MovieUpdate {
	mu.mutation.SetTitle(s)
	return mu
}

// SetYear sets the "year" field.
func (mu *MovieUpdate) SetYear(s string) *MovieUpdate {
	mu.mutation.SetYear(s)
	return mu
}

// SetImageURL sets the "image_url" field.
func (mu *MovieUpdate) SetImageURL(s string) *MovieUpdate {
	mu.mutation.SetImageURL(s)
	return mu
}

// SetRuntimeStr sets the "runtimeStr" field.
func (mu *MovieUpdate) SetRuntimeStr(s string) *MovieUpdate {
	mu.mutation.SetRuntimeStr(s)
	return mu
}

// SetGenres sets the "genres" field.
func (mu *MovieUpdate) SetGenres(s string) *MovieUpdate {
	mu.mutation.SetGenres(s)
	return mu
}

// SetImDbRating sets the "imDbRating" field.
func (mu *MovieUpdate) SetImDbRating(s string) *MovieUpdate {
	mu.mutation.SetImDbRating(s)
	return mu
}

// SetPlot sets the "plot" field.
func (mu *MovieUpdate) SetPlot(s string) *MovieUpdate {
	mu.mutation.SetPlot(s)
	return mu
}

// SetStars sets the "stars" field.
func (mu *MovieUpdate) SetStars(s string) *MovieUpdate {
	mu.mutation.SetStars(s)
	return mu
}

// SetMetacriticRating sets the "metacriticRating" field.
func (mu *MovieUpdate) SetMetacriticRating(s string) *MovieUpdate {
	mu.mutation.SetMetacriticRating(s)
	return mu
}

// Mutation returns the MovieMutation object of the builder.
func (mu *MovieUpdate) Mutation() *MovieMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MovieUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MovieMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MovieUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MovieUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MovieUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MovieUpdate) check() error {
	if v, ok := mu.mutation.Title(); ok {
		if err := movie.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Movie.title": %w`, err)}
		}
	}
	return nil
}

func (mu *MovieUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   movie.Table,
			Columns: movie.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: movie.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldTitle,
		})
	}
	if value, ok := mu.mutation.Year(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldYear,
		})
	}
	if value, ok := mu.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldImageURL,
		})
	}
	if value, ok := mu.mutation.RuntimeStr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldRuntimeStr,
		})
	}
	if value, ok := mu.mutation.Genres(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldGenres,
		})
	}
	if value, ok := mu.mutation.ImDbRating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldImDbRating,
		})
	}
	if value, ok := mu.mutation.Plot(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldPlot,
		})
	}
	if value, ok := mu.mutation.Stars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldStars,
		})
	}
	if value, ok := mu.mutation.MetacriticRating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldMetacriticRating,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{movie.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MovieUpdateOne is the builder for updating a single Movie entity.
type MovieUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MovieMutation
}

// SetTitle sets the "title" field.
func (muo *MovieUpdateOne) SetTitle(s string) *MovieUpdateOne {
	muo.mutation.SetTitle(s)
	return muo
}

// SetYear sets the "year" field.
func (muo *MovieUpdateOne) SetYear(s string) *MovieUpdateOne {
	muo.mutation.SetYear(s)
	return muo
}

// SetImageURL sets the "image_url" field.
func (muo *MovieUpdateOne) SetImageURL(s string) *MovieUpdateOne {
	muo.mutation.SetImageURL(s)
	return muo
}

// SetRuntimeStr sets the "runtimeStr" field.
func (muo *MovieUpdateOne) SetRuntimeStr(s string) *MovieUpdateOne {
	muo.mutation.SetRuntimeStr(s)
	return muo
}

// SetGenres sets the "genres" field.
func (muo *MovieUpdateOne) SetGenres(s string) *MovieUpdateOne {
	muo.mutation.SetGenres(s)
	return muo
}

// SetImDbRating sets the "imDbRating" field.
func (muo *MovieUpdateOne) SetImDbRating(s string) *MovieUpdateOne {
	muo.mutation.SetImDbRating(s)
	return muo
}

// SetPlot sets the "plot" field.
func (muo *MovieUpdateOne) SetPlot(s string) *MovieUpdateOne {
	muo.mutation.SetPlot(s)
	return muo
}

// SetStars sets the "stars" field.
func (muo *MovieUpdateOne) SetStars(s string) *MovieUpdateOne {
	muo.mutation.SetStars(s)
	return muo
}

// SetMetacriticRating sets the "metacriticRating" field.
func (muo *MovieUpdateOne) SetMetacriticRating(s string) *MovieUpdateOne {
	muo.mutation.SetMetacriticRating(s)
	return muo
}

// Mutation returns the MovieMutation object of the builder.
func (muo *MovieUpdateOne) Mutation() *MovieMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MovieUpdateOne) Select(field string, fields ...string) *MovieUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Movie entity.
func (muo *MovieUpdateOne) Save(ctx context.Context) (*Movie, error) {
	var (
		err  error
		node *Movie
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MovieMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MovieUpdateOne) SaveX(ctx context.Context) *Movie {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MovieUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MovieUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MovieUpdateOne) check() error {
	if v, ok := muo.mutation.Title(); ok {
		if err := movie.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Movie.title": %w`, err)}
		}
	}
	return nil
}

func (muo *MovieUpdateOne) sqlSave(ctx context.Context) (_node *Movie, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   movie.Table,
			Columns: movie.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: movie.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Movie.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, movie.FieldID)
		for _, f := range fields {
			if !movie.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != movie.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldTitle,
		})
	}
	if value, ok := muo.mutation.Year(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldYear,
		})
	}
	if value, ok := muo.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldImageURL,
		})
	}
	if value, ok := muo.mutation.RuntimeStr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldRuntimeStr,
		})
	}
	if value, ok := muo.mutation.Genres(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldGenres,
		})
	}
	if value, ok := muo.mutation.ImDbRating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldImDbRating,
		})
	}
	if value, ok := muo.mutation.Plot(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldPlot,
		})
	}
	if value, ok := muo.mutation.Stars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldStars,
		})
	}
	if value, ok := muo.mutation.MetacriticRating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldMetacriticRating,
		})
	}
	_node = &Movie{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{movie.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
