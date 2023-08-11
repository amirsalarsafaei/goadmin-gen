package gohadmin

import (
	"context"

	adminContext "github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type TableController interface {
	Create(ctx context.Context, data form.Values, user *models.UserModel) error
	Retrieve(ctx context.Context, params parameter.Parameters, user *models.UserModel) ([]map[string]interface{}, int)
	Update(ctx context.Context, data form.Values, user *models.UserModel) error
	Delete(ctx context.Context, id string, user *models.UserModel) error
}

type Table interface {
	GetTable(ctx *adminContext.Context) table.Table
}

type CustomTable struct {
	*table.DefaultTable
	user *models.UserModel
	Conn db.Connection

	ctx        context.Context
	controller TableController
}

type CustomTableConfig struct {
	*table.Config
	User       *models.UserModel
	Conn       db.Connection
	Controller TableController
}

func NewCustomTable(ctx context.Context, cfg CustomTableConfig) table.Table {
	return &CustomTable{
		DefaultTable: table.NewDefaultTable(*cfg.Config).(*table.DefaultTable),
		user:         cfg.User,
		Conn:         cfg.Conn,
		controller:   cfg.Controller,
		ctx:          ctx,
	}
}

func (tb *CustomTable) GetNewFormInfo() table.FormInfo {
	f := tb.GetActualNewForm()

	return table.FormInfo{FieldList: f.FieldsWithDefaultValue(tb.sqlObjOrNil)}
}

func (tb *CustomTable) sqlObjOrNil() *db.SQL {
	return db.WithDriverAndConnection("default", tb.Conn)
}

func getDataRes(list []map[string]interface{}, _ int) map[string]interface{} {
	if len(list) > 0 {
		return list[0]
	}
	return nil
}

func (tb *CustomTable) GetDataWithId(param parameter.Parameters) (table.FormInfo, error) { //nolint:revive

	var (
		res     map[string]interface{}
		columns table.Columns
		id      = param.PK()
	)

	res = getDataRes(tb.controller.Retrieve(tb.ctx, param, tb.user))

	var (
		groupFormList = make([]types.FormFields, 0)
		groupHeaders  = make([]string, 0)
	)

	if len(tb.Form.TabGroups) > 0 {
		groupFormList, groupHeaders = tb.Form.GroupFieldWithValue(tb.PrimaryKey.Name, id, columns, res, tb.sqlObjOrNil)
		return table.FormInfo{
			FieldList:         tb.Form.FieldList,
			GroupFieldList:    groupFormList,
			GroupFieldHeaders: groupHeaders,
			Title:             tb.Form.Title,
			Description:       tb.Form.Description,
		}, nil
	}

	var fieldList = tb.Form.FieldsWithValue(tb.PrimaryKey.Name, id, columns, res, tb.sqlObjOrNil)

	return table.FormInfo{
		FieldList:         fieldList,
		GroupFieldList:    groupFormList,
		GroupFieldHeaders: groupHeaders,
		Title:             tb.Form.Title,
		Description:       tb.Form.Description,
	}, nil
}

func (tb *CustomTable) UpdateData(dataList form.Values) error {
	return tb.controller.Update(tb.ctx, dataList, tb.user)
}

func (tb *CustomTable) InsertData(dataList form.Values) error {
	return tb.controller.Create(tb.ctx, dataList, tb.user)
}

func (tb *CustomTable) DeleteData(id string) error {
	return tb.controller.Delete(tb.ctx, id, tb.user)
}
