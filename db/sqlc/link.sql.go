// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: link.sql

package sqlc

import (
	"context"
	"database/sql"
)

const addLink = `-- name: AddLink :one
INSERT INTO link (link_id, link_title, link_hostname, link_url, link_favicon, account_id, folder_id, link_thumbnail)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col
`

type AddLinkParams struct {
	LinkID        string         `json:"link_id"`
	LinkTitle     string         `json:"link_title"`
	LinkHostname  string         `json:"link_hostname"`
	LinkUrl       string         `json:"link_url"`
	LinkFavicon   string         `json:"link_favicon"`
	AccountID     int64          `json:"account_id"`
	FolderID      sql.NullString `json:"folder_id"`
	LinkThumbnail string         `json:"link_thumbnail"`
}

func (q *Queries) AddLink(ctx context.Context, arg AddLinkParams) (Link, error) {
	row := q.db.QueryRowContext(ctx, addLink,
		arg.LinkID,
		arg.LinkTitle,
		arg.LinkHostname,
		arg.LinkUrl,
		arg.LinkFavicon,
		arg.AccountID,
		arg.FolderID,
		arg.LinkThumbnail,
	)
	var i Link
	err := row.Scan(
		&i.LinkID,
		&i.LinkTitle,
		&i.LinkThumbnail,
		&i.LinkFavicon,
		&i.LinkHostname,
		&i.LinkUrl,
		&i.LinkNotes,
		&i.AccountID,
		&i.FolderID,
		&i.AddedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.TextsearchableIndexCol,
	)
	return i, err
}

const deleteLinkForever = `-- name: DeleteLinkForever :one
DELETE FROM link WHERE link_id = $1 RETURNING link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col
`

func (q *Queries) DeleteLinkForever(ctx context.Context, linkID string) (Link, error) {
	row := q.db.QueryRowContext(ctx, deleteLinkForever, linkID)
	var i Link
	err := row.Scan(
		&i.LinkID,
		&i.LinkTitle,
		&i.LinkThumbnail,
		&i.LinkFavicon,
		&i.LinkHostname,
		&i.LinkUrl,
		&i.LinkNotes,
		&i.AccountID,
		&i.FolderID,
		&i.AddedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.TextsearchableIndexCol,
	)
	return i, err
}

const getFolderLinks = `-- name: GetFolderLinks :many
SELECT link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col FROM link WHERE account_id = $1 AND folder_id = $2 AND deleted_at IS NULL ORDER BY added_at DESC
`

type GetFolderLinksParams struct {
	AccountID int64          `json:"account_id"`
	FolderID  sql.NullString `json:"folder_id"`
}

