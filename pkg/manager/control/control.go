package control

import (
	"database/sql"

	"github.com/jaehoonkim/sentinel/pkg/manager/database/vanilla/excute"

	"github.com/labstack/echo/v4"
)

// type Control struct {
// 	db *database.DBManipulator
// }

// func New(d *database.DBManipulator) *Control {
// 	return &Control{db: d}
// }

// func (ctl Control) Scope(fn func(database.Context) (interface{}, error)) (v interface{}, err error) {
// 	block.Block{
// 		Try: func() {
// 			_, lockerr := ctl.db.Engine().Transaction(func(s *xorm.Session) (interface{}, error) {
// 				v, err = fn(database.NewXormContext(s))
// 				return nil, err
// 			})
// 			if err == nil && lockerr != nil {
// 				err = errors.Wrapf(lockerr, "xorm commit")
// 			}
// 		},
// 		Catch: func(ex error) {
// 			err = errors.Wrapf(ex, "catch")
// 		},
// 	}.Do()

// 	return
// }

// func (ctl Control) ScopeSession(fn func(tx *xorm.Session) (interface{}, error)) (v interface{}, err error) {
// 	block.Block{
// 		Try: func() {
// 			_, lockerr := ctl.db.Engine().Transaction(func(s *xorm.Session) (interface{}, error) {
// 				v, err = fn(s)
// 				return nil, err
// 			})
// 			if err == nil && lockerr != nil {
// 				err = errors.Wrapf(lockerr, "xorm commit")
// 			}
// 		},
// 		Catch: func(ex error) {
// 			err = errors.Wrapf(ex, "catch")
// 		},
// 	}.Do()

// 	return
// }

// func (ctl Control) NewSession() database.Context {
// 	return database.NewXormContext(ctl.db.Engine().NewSession())
// }

type ControlVanilla struct {
	*sql.DB
	dialect excute.SqlExcutor
}

func NewVanilla(db *sql.DB, dialect excute.SqlExcutor) *ControlVanilla {
	return &ControlVanilla{
		DB:      db,
		dialect: dialect,
	}
}

// func (ctl ControlVanilla) Dialect() string {
// 	return ctl.dialect
// }

// HttpError
func HttpError(err error, code int) error {
	if err == nil {
		return nil
	}
	return echo.NewHTTPError(code).SetInternal(err)
}
