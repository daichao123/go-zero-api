package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CountriesModel = (*customCountriesModel)(nil)

type (
	// CountriesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCountriesModel.
	CountriesModel interface {
		countriesModel
	}

	customCountriesModel struct {
		*defaultCountriesModel
	}
)

// NewCountriesModel returns a model for the database table.
func NewCountriesModel(conn sqlx.SqlConn) CountriesModel {
	return &customCountriesModel{
		defaultCountriesModel: newCountriesModel(conn),
	}
}
