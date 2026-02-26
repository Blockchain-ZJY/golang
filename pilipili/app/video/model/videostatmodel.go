package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoStatModel = (*customVideoStatModel)(nil)

type (
	// VideoStatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoStatModel.
	VideoStatModel interface {
		videoStatModel
	}

	customVideoStatModel struct {
		*defaultVideoStatModel
	}
)

// NewVideoStatModel returns a model for the database table.
func NewVideoStatModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VideoStatModel {
	return &customVideoStatModel{
		defaultVideoStatModel: newVideoStatModel(conn, c, opts...),
	}
}
