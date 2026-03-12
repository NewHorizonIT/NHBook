package helper

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func NullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

func NullUUID(id *uuid.UUID) uuid.NullUUID {
	if id == nil {
		return uuid.NullUUID{Valid: false}
	}
	return uuid.NullUUID{UUID: *id, Valid: true}
}

func NullBool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}

func PtrString(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func PtrUUID(nu uuid.NullUUID) *uuid.UUID {
	if !nu.Valid {
		return nil
	}
	return &nu.UUID
}

func PtrBool(nb sql.NullBool) bool {
	if !nb.Valid {
		return false
	}
	return nb.Bool
}

func PtrTime(nt sql.NullTime) time.Time {
	if !nt.Valid {
		return time.Time{}
	}
	return nt.Time
}
