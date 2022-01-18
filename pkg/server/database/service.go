package database

import (
	servicev1 "github.com/NexClipper/sudory/pkg/server/model/service/v1"
)

func (d *DBManipulator) CreateService(m servicev1.DbSchemaService) error {
	var err error
	tx := d.session()
	tx.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	affect, err := tx.Insert(&m)
	if err != nil {
		return err
	}
	if !(0 < affect) {
		return ErrorNoAffecte()
	}
	return nil
}

/* GetService
   @return DbSchemaService, error
   @method get
   @from Service
   @condition uuid
*/
func (d *DBManipulator) GetService(uuid string) (*servicev1.DbSchemaService, error) {
	tx := d.session()

	record := new(servicev1.DbSchemaService)
	has, err := tx.Where("uuid = ?", uuid).
		Get(record)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrorRecordWasNotFound()
	}

	return record, err
}

/* FindService
   @return []servicev1.DbSchemaService, error
   @method find
   @from Service
   @condition where, args
*/
func (d *DBManipulator) FindService(where string, args ...interface{}) ([]servicev1.DbSchemaService, error) {
	tx := d.session()

	//SELECT * FROM {table} WHERE [cond]
	var model = make([]servicev1.DbSchemaService, 0)
	err := tx.Where(where, args...).
		Find(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

/* UpdateService
   @return error
   @method update
   @from Service
   @condition DbSchemaService
*/
func (d *DBManipulator) UpdateService(m servicev1.DbSchemaService) error {
	var err error
	tx := d.session()
	tx.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	affect, err := tx.Where("uuid = ?", m.Uuid).
		Update(&m)
	if err != nil {
		return err
	}
	if !(0 < affect) {
		return ErrorNoAffecte()
	}
	return nil
}

/* DeleteService
   @return error
   @method delete
   @from Service
   @condition uuid
*/
func (d *DBManipulator) DeleteService(uuid string) error {
	var err error
	tx := d.session()
	tx.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	record := new(servicev1.DbSchemaService)
	//DELETE FROM {table} WHERE uuid = ?
	affect, err := tx.Where("uuid = ?", uuid).
		Delete(record)
	if err != nil {
		return err
	}
	if !(0 < affect) {
		return ErrorNoAffecte()
	}
	return nil
}
