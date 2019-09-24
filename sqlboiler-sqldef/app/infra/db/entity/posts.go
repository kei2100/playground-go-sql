// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Post is an object representing the database table.
type Post struct {
	ID       string `boil:"id" json:"id" toml:"id" yaml:"id"`
	AuthorID string `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	Body     string `boil:"body" json:"body" toml:"body" yaml:"body"`

	R *postR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L postL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PostColumns = struct {
	ID       string
	AuthorID string
	Body     string
}{
	ID:       "id",
	AuthorID: "author_id",
	Body:     "body",
}

// Generated where

var PostWhere = struct {
	ID       whereHelperstring
	AuthorID whereHelperstring
	Body     whereHelperstring
}{
	ID:       whereHelperstring{field: "\"posts\".\"id\""},
	AuthorID: whereHelperstring{field: "\"posts\".\"author_id\""},
	Body:     whereHelperstring{field: "\"posts\".\"body\""},
}

// PostRels is where relationship names are stored.
var PostRels = struct {
	Author   string
	Comments string
}{
	Author:   "Author",
	Comments: "Comments",
}

// postR is where relationships are stored.
type postR struct {
	Author   *Account
	Comments CommentSlice
}

// NewStruct creates a new relationship struct
func (*postR) NewStruct() *postR {
	return &postR{}
}

// postL is where Load methods for each relationship are stored.
type postL struct{}

var (
	postAllColumns            = []string{"id", "author_id", "body"}
	postColumnsWithoutDefault = []string{"id", "author_id", "body"}
	postColumnsWithDefault    = []string{}
	postPrimaryKeyColumns     = []string{"id"}
)

type (
	// PostSlice is an alias for a slice of pointers to Post.
	// This should generally be used opposed to []Post.
	PostSlice []*Post
	// PostHook is the signature for custom Post hook methods
	PostHook func(context.Context, boil.ContextExecutor, *Post) error

	postQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	postType                 = reflect.TypeOf(&Post{})
	postMapping              = queries.MakeStructMapping(postType)
	postPrimaryKeyMapping, _ = queries.BindMapping(postType, postMapping, postPrimaryKeyColumns)
	postInsertCacheMut       sync.RWMutex
	postInsertCache          = make(map[string]insertCache)
	postUpdateCacheMut       sync.RWMutex
	postUpdateCache          = make(map[string]updateCache)
	postUpsertCacheMut       sync.RWMutex
	postUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var postBeforeInsertHooks []PostHook
var postBeforeUpdateHooks []PostHook
var postBeforeDeleteHooks []PostHook
var postBeforeUpsertHooks []PostHook

var postAfterInsertHooks []PostHook
var postAfterSelectHooks []PostHook
var postAfterUpdateHooks []PostHook
var postAfterDeleteHooks []PostHook
var postAfterUpsertHooks []PostHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Post) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Post) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Post) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Post) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Post) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Post) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Post) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Post) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Post) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range postAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPostHook registers your hook function for all future operations.
func AddPostHook(hookPoint boil.HookPoint, postHook PostHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		postBeforeInsertHooks = append(postBeforeInsertHooks, postHook)
	case boil.BeforeUpdateHook:
		postBeforeUpdateHooks = append(postBeforeUpdateHooks, postHook)
	case boil.BeforeDeleteHook:
		postBeforeDeleteHooks = append(postBeforeDeleteHooks, postHook)
	case boil.BeforeUpsertHook:
		postBeforeUpsertHooks = append(postBeforeUpsertHooks, postHook)
	case boil.AfterInsertHook:
		postAfterInsertHooks = append(postAfterInsertHooks, postHook)
	case boil.AfterSelectHook:
		postAfterSelectHooks = append(postAfterSelectHooks, postHook)
	case boil.AfterUpdateHook:
		postAfterUpdateHooks = append(postAfterUpdateHooks, postHook)
	case boil.AfterDeleteHook:
		postAfterDeleteHooks = append(postAfterDeleteHooks, postHook)
	case boil.AfterUpsertHook:
		postAfterUpsertHooks = append(postAfterUpsertHooks, postHook)
	}
}

// One returns a single post record from the query.
func (q postQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Post, error) {
	o := &Post{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for posts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Post records from the query.
func (q postQuery) All(ctx context.Context, exec boil.ContextExecutor) (PostSlice, error) {
	var o []*Post

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Post slice")
	}

	if len(postAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Post records in the query.
func (q postQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count posts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q postQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if posts exists")
	}

	return count > 0, nil
}

// Author pointed to by the foreign key.
func (o *Post) Author(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AuthorID),
	}

	queryMods = append(queryMods, mods...)

	query := Accounts(queryMods...)
	queries.SetFrom(query.Query, "\"accounts\"")

	return query
}

// Comments retrieves all the comment's Comments with an executor.
func (o *Post) Comments(mods ...qm.QueryMod) commentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"comments\".\"post_id\"=?", o.ID),
	)

	query := Comments(queryMods...)
	queries.SetFrom(query.Query, "\"comments\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"comments\".*"})
	}

	return query
}

// LoadAuthor allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (postL) LoadAuthor(ctx context.Context, e boil.ContextExecutor, singular bool, maybePost interface{}, mods queries.Applicator) error {
	var slice []*Post
	var object *Post

	if singular {
		object = maybePost.(*Post)
	} else {
		slice = *maybePost.(*[]*Post)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &postR{}
		}
		args = append(args, object.AuthorID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &postR{}
			}

			for _, a := range args {
				if a == obj.AuthorID {
					continue Outer
				}
			}

			args = append(args, obj.AuthorID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`accounts`), qm.WhereIn(`accounts.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Account")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Account")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for accounts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for accounts")
	}

	if len(postAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Author = foreign
		if foreign.R == nil {
			foreign.R = &accountR{}
		}
		foreign.R.AuthorPosts = append(foreign.R.AuthorPosts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AuthorID == foreign.ID {
				local.R.Author = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.AuthorPosts = append(foreign.R.AuthorPosts, local)
				break
			}
		}
	}

	return nil
}

