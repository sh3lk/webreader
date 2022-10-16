// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"webreader/ent/migrate"
	"webreader/ent/schema/ulid"

	"webreader/ent/category"
	"webreader/ent/ranobe"
	"webreader/ent/tag"
	"webreader/ent/todo"
	"webreader/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Ranobe is the client for interacting with the Ranobe builders.
	Ranobe *RanobeClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// Todo is the client for interacting with the Todo builders.
	Todo *TodoClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Category = NewCategoryClient(c.config)
	c.Ranobe = NewRanobeClient(c.config)
	c.Tag = NewTagClient(c.config)
	c.Todo = NewTodoClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Category: NewCategoryClient(cfg),
		Ranobe:   NewRanobeClient(cfg),
		Tag:      NewTagClient(cfg),
		Todo:     NewTodoClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Category: NewCategoryClient(cfg),
		Ranobe:   NewRanobeClient(cfg),
		Tag:      NewTagClient(cfg),
		Todo:     NewTodoClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Category.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Category.Use(hooks...)
	c.Ranobe.Use(hooks...)
	c.Tag.Use(hooks...)
	c.Todo.Use(hooks...)
	c.User.Use(hooks...)
}

// CategoryClient is a client for the Category schema.
type CategoryClient struct {
	config
}

// NewCategoryClient returns a client for the Category from the given config.
func NewCategoryClient(c config) *CategoryClient {
	return &CategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `category.Hooks(f(g(h())))`.
func (c *CategoryClient) Use(hooks ...Hook) {
	c.hooks.Category = append(c.hooks.Category, hooks...)
}

// Create returns a builder for creating a Category entity.
func (c *CategoryClient) Create() *CategoryCreate {
	mutation := newCategoryMutation(c.config, OpCreate)
	return &CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Category entities.
func (c *CategoryClient) CreateBulk(builders ...*CategoryCreate) *CategoryCreateBulk {
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Category.
func (c *CategoryClient) Update() *CategoryUpdate {
	mutation := newCategoryMutation(c.config, OpUpdate)
	return &CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CategoryClient) UpdateOne(ca *Category) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategory(ca))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CategoryClient) UpdateOneID(id ulid.ID) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategoryID(id))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Category.
