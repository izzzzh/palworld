package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./rpc/goods/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式
	})

	gormdb, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/palworld?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	g.ApplyBasic(
		// 从当前数据库中生成所有表的结构
		g.GenerateModel("goods"),
	)

	// 生成代码
	g.Execute()
}
