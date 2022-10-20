// Package record wraps the CRUD operations for a models.Record
package record

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.hollow.sh/dnscontroller/internal/models"
)

func qmRecordNameAndType(rname, rtype string) qm.QueryMod {
	mods := []qm.QueryMod{}

	mods = append(mods, qm.Where("record=?", rname))
	mods = append(mods, qm.Where("record_type=?", rtype))

	return qm.Expr(mods...)
}

// Delete removes a record from the DB
func (r *Record) Delete(ctx context.Context, db *sqlx.DB) error {
	err := r.validate()
	if err != nil {
		return err
	}

	err = r.Find(ctx, db)
	if err != nil {
		return err
	}

	dbRecord, err := r.ToDBModel()
	if err != nil {
		return err
	}

	_, err = dbRecord.Delete(ctx, db)
	if err != nil {
		return err
	}

	return nil
}

// FindOrCreate is the upsert function
func (r *Record) FindOrCreate(ctx context.Context, db *sqlx.DB) error {
	err := r.Find(ctx, db)
	if errors.Is(err, sql.ErrNoRows) {
		return r.Create(ctx, db)
	} else if err != nil {
		return err
	}

	return nil
}

// Find looks the record up by name,type
func (r *Record) Find(ctx context.Context, db *sqlx.DB) error {
	if err := r.validate(); err != nil {
		return err
	}

	qm := qmRecordNameAndType(r.Name, r.Type)

	dbRecord, err := models.Records(qm).One(ctx, db)
	if err != nil {
		return err
	}

	return r.FromDBModel(dbRecord)
}

// Create inserts a record
func (r *Record) Create(ctx context.Context, db *sqlx.DB) error {
	dbRecord, err := r.ToDBModel()
	if err != nil {
		return err
	}

	if err := dbRecord.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	// Set the values back
	return r.FromDBModel(dbRecord)
}

// FromDBModel converts a db type to an api type
func (r *Record) FromDBModel(dbT *models.Record) error {
	r.CreatedAt = dbT.CreatedAt
	r.UpdatedAt = dbT.UpdatedAt
	r.Name = dbT.Record
	r.Type = dbT.RecordType

	var err error

	r.UUID, err = uuid.Parse(dbT.ID)
	if err != nil {
		return err
	}

	return r.validate()
}

// ToDBModel converts the api type to db type
func (r *Record) ToDBModel() (*models.Record, error) {
	if err := r.validate(); err != nil {
		return nil, err
	}

	dbModel := &models.Record{
		Record:    r.Name,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}

	switch strings.ToUpper(r.Type) {
	case "A", "CNAME":
		dbModel.RecordType = r.Type
	default:
		return nil, ErrorUnsupportedType
	}

	if r.UUID.String() != uuid.Nil.String() {
		dbModel.ID = r.UUID.String()
	}

	return dbModel, nil
}

// GetPath return the name/type
func (r *Record) GetPath() string {
	return r.path
}

// NewRecord creates a record from the URL params and validates it
func NewRecord(c *gin.Context) (*Record, error) {
	var record *Record

	// Try to get record info from URL params
	rname := c.Param("record")
	rtype := c.Param("recordtype")

	record = &Record{
		Name: rname,
		Type: rtype,
	}

	// Try to validate
	if err := record.validate(); err != nil {
		return nil, err
	}

	// Sanitize input
	record.Name = strings.ToLower(record.Name)
	record.Type = strings.ToUpper(record.Type)
	record.path = record.Name + "/" + record.Type

	return record, nil
}

func (r *Record) validate() error {
	if r.Name == "" {
		return ErrorNoRecordName
	}

	if r.Type == "" {
		return ErrorNoRecordType
	}

	return nil
}
