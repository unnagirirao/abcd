package services

import (
	"github.com/unnagirirao/abcd/bcc/pkg/rest/server/daos"
	"github.com/unnagirirao/abcd/bcc/pkg/rest/server/models"
)

type CbncService struct {
	cbncDao *daos.CbncDao
}

func NewCbncService() (*CbncService, error) {
	cbncDao, err := daos.NewCbncDao()
	if err != nil {
		return nil, err
	}
	return &CbncService{
		cbncDao: cbncDao,
	}, nil
}

func (cbncService *CbncService) CreateCbnc(cbnc *models.Cbnc) (*models.Cbnc, error) {
	return cbncService.cbncDao.CreateCbnc(cbnc)
}

func (cbncService *CbncService) UpdateCbnc(id int64, cbnc *models.Cbnc) (*models.Cbnc, error) {
	return cbncService.cbncDao.UpdateCbnc(id, cbnc)
}

func (cbncService *CbncService) DeleteCbnc(id int64) error {
	return cbncService.cbncDao.DeleteCbnc(id)
}

func (cbncService *CbncService) ListCbncs() ([]*models.Cbnc, error) {
	return cbncService.cbncDao.ListCbncs()
}

func (cbncService *CbncService) GetCbnc(id int64) (*models.Cbnc, error) {
	return cbncService.cbncDao.GetCbnc(id)
}
