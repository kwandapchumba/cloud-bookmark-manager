// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: shareCollection.sql

package sqlc

import (
	"context"
)

const getSharedCollection = `-- name: GetSharedCollection :one
SELECT collection_id, collection_shared_by, collection_shared_with, collection_shared_at, collection_access_level FROM shared_collection WHERE collection_id = $1 LIMIT 1
`

func (q *Queries) GetSharedCollection(ctx context.Context, collectionID string) (SharedCollection, error) {
	row := q.db.QueryRowContext(ctx, getSharedCollection, collectionID)
	var i SharedCollection
	err := row.Scan(
		&i.CollectionID,
		&i.CollectionSharedBy,
		&i.CollectionSharedWith,
		&i.CollectionSharedAt,
		&i.CollectionAccessLevel,
	)
	return i, err
}

const getSharedCollectionByCollectionIDandAccountID = `-- name: GetSharedCollectionByCollectionIDandAccountID :one
SELECT collection_id, collection_shared_by, collection_shared_with, collection_shared_at, collection_access_level FROM shared_collection WHERE collection_id = $1 AND collection_shared_with = $2 LIMIT 1
`

type GetSharedCollectionByCollectionIDandAccountIDParams struct {
	CollectionID         string `json:"collection_id"`
	CollectionSharedWith int64  `json:"collection_shared_with"`
}

func (q *Queries) GetSharedCollectionByCollectionIDandAccountID(ctx context.Context, arg GetSharedCollectionByCollectionIDandAccountIDParams) (SharedCollection, error) {
	row := q.db.QueryRowContext(ctx, getSharedCollectionByCollectionIDandAccountID, arg.CollectionID, arg.CollectionSharedWith)
	var i SharedCollection
	err := row.Scan(
		&i.CollectionID,
		&i.CollectionSharedBy,
		&i.CollectionSharedWith,
		&i.CollectionSharedAt,
		&i.CollectionAccessLevel,
	)
	return i, err
}

const shareCollection = `-- name: ShareCollection :one
INSERT INTO shared_collection (collection_id, collection_shared_by, collection_shared_with, collection_access_level)
VALUES ($1, $2, $3, $4)
RETURNING collection_id, collection_shared_by, collection_shared_with, collection_shared_at, collection_access_level
`

type ShareCollectionParams struct {
	CollectionID          string                `json:"collection_id"`
	CollectionSharedBy    int64                 `json:"collection_shared_by"`
	CollectionSharedWith  int64                 `json:"collection_shared_with"`
	CollectionAccessLevel CollectionAccessLevel `json:"collection_access_level"`
}

func (q *Queries) ShareCollection(ctx context.Context, arg ShareCollectionParams) (SharedCollection, error) {
	row := q.db.QueryRowContext(ctx, shareCollection,
		arg.CollectionID,
		arg.CollectionSharedBy,
		arg.CollectionSharedWith,
		arg.CollectionAccessLevel,
	)
	var i SharedCollection
	err := row.Scan(
		&i.CollectionID,
		&i.CollectionSharedBy,
		&i.CollectionSharedWith,
		&i.CollectionSharedAt,
		&i.CollectionAccessLevel,
	)
	return i, err
}