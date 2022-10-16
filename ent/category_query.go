// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"webreader/ent/category"
	"webreader/ent/predicate"
	"webreader/ent/ranobe"
	"webreader/ent/schema/ulid"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryQuery is the builder for querying Category entities.
type CategoryQuery struct {
	config
	limit            *int
	offset           *int
	unique           *bool
	order            []OrderFunc
	fields           []string
	predicates       []predicate.Category
	withRanobes      *RanobeQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*Category) error
	withNamedRanobes map[string]*RanobeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CategoryQuery builder.
func (cq *CategoryQuery) Where(ps ...predicate.Category) *CategoryQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CategoryQuery) Limit(limit int) *CategoryQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CategoryQuery) Offset(offset int) *CategoryQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CategoryQuery) Unique(unique bool) *CategoryQuery {
	cq.unique = &unique
	return cq
}

// Order adds an order step to the query.
func (cq *CategoryQuery) Order(o ...OrderFunc) *CategoryQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryRanobes chains the current query on the "ranobes" edge.
func (cq *CategoryQuery) QueryRanobes() *RanobeQuery {
	query := &RanobeQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, selector),
			sqlgraph.To(ranobe.Table, ranobe.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, category.RanobesTable, category.RanobesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Category entity from the query.
// Returns a *NotFoundError when no Category was found.
func (cq *CategoryQuery) First(ctx context.Context) (*Category, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{category.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CategoryQuery) FirstX(ctx context.Context) *Category {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Category ID from the query.
// Returns a *NotFoundError when no Category ID was found.
func (cq *CategoryQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{category.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CategoryQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Category entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Category entity is found.
// Returns a *NotFoundError when no Category entities are found.
func (cq *CategoryQuery) Only(ctx context.Context) (*Category, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{category.Label}
	default:
		return nil, &NotSingularError{category.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CategoryQuery) OnlyX(ctx context.Context) *Category {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Category ID in the query.
// Returns a *NotSingularError when more than one Category ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CategoryQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{category.Label}
	default:
		err = &NotSingularError{category.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CategoryQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Categories.
func (cq *CategoryQuery) All(ctx context.Context) ([]*Category, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CategoryQuery) AllX(ctx context.Context) []*Category {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Category IDs.
func (cq *CategoryQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := cq.Select(category.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CategoryQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CategoryQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CategoryQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CategoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CategoryQuery) Clone() *CategoryQuery {
	if cq == nil {
		return nil
	}
	return &CategoryQuery{
		config:      cq.config,
		limit:       cq.limit,
		offset:      cq.offset,
		order:       append([]OrderFunc{}, cq.order...),
		predicates:  append([]predicate.Category{}, cq.predicates...),
		withRanobes: cq.withRanobes.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithRanobes tells the query-builder to eager-load the nodes that are connected to
// the "ranobes" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CategoryQuery) WithRanobes(opts ...func(*RanobeQuery)) *CategoryQuery {
	query := &RanobeQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withRanobes = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Category.Query().
//		GroupBy(category.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CategoryQuery) GroupBy(field string, fields ...string) *CategoryGroupBy {
	grbuild := &CategoryGroupBy{config: cq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(ctx), nil
	}
	grbuild.label = category.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Category.Query().
//		Select(category.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CategoryQuery) Select(fields ...string) *CategorySelect {
	cq.fields = append(cq.fields, fields...)
	selbuild := &CategorySelect{CategoryQuery: cq}
	selbuild.label = category.Label
	selbuild.flds, selbuild.scan = &cq.fields, selbuild.Scan
	return selbuild
}

func (cq *CategoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !category.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Category, error) {
	var (
		nodes       = []*Category{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withRanobes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Category).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Category{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withRanobes; query != nil {
		if err := cq.loadRanobes(ctx, query, nodes,
			func(n *Category) { n.Edges.Ranobes = []*Ranobe{} },
			func(n *Category, e *Ranobe) { n.Edges.Ranobes = append(n.Edges.Ranobes, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedRanobes {
		if err := cq.loadRanobes(ctx, query, nodes,
			func(n *Category) { n.appendNamedRanobes(name) },
			func(n *Category, e *Ranobe) { n.appendNamedRanobes(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CategoryQuery) loadRanobes(ctx context.Context, query *RanobeQuery, nodes []*Category, init func(*Category), assign func(*Category, *Ranobe)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[ulid.ID]*Category)
	nids := make(map[ulid.ID]map[*Category]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(category.RanobesTable)
		s.Join(joinT).On(s.C(ranobe.FieldID), joinT.C(category.RanobesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(category.RanobesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(category.RanobesPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(ulid.ID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*ulid.ID)
			inValue := *values[1].(*ulid.ID)
			if nids[inValue] == nil {
				nids[inValue] = map[*Category]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "ranobes" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *CategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CategoryQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (cq *CategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: category.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for i := range fields {
			if fields[i] != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(category.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = category.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedRanobes tells the query-builder to eager-load the nodes that are connected to the "ranobes"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CategoryQuery) WithNamedRanobes(name string, opts ...func(*RanobeQuery)) *CategoryQuery {
	query := &RanobeQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedRanobes == nil {
		cq.withNamedRanobes = make(map[string]*RanobeQuery)
	}
	cq.withNamedRanobes[name] = query
	return cq
}

// CategoryGroupBy is the group-by builder for Category entities.
type CategoryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CategoryGroupBy) Aggregate(fns ...AggregateFunc) *CategoryGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *CategoryGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

func (cgb *CategoryGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cgb.fields {
		if !category.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *CategoryGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql.Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
		for _, f := range cgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cgb.fields...)...)
}

// CategorySelect is the builder for selecting fields of Category entities.
type CategorySelect struct {
	*CategoryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CategorySelect) Scan(ctx context.Context, v any) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.CategoryQuery.sqlQuery(ctx)
	return cs.sqlScan(ctx, v)
}

func (cs *CategorySelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := cs.sql.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