// LoadComments allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (postL) LoadComments(ctx context.Context, e boil.ContextExecutor, singular bool, maybePost interface{}, mods queries.Applicator) error {
	var slice []*Post
	var object *Post

	if singular {
		object = maybePost.(*Post)
	} else {
		slice = *maybePost.(*[]*Post)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &postR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &postR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`comments`), qm.WhereIn(`comments.post_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load comments")
	}

	var resultSlice []*Comment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice comments")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on comments")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for comments")
	}

	if len(commentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Comments = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &commentR{}
			}
			foreign.R.Post = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PostID {
				local.R.Comments = append(local.R.Comments, foreign)
				if foreign.R == nil {
					foreign.R = &commentR{}
				}
				foreign.R.Post = local
				break
			}
		}
	}

	return nil
}

// SetAuthor of the post to the related item.
// Sets o.R.Author to related.
// Adds o to related.R.AuthorPosts.
func (o *Post) SetAuthor(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"posts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"author_id"}),
		strmangle.WhereClause("\"", "\"", 2, postPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AuthorID = related.ID
	if o.R == nil {
		o.R = &postR{
			Author: related,
		}
	} else {
		o.R.Author = related
	}

	if related.R == nil {
		related.R = &accountR{
			AuthorPosts: PostSlice{o},
		}
	} else {
		related.R.AuthorPosts = append(related.R.AuthorPosts, o)
	}

	return nil
}

// AddComments adds the given related objects to the existing relationships
// of the post, optionally inserting them as new records.
// Appends related to o.R.Comments.
// Sets related.R.Post appropriately.
func (o *Post) AddComments(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Comment) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PostID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"comments\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"post_id"}),
				strmangle.WhereClause("\"", "\"", 2, commentPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PostID = o.ID
		}
	}

	if o.R == nil {
		o.R = &postR{
			Comments: related,
		}
	} else {
		o.R.Comments = append(o.R.Comments, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &commentR{
				Post: o,
			}
		} else {
			rel.R.Post = o
		}
	}
	return nil
}

// Posts retrieves all the records using an executor.
func Posts(mods ...qm.QueryMod) postQuery {
	mods = append(mods, qm.From("\"posts\""))
	return postQuery{NewQuery(mods...)}
}

// FindPost retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPost(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Post, error) {
	postObj := &Post{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"posts\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, postObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from posts")
	}

	return postObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Post) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no posts provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(postColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	postInsertCacheMut.RLock()
	cache, cached := postInsertCache[key]
	postInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			postAllColumns,
			postColumnsWithDefault,
			postColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(postType, postMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(postType, postMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"posts\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"posts\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "entity: unable to insert into posts")
	}

	if !cached {
		postInsertCacheMut.Lock()
		postInsertCache[key] = cache
		postInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Post.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Post) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	postUpdateCacheMut.RLock()
	cache, cached := postUpdateCache[key]
	postUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			postAllColumns,
			postPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update posts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"posts\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, postPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(postType, postMapping, append(wl, postPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update posts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for posts")
	}

	if !cached {
		postUpdateCacheMut.Lock()
		postUpdateCache[key] = cache
		postUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q postQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for posts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for posts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PostSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("entity: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"posts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, postPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in post slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all post")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Post) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no posts provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(postColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	postUpsertCacheMut.RLock()
	cache, cached := postUpsertCache[key]
	postUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			postAllColumns,
			postColumnsWithDefault,
			postColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			postAllColumns,
			postPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert posts, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(postPrimaryKeyColumns))
			copy(conflict, postPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"posts\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(postType, postMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(postType, postMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "entity: unable to upsert posts")
	}

	if !cached {
		postUpsertCacheMut.Lock()
		postUpsertCache[key] = cache
		postUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Post record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Post) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no Post provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), postPrimaryKeyMapping)
	sql := "DELETE FROM \"posts\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from posts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for posts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q postQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no postQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from posts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for posts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PostSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(postBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"posts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, postPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from post slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for posts")
	}

	if len(postAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Post) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPost(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PostSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PostSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"posts\".* FROM \"posts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, postPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in PostSlice")
	}

	*o = slice

	return nil
}

// PostExists checks if the Post row exists.
func PostExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"posts\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if posts exists")
	}

	return exists, nil
}
