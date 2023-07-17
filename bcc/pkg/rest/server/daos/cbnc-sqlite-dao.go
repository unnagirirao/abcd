package daos

import (
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/unnagirirao/abcd/bcc/pkg/rest/server/daos/clients/sqls"
	"github.com/unnagirirao/abcd/bcc/pkg/rest/server/models"
)

type CbncDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateCbncs(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS cbncs(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Can INTEGER NOT NULL,
		Canot TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewCbncDao() (*CbncDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateCbncs(sqlClient)
	if err != nil {
		return nil, err
	}
	return &CbncDao{
		sqlClient,
	}, nil
}

func (cbncDao *CbncDao) CreateCbnc(m *models.Cbnc) (*models.Cbnc, error) {
	insertQuery := "INSERT INTO cbncs(Can, Canot)values(?, ?)"
	res, err := cbncDao.sqlClient.DB.Exec(insertQuery, m.Can, m.Canot)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("cbnc created")
	return m, nil
}

func (cbncDao *CbncDao) UpdateCbnc(id int64, m *models.Cbnc) (*models.Cbnc, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	cbnc, err := cbncDao.GetCbnc(id)
	if err != nil {
		return nil, err
	}
	if cbnc == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE cbncs SET Can = ?, Canot = ? WHERE Id = ?"
	res, err := cbncDao.sqlClient.DB.Exec(updateQuery, m.Can, m.Canot, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("cbnc updated")
	return m, nil
}

func (cbncDao *CbncDao) DeleteCbnc(id int64) error {
	deleteQuery := "DELETE FROM cbncs WHERE Id = ?"
	res, err := cbncDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("cbnc deleted")
	return nil
}

func (cbncDao *CbncDao) ListCbncs() ([]*models.Cbnc, error) {
	selectQuery := "SELECT * FROM cbncs"
	rows, err := cbncDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var cbncs []*models.Cbnc
	for rows.Next() {
		m := models.Cbnc{}
		if err = rows.Scan(&m.Id, &m.Can, &m.Canot); err != nil {
			return nil, err
		}
		cbncs = append(cbncs, &m)
	}
	if cbncs == nil {
		cbncs = []*models.Cbnc{}
	}

	log.Debugf("cbnc listed")
	return cbncs, nil
}

func (cbncDao *CbncDao) GetCbnc(id int64) (*models.Cbnc, error) {
	selectQuery := "SELECT * FROM cbncs WHERE Id = ?"
	row := cbncDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Cbnc{}
	if err := row.Scan(&m.Id, &m.Can, &m.Canot); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("cbnc retrieved")
	return &m, nil
}