func (c *CategoryClient) Delete() *CategoryDelete {
	mutation := newCategoryMutation(c.config, OpDelete)
	return &CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CategoryClient) DeleteOne(ca *Category) *CategoryDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CategoryClient) DeleteOneID(id ulid.ID) *CategoryDeleteOne {
	builder := c.Delete().Where(category.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CategoryDeleteOne{builder}
}

// Query returns a query builder for Category.
func (c *CategoryClient) Query() *CategoryQuery {
	return &CategoryQuery{
		config: c.config,
	}
}

// Get returns a Category entity by its id.
func (c *CategoryClient) Get(ctx context.Context, id ulid.ID) (*Category, error) {
	return c.Query().Where(category.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CategoryClient) GetX(ctx context.Context, id ulid.ID) *Category {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRanobes queries the ranobes edge of a Category.
func (c *CategoryClient) QueryRanobes(ca *Category) *RanobeQuery {
	query := &RanobeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(ranobe.Table, ranobe.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, category.RanobesTable, category.RanobesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CategoryClient) Hooks() []Hook {
	return c.hooks.Category
}

// RanobeClient is a client for the Ranobe schema.
type RanobeClient struct {
	config
}

// NewRanobeClient returns a client for the Ranobe from the given config.
func NewRanobeClient(c config) *RanobeClient {
	return &RanobeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ranobe.Hooks(f(g(h())))`.
func (c *RanobeClient) Use(hooks ...Hook) {
	c.hooks.Ranobe = append(c.hooks.Ranobe, hooks...)
}

// Create returns a builder for creating a Ranobe entity.
func (c *RanobeClient) Create() *RanobeCreate {
	mutation := newRanobeMutation(c.config, OpCreate)
	return &RanobeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ranobe entities.
func (c *RanobeClient) CreateBulk(builders ...*RanobeCreate) *RanobeCreateBulk {
	return &RanobeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ranobe.
func (c *RanobeClient) Update() *RanobeUpdate {
	mutation := newRanobeMutation(c.config, OpUpdate)
	return &RanobeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RanobeClient) UpdateOne(r *Ranobe) *RanobeUpdateOne {
	mutation := newRanobeMutation(c.config, OpUpdateOne, withRanobe(r))
	return &RanobeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RanobeClient) UpdateOneID(id ulid.ID) *RanobeUpdateOne {
	mutation := newRanobeMutation(c.config, OpUpdateOne, withRanobeID(id))
	return &RanobeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ranobe.
func (c *RanobeClient) Delete() *RanobeDelete {
	mutation := newRanobeMutation(c.config, OpDelete)
	return &RanobeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RanobeClient) DeleteOne(r *Ranobe) *RanobeDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *RanobeClient) DeleteOneID(id ulid.ID) *RanobeDeleteOne {
	builder := c.Delete().Where(ranobe.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RanobeDeleteOne{builder}
}

// Query returns a query builder for Ranobe.
func (c *RanobeClient) Query() *RanobeQuery {
	return &RanobeQuery{
		config: c.config,
	}
}

// Get returns a Ranobe entity by its id.
func (c *RanobeClient) Get(ctx context.Context, id ulid.ID) (*Ranobe, error) {
	return c.Query().Where(ranobe.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RanobeClient) GetX(ctx context.Context, id ulid.ID) *Ranobe {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCategories queries the categories edge of a Ranobe.
func (c *RanobeClient) QueryCategories(r *Ranobe) *CategoryQuery {
	query := &CategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ranobe.Table, ranobe.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ranobe.CategoriesTable, ranobe.CategoriesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTags queries the tags edge of a Ranobe.
func (c *RanobeClient) QueryTags(r *Ranobe) *TagQuery {
	query := &TagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ranobe.Table, ranobe.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ranobe.TagsTable, ranobe.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RanobeClient) Hooks() []Hook {
	return c.hooks.Ranobe
}

// TagClient is a client for the Tag schema.
type TagClient struct {
	config
}

// NewTagClient returns a client for the Tag from the given config.
func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tag.Hooks(f(g(h())))`.
func (c *TagClient) Use(hooks ...Hook) {
	c.hooks.Tag = append(c.hooks.Tag, hooks...)
}

// Create returns a builder for creating a Tag entity.
func (c *TagClient) Create() *TagCreate {
	mutation := newTagMutation(c.config, OpCreate)
	return &TagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tag entities.
func (c *TagClient) CreateBulk(builders ...*TagCreate) *TagCreateBulk {
	return &TagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tag.
func (c *TagClient) Update() *TagUpdate {
	mutation := newTagMutation(c.config, OpUpdate)
	return &TagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagClient) UpdateOne(t *Tag) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTag(t))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TagClient) UpdateOneID(id ulid.ID) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTagID(id))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tag.
func (c *TagClient) Delete() *TagDelete {
	mutation := newTagMutation(c.config, OpDelete)
	return &TagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TagClient) DeleteOne(t *Tag) *TagDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TagClient) DeleteOneID(id ulid.ID) *TagDeleteOne {
	builder := c.Delete().Where(tag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagDeleteOne{builder}
}

// Query returns a query builder for Tag.
func (c *TagClient) Query() *TagQuery {
	return &TagQuery{
		config: c.config,
	}
}

// Get returns a Tag entity by its id.
func (c *TagClient) Get(ctx context.Context, id ulid.ID) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id ulid.ID) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRanobes queries the ranobes edge of a Tag.
func (c *TagClient) QueryRanobes(t *Tag) *RanobeQuery {
	query := &RanobeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(ranobe.Table, ranobe.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tag.RanobesTable, tag.RanobesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TagClient) Hooks() []Hook {
	return c.hooks.Tag
}

// TodoClient is a client for the Todo schema.
type TodoClient struct {
	config
}

// NewTodoClient returns a client for the Todo from the given config.
func NewTodoClient(c config) *TodoClient {
	return &TodoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `todo.Hooks(f(g(h())))`.
func (c *TodoClient) Use(hooks ...Hook) {
	c.hooks.Todo = append(c.hooks.Todo, hooks...)
}

// Create returns a builder for creating a Todo entity.
func (c *TodoClient) Create() *TodoCreate {
	mutation := newTodoMutation(c.config, OpCreate)
	return &TodoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Todo entities.
func (c *TodoClient) CreateBulk(builders ...*TodoCreate) *TodoCreateBulk {
	return &TodoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Todo.
func (c *TodoClient) Update() *TodoUpdate {
	mutation := newTodoMutation(c.config, OpUpdate)
	return &TodoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TodoClient) UpdateOne(t *Todo) *TodoUpdateOne {
	mutation := newTodoMutation(c.config, OpUpdateOne, withTodo(t))
	return &TodoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TodoClient) UpdateOneID(id ulid.ID) *TodoUpdateOne {
	mutation := newTodoMutation(c.config, OpUpdateOne, withTodoID(id))
	return &TodoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Todo.
func (c *TodoClient) Delete() *TodoDelete {
	mutation := newTodoMutation(c.config, OpDelete)
	return &TodoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TodoClient) DeleteOne(t *Todo) *TodoDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TodoClient) DeleteOneID(id ulid.ID) *TodoDeleteOne {
	builder := c.Delete().Where(todo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TodoDeleteOne{builder}
}

// Query returns a query builder for Todo.
func (c *TodoClient) Query() *TodoQuery {
	return &TodoQuery{
		config: c.config,
	}
}

// Get returns a Todo entity by its id.
func (c *TodoClient) Get(ctx context.Context, id ulid.ID) (*Todo, error) {
	return c.Query().Where(todo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TodoClient) GetX(ctx context.Context, id ulid.ID) *Todo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Todo.
func (c *TodoClient) QueryUser(t *Todo) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, todo.UserTable, todo.UserColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TodoClient) Hooks() []Hook {
	return c.hooks.Todo
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id ulid.ID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id ulid.ID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id ulid.ID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id ulid.ID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTodos queries the todos edge of a User.
func (c *UserClient) QueryTodos(u *User) *TodoQuery {
	query := &TodoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.TodosTable, user.TodosColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
