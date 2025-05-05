package manager

import (
	"github.com/melkor217/bitmagnet/internal/database/dao"
	"gorm.io/gorm"
)

type manager struct {
	dao *dao.Query
	db  *gorm.DB
}
