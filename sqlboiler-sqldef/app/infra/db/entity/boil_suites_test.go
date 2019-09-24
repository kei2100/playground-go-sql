// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Accounts", testAccounts)
	t.Run("Comments", testComments)
	t.Run("Posts", testPosts)
	t.Run("Relationships", testRelationships)
}

func TestDelete(t *testing.T) {
	t.Run("Accounts", testAccountsDelete)
	t.Run("Comments", testCommentsDelete)
	t.Run("Posts", testPostsDelete)
	t.Run("Relationships", testRelationshipsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsQueryDeleteAll)
	t.Run("Comments", testCommentsQueryDeleteAll)
	t.Run("Posts", testPostsQueryDeleteAll)
	t.Run("Relationships", testRelationshipsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceDeleteAll)
	t.Run("Comments", testCommentsSliceDeleteAll)
	t.Run("Posts", testPostsSliceDeleteAll)
	t.Run("Relationships", testRelationshipsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Accounts", testAccountsExists)
	t.Run("Comments", testCommentsExists)
	t.Run("Posts", testPostsExists)
	t.Run("Relationships", testRelationshipsExists)
}

func TestFind(t *testing.T) {
	t.Run("Accounts", testAccountsFind)
	t.Run("Comments", testCommentsFind)
	t.Run("Posts", testPostsFind)
	t.Run("Relationships", testRelationshipsFind)
}

func TestBind(t *testing.T) {
	t.Run("Accounts", testAccountsBind)
	t.Run("Comments", testCommentsBind)
	t.Run("Posts", testPostsBind)
	t.Run("Relationships", testRelationshipsBind)
}

func TestOne(t *testing.T) {
	t.Run("Accounts", testAccountsOne)
	t.Run("Comments", testCommentsOne)
	t.Run("Posts", testPostsOne)
	t.Run("Relationships", testRelationshipsOne)
}

func TestAll(t *testing.T) {
	t.Run("Accounts", testAccountsAll)
	t.Run("Comments", testCommentsAll)
	t.Run("Posts", testPostsAll)
	t.Run("Relationships", testRelationshipsAll)
}

func TestCount(t *testing.T) {
	t.Run("Accounts", testAccountsCount)
	t.Run("Comments", testCommentsCount)
	t.Run("Posts", testPostsCount)
	t.Run("Relationships", testRelationshipsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Accounts", testAccountsHooks)
	t.Run("Comments", testCommentsHooks)
	t.Run("Posts", testPostsHooks)
	t.Run("Relationships", testRelationshipsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Accounts", testAccountsInsert)
	t.Run("Accounts", testAccountsInsertWhitelist)
	t.Run("Comments", testCommentsInsert)
	t.Run("Comments", testCommentsInsertWhitelist)
	t.Run("Posts", testPostsInsert)
	t.Run("Posts", testPostsInsertWhitelist)
	t.Run("Relationships", testRelationshipsInsert)
	t.Run("Relationships", testRelationshipsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CommentToAccountUsingAuthor", testCommentToOneAccountUsingAuthor)
	t.Run("CommentToPostUsingPost", testCommentToOnePostUsingPost)
	t.Run("PostToAccountUsingAuthor", testPostToOneAccountUsingAuthor)
	t.Run("RelationshipToAccountUsingFollowee", testRelationshipToOneAccountUsingFollowee)
	t.Run("RelationshipToAccountUsingFollower", testRelationshipToOneAccountUsingFollower)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AccountToAuthorComments", testAccountToManyAuthorComments)
	t.Run("AccountToAuthorPosts", testAccountToManyAuthorPosts)
	t.Run("AccountToFolloweeRelationships", testAccountToManyFolloweeRelationships)
	t.Run("AccountToFollowerRelationships", testAccountToManyFollowerRelationships)
	t.Run("PostToComments", testPostToManyComments)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CommentToAccountUsingAuthorComments", testCommentToOneSetOpAccountUsingAuthor)
	t.Run("CommentToPostUsingComments", testCommentToOneSetOpPostUsingPost)
	t.Run("PostToAccountUsingAuthorPosts", testPostToOneSetOpAccountUsingAuthor)
	t.Run("RelationshipToAccountUsingFolloweeRelationships", testRelationshipToOneSetOpAccountUsingFollowee)
	t.Run("RelationshipToAccountUsingFollowerRelationships", testRelationshipToOneSetOpAccountUsingFollower)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AccountToAuthorComments", testAccountToManyAddOpAuthorComments)
	t.Run("AccountToAuthorPosts", testAccountToManyAddOpAuthorPosts)
	t.Run("AccountToFolloweeRelationships", testAccountToManyAddOpFolloweeRelationships)
	t.Run("AccountToFollowerRelationships", testAccountToManyAddOpFollowerRelationships)
	t.Run("PostToComments", testPostToManyAddOpComments)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Accounts", testAccountsReload)
	t.Run("Comments", testCommentsReload)
	t.Run("Posts", testPostsReload)
	t.Run("Relationships", testRelationshipsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Accounts", testAccountsReloadAll)
	t.Run("Comments", testCommentsReloadAll)
	t.Run("Posts", testPostsReloadAll)
	t.Run("Relationships", testRelationshipsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Accounts", testAccountsSelect)
	t.Run("Comments", testCommentsSelect)
	t.Run("Posts", testPostsSelect)
	t.Run("Relationships", testRelationshipsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Accounts", testAccountsUpdate)
	t.Run("Comments", testCommentsUpdate)
	t.Run("Posts", testPostsUpdate)
	t.Run("Relationships", testRelationshipsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceUpdateAll)
	t.Run("Comments", testCommentsSliceUpdateAll)
	t.Run("Posts", testPostsSliceUpdateAll)
	t.Run("Relationships", testRelationshipsSliceUpdateAll)
}
