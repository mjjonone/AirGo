package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"github.com/gin-gonic/gin"
)

// 获取数据库的所有数据库名
func GetDB(ctx *gin.Context) {
	//var database model.DbNameAndTableReq
	//err := ctx.ShouldBind(&database)
	//if err != nil {
	//	global.Logrus.Error("获取数据库的所有数据库名 error:", err.Error())
	//	response.Fail("获取数据库的所有数据库名 error:"+err.Error(), nil, ctx)
	//	return
	//}
	res, err := service.GetDB()
	if err != nil {
		global.Logrus.Error("获取数据库的所有数据库名 error:", err.Error())
		response.Fail("获取数据库的所有数据库名 error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("获取数据库的所有数据库名成功", res, ctx)
}

// 获取数据库的所有表名
func GetTables(ctx *gin.Context) {
	var dbName model.DbNameAndTableReq
	err := ctx.ShouldBind(&dbName)
	if err != nil {
		global.Logrus.Error("获取数据库的所有表名 error:", err.Error())
		response.Fail("获取数据库的所有表名 error:"+err.Error(), nil, ctx)
		return
	}
	if dbName.Database == "" {
		response.Fail("数据库名为空", nil, ctx)
		return
	}
	res, err := service.GetTables(dbName.Database)
	if err != nil {
		global.Logrus.Error("获取数据库的所有表名 error:", err.Error())
		response.Fail("获取数据库的所有表名 error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("获取数据库的所有表名成功", res, ctx)
}

// 获取字段名,类型值
func GetColumn(ctx *gin.Context) {
	//body, err := ioutil.ReadAll(ctx.Request.Body)
	//fmt.Println("body:", string(body), err)
	var dbNameAndTable model.DbNameAndTableReq
	err := ctx.ShouldBind(&dbNameAndTable)
	if err != nil {
		global.Logrus.Error("获取数据表所有字段名, error:", err.Error())
		response.Fail("获取数据表所有字段名, error:"+err.Error(), nil, ctx)
		return
	}
	if dbNameAndTable.Database == "" {
		if global.Config.SystemParams.DbType == "mysql" {
			dbNameAndTable.Database = global.Config.Mysql.Dbname
		} else {
			dbNameAndTable.Database = global.Config.Sqlite.Path
		}
	}
	//fmt.Println("所有字段名:", dbNameAndTable)
	res, err := service.GetColumnByDB(dbNameAndTable.Database, dbNameAndTable.TableName)
	if err != nil {
		global.Logrus.Error("获取数据表所有字段名, error:", err.Error())
		response.Fail("获取数据表所有字段名, error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("获取数据库所有字段名成功", res, ctx)
}

// 获取字段名,类型值
func GetColumnNew(ctx *gin.Context) {
	var dbNameAndTable model.DbNameAndTableReq
	err := ctx.ShouldBind(&dbNameAndTable)
	if err != nil {
		global.Logrus.Error("获取数据表所有字段名, error:", err.Error())
		response.Fail("获取数据表所有字段名, error:"+err.Error(), nil, ctx)
		return
	}
	m1, m2, m3 := service.GetColumnByReflect(dbNameAndTable.TableName)

	response.OK("获取库表字段信息成功", gin.H{
		"field_list":              m1,
		"field_chinese_name_list": m2,
		"field_type_list":         m3,
	}, ctx)
}

// 获取报表
func ReportSubmit(ctx *gin.Context) {
	//body, err := ioutil.ReadAll(ctx.Request.Body)
	//fmt.Println("body:", string(body), err)

	var fieldParams model.FieldParamsReq
	err := ctx.ShouldBind(&fieldParams)
	if err != nil {
		global.Logrus.Error("获取报表, error:", err.Error())
		response.Fail("获取报表,error:"+err.Error(), nil, ctx)
		return
	}
	//fmt.Println("获取报表:", fieldParams)
	res, total, err := service.GetReport(fieldParams)
	if err != nil {
		global.Logrus.Error("获取报表,error:", err.Error())
		response.Fail("获取报表, error:"+err.Error(), nil, ctx)
		return
	}
	var tableHeader interface{}
	switch fieldParams.TableName {
	case "user":
		tableHeader = model.UserHeaderItem
	case "orders":
		tableHeader = model.OrdersHeaderItem
	case "gallery":
		tableHeader = model.GalleryHeaderItem
	default:
	}
	response.OK("获取报表成功", gin.H{
		"table_header": tableHeader,
		"table_data":   res,
		"total":        total,
	}, ctx)

}
