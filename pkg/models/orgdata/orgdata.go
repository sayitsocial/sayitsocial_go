package orgdata

import (
	"database/sql"

	"github.com/Sayitsocial/Sayitsocial_go/pkg/helpers"
	"github.com/Sayitsocial/Sayitsocial_go/pkg/models"
)

const (
	tableName = helpers.DbTableOrganisationData
	schema    = helpers.DbSchemaOrg
)

// swagger:model
type OrgData struct {
	OrganisationID string `row:"organisation_id" type:"exact" pk:"manual" json:"organisation_id"`
	DisplayName    string `row:"display_name" type:"like" json:"display_name"`
	Locality       string `row:"locality" type:"like" json:"locality"`
	RegistrationNo string `row:"registration_no" type:"exact" json:"registration_no"`
	ContactEmail   string `row:"contact_email" type:"like" json:"contact_email"`
	ContactPhone   string `row:"contact_phone" type:"like" json:"contact_phone"`
	Desc           string `row:"description" type:"like" json:"desc"`
	Owner          string `row:"owner" type:"like" json:"owner"`
	Achievements   string `row:"achievements" type:"like" json:"achievements"`
	TypeOfOrg      int    `row:"type_of_org" type:"like" json:"type_of_org"`
}

type Model struct {
	trans *sql.Tx
	conn  *sql.DB
}

// Initialize returns model of db with active connection
func Initialize(tx *sql.Tx) *Model {
	if tx != nil {
		return &Model{
			trans: tx,
		}
	}
	return &Model{
		conn: models.GetConn(schema, tableName),
	}
}

// Close closes the connection to db
// Model should not be used after close is called
func (a Model) Close() {
	err := a.conn.Close()
	if err != nil {
		helpers.LogError(err.Error())
	}
}

// Create creates a value in database
func (a Model) Create(data OrgData) error {
	query, args := models.QueryBuilderCreate(data, schema+"."+tableName)

	var err error
	if a.trans != nil {
		_, err = a.trans.Exec(query, args...)
	} else {
		_, err = a.conn.Exec(query, args...)
	}
	return err
}

// Get data from db into slice of struct
// Searches by the member provided in input struct
func (a Model) Get(data OrgData) (orgData []OrgData) {
	query, args := models.QueryBuilderGet(data, schema+"."+tableName)

	row, err := a.conn.Query(query, args...)
	if err != nil {
		helpers.LogError(err.Error())
		return
	}

	models.GetIntoStruct(row, &orgData)
	return
}