func (q *Queries) GetFolderLinks(ctx context.Context, arg GetFolderLinksParams) ([]Link, error) {
	rows, err := q.db.QueryContext(ctx, getFolderLinks, arg.AccountID, arg.FolderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.LinkID,
			&i.LinkTitle,
			&i.LinkThumbnail,
			&i.LinkFavicon,
			&i.LinkHostname,
			&i.LinkUrl,
			&i.LinkNotes,
			&i.AccountID,
			&i.FolderID,
			&i.AddedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TextsearchableIndexCol,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLink = `-- name: GetLink :one
SELECT link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col FROM link
WHERE link_id = $1
LIMIT 1
`

func (q *Queries) GetLink(ctx context.Context, linkID string) (Link, error) {
	row := q.db.QueryRowContext(ctx, getLink, linkID)
	var i Link
	err := row.Scan(
		&i.LinkID,
		&i.LinkTitle,
		&i.LinkThumbnail,
		&i.LinkFavicon,
		&i.LinkHostname,
		&i.LinkUrl,
		&i.LinkNotes,
		&i.AccountID,
		&i.FolderID,
		&i.AddedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.TextsearchableIndexCol,
	)
	return i, err
}

const getLinksByUserID = `-- name: GetLinksByUserID :many
SELECT link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col FROM link WHERE account_id = $1
`

func (q *Queries) GetLinksByUserID(ctx context.Context, accountID int64) ([]Link, error) {
	rows, err := q.db.QueryContext(ctx, getLinksByUserID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.LinkID,
			&i.LinkTitle,
			&i.LinkThumbnail,
			&i.LinkFavicon,
			&i.LinkHostname,
			&i.LinkUrl,
			&i.LinkNotes,
			&i.AccountID,
			&i.FolderID,
			&i.AddedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TextsearchableIndexCol,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLinksMovedToTrash = `-- name: GetLinksMovedToTrash :many
SELECT link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col FROM link WHERE deleted_at IS NOT NULL AND account_id = $1 ORDER BY deleted_at DESC
`

func (q *Queries) GetLinksMovedToTrash(ctx context.Context, accountID int64) ([]Link, error) {
	rows, err := q.db.QueryContext(ctx, getLinksMovedToTrash, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.LinkID,
			&i.LinkTitle,
			&i.LinkThumbnail,
			&i.LinkFavicon,
			&i.LinkHostname,
			&i.LinkUrl,
			&i.LinkNotes,
			&i.AccountID,
			&i.FolderID,
			&i.AddedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TextsearchableIndexCol,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRootLinks = `-- name: GetRootLinks :many
SELECT link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col FROM link WHERE account_id = $1 AND folder_id IS NULL AND deleted_at IS NULL ORDER BY added_at DESC
`

func (q *Queries) GetRootLinks(ctx context.Context, accountID int64) ([]Link, error) {
	rows, err := q.db.QueryContext(ctx, getRootLinks, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.LinkID,
			&i.LinkTitle,
			&i.LinkThumbnail,
			&i.LinkFavicon,
			&i.LinkHostname,
			&i.LinkUrl,
			&i.LinkNotes,
			&i.AccountID,
			&i.FolderID,
			&i.AddedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TextsearchableIndexCol,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const moveLinkToFolder = `-- name: MoveLinkToFolder :one
UPDATE link SET folder_id = $1 WHERE link_id = $2 RETURNING link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col
`

type MoveLinkToFolderParams struct {
	FolderID sql.NullString `json:"folder_id"`
	LinkID   string         `json:"link_id"`
}

func (q *Queries) MoveLinkToFolder(ctx context.Context, arg MoveLinkToFolderParams) (Link, error) {
	row := q.db.QueryRowContext(ctx, moveLinkToFolder, arg.FolderID, arg.LinkID)
	var i Link
	err := row.Scan(
		&i.LinkID,
		&i.LinkTitle,
		&i.LinkThumbnail,
		&i.LinkFavicon,
		&i.LinkHostname,
		&i.LinkUrl,
		&i.LinkNotes,
		&i.AccountID,
		&i.FolderID,
		&i.AddedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.TextsearchableIndexCol,
	)
	return i, err
}

const moveLinkToRoot = `-- name: MoveLinkToRoot :one
UPDATE link SET folder_id = NULL WHERE link_id = $1 RETURNING link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col
`

func (q *Queries) MoveLinkToRoot(ctx context.Context, linkID string) (Link, error) {
	row := q.db.QueryRowContext(ctx, moveLinkToRoot, linkID)
	var i Link
	err := row.Scan(
		&i.LinkID,
		&i.LinkTitle,
		&i.LinkThumbnail,
		&i.LinkFavicon,
		&i.LinkHostname,
		&i.LinkUrl,
		&i.LinkNotes,
		&i.AccountID,
		&i.FolderID,
		&i.AddedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.TextsearchableIndexCol,
	)
	return i, err
}

const moveLinkToTrash = `-- name: MoveLinkToTrash :one
UPDATE link SET deleted_at = CURRENT_TIMESTAMP WHERE link_id = $1 RETURNING link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col
`

func (q *Queries) MoveLinkToTrash(ctx context.Context, linkID string) (Link, error) {
	row := q.db.QueryRowContext(ctx, moveLinkToTrash, linkID)
	var i Link
	err := row.Scan(
		&i.LinkID,
		&i.LinkTitle,
		&i.LinkThumbnail,
		&i.LinkFavicon,
		&i.LinkHostname,
		&i.LinkUrl,
		&i.LinkNotes,
		&i.AccountID,
		&i.FolderID,
		&i.AddedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.TextsearchableIndexCol,
	)
	return i, err
}

const renameLink = `-- name: RenameLink :one
UPDATE link SET link_title = $1 WHERE link_id = $2 RETURNING link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col
`

type RenameLinkParams struct {
	LinkTitle string `json:"link_title"`
	LinkID    string `json:"link_id"`
}

func (q *Queries) RenameLink(ctx context.Context, arg RenameLinkParams) (Link, error) {
	row := q.db.QueryRowContext(ctx, renameLink, arg.LinkTitle, arg.LinkID)
	var i Link
	err := row.Scan(
		&i.LinkID,
		&i.LinkTitle,
		&i.LinkThumbnail,
		&i.LinkFavicon,
		&i.LinkHostname,
		&i.LinkUrl,
		&i.LinkNotes,
		&i.AccountID,
		&i.FolderID,
		&i.AddedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.TextsearchableIndexCol,
	)
	return i, err
}

const restoreLinkFromTrash = `-- name: RestoreLinkFromTrash :one
UPDATE link SET deleted_at = NULL WHERE link_id = $1 RETURNING link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col
`

func (q *Queries) RestoreLinkFromTrash(ctx context.Context, linkID string) (Link, error) {
	row := q.db.QueryRowContext(ctx, restoreLinkFromTrash, linkID)
	var i Link
	err := row.Scan(
		&i.LinkID,
		&i.LinkTitle,
		&i.LinkThumbnail,
		&i.LinkFavicon,
		&i.LinkHostname,
		&i.LinkUrl,
		&i.LinkNotes,
		&i.AccountID,
		&i.FolderID,
		&i.AddedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.TextsearchableIndexCol,
	)
	return i, err
}

const searchLinks = `-- name: SearchLinks :many
SELECT link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col
FROM link
WHERE textsearchable_index_col @@ plainto_tsquery($1) AND account_id = $2 AND deleted_at IS NULL
ORDER BY added_at DESC
`

type SearchLinksParams struct {
	PlaintoTsquery string `json:"plainto_tsquery"`
	AccountID      int64  `json:"account_id"`
}

func (q *Queries) SearchLinks(ctx context.Context, arg SearchLinksParams) ([]Link, error) {
	rows, err := q.db.QueryContext(ctx, searchLinks, arg.PlaintoTsquery, arg.AccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.LinkID,
			&i.LinkTitle,
			&i.LinkThumbnail,
			&i.LinkFavicon,
			&i.LinkHostname,
			&i.LinkUrl,
			&i.LinkNotes,
			&i.AccountID,
			&i.FolderID,
			&i.AddedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TextsearchableIndexCol,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchLinkz = `-- name: SearchLinkz :many
SELECT link_id, link_title, link_thumbnail, link_favicon, link_hostname, link_url, link_notes, account_id, folder_id, added_at, updated_at, deleted_at, textsearchable_index_col
FROM link
WHERE link_title ILIKE $1 AND account_id = $2 AND deleted_at IS NULL
ORDER BY added_at DESC
`

type SearchLinkzParams struct {
	LinkTitle string `json:"link_title"`
	AccountID int64  `json:"account_id"`
}

func (q *Queries) SearchLinkz(ctx context.Context, arg SearchLinkzParams) ([]Link, error) {
	rows, err := q.db.QueryContext(ctx, searchLinkz, arg.LinkTitle, arg.AccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.LinkID,
			&i.LinkTitle,
			&i.LinkThumbnail,
			&i.LinkFavicon,
			&i.LinkHostname,
			&i.LinkUrl,
			&i.LinkNotes,
			&i.AccountID,
			&i.FolderID,
			&i.AddedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TextsearchableIndexCol,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
