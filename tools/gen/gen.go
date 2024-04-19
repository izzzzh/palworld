package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var gormdb *gorm.DB

func InitDb() {
	gormdb, _ = gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/palworld?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
}

func genUserServer() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./rpc/user/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式,
	})

	g.UseDB(gormdb) // reuse your gorm db
	g.ApplyBasic(
		// 从当前数据库中生成所有表的结构
		g.GenerateModel("user"),
	)

	// 生成代码
	g.Execute()
}

func genSkillServer() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./rpc/skill/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式,
	})

	g.UseDB(gormdb) // reuse your gorm db
	g.ApplyBasic(
		// 从当前数据库中生成所有表的结构
		g.GenerateModel("skill"),
		g.GenerateModel("pal_skill_map"),
	)

	// 生成代码
	g.Execute()
}

func genPalServer() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./rpc/pal/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式,
	})

	g.UseDB(gormdb) // reuse your gorm db
	g.ApplyBasic(
		// 从当前数据库中生成所有表的结构
		g.GenerateModel("pal"),
		g.GenerateModel("pal_skill_map"),
		g.GenerateModel("skill"),
	)

	// 生成代码
	g.Execute()
}

func genGoodsServer() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./rpc/goods/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式,
	})

	g.UseDB(gormdb) // reuse your gorm db
	g.ApplyBasic(
		// 从当前数据库中生成所有表的结构
		g.GenerateModel("goods"),
		g.GenerateModel("goods_material"),
	)

	// 生成代码
	g.Execute()
}

func genTechServer() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./rpc/technology_tree/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式,
	})

	g.UseDB(gormdb) // reuse your gorm db
	g.ApplyBasic(
		// 从当前数据库中生成所有表的结构
		g.GenerateModel("technology_tree"),
		g.GenerateModel("technology_material"),
	)

	// 生成代码
	g.Execute()
}

func genCommentsServer() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./rpc/comments/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式,
	})

	g.UseDB(gormdb) // reuse your gorm db
	g.ApplyBasic(
		// 从当前数据库中生成所有表的结构
		g.GenerateModel("comments"),
	)

	// 生成代码
	g.Execute()
}

func genPalMateServer() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./rpc/pal_mate/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式,
	})

	g.UseDB(gormdb) // reuse your gorm db
	g.ApplyBasic(
		// 从当前数据库中生成所有表的结构
		g.GenerateModel("pal_mate_map"),
	)

	// 生成代码
	g.Execute()
}
func main() {
	InitDb()
	genUserServer()
	genCommentsServer()
	genSkillServer()
	genPalServer()
	genTechServer()
	genGoodsServer()
	genPalMateServer()
}
