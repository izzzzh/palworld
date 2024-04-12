package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Pal struct {
	ID           int64  `json:"id"`
	Number       string `json:"number"`
	Name         string `json:"name"`
	EnName       string `json:"en_name"`
	Icon         string `json:"icon"`
	AttributeIds string `json:"attribute_ids"`
	Abilities    string `json:"abilities"`
	HP           int    `json:"hp"`
	Energy       int    `json:"energy"`
	Defensively  int    `json:"defensively"`
	Eat          int    `json:"eat"`
	Passive      string `json:"passive"`
	PassiveDesc  string `json:"passive_desc"`
	Description  string `json:"description"`
}

type Skill struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AttributeId int    `json:"attribute_id"`
	CT          int    `json:"ct"`
	Power       int    `json:"power"`
}

type Goods struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Types       string `json:"types"`
	Quality     int    `json:"quality"`
}

type PalMateMap struct {
	ID        int64 `json:"id"`
	ParentOne int64 `json:"parent_one"`
	ParentTwo int64 `json:"parent_two"`
	Result    int64 `json:"result"`
}

var mapPal = make(map[string]Pal)

var mapAttribute = map[string]int{
	"一般": 1,
	"暗":  2,
	"龙":  3,
	"冰":  4,
	"火":  5,
	"木":  6,
	"地":  7,
	"电":  8,
	"水":  9,
}

type Abillity struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	Icon  int    `json:"icon"`
}

var mapAbility = map[int]string{
	7:  "生火",
	8:  "浇水",
	9:  "种植",
	10: "发电",
	11: "手工",
	12: "采集",
	13: "伐木",
	14: "挖矿",
	15: "制药",
	16: "冷却",
	17: "牧场",
	18: "搬运",
}

type PalSkillMap struct {
	ID      int64 `json:"id"`
	PalID   int64 `json:"pal_id"`
	SkillID int64 `json:"skill_id"`
}

type TechnologyTree struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Cost        int    `json:"cost"`
	Level       int    `json:"level"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

type TechMaterial struct {
	TechnologyID int64 `json:"technology_id"`
	MaterialID   int64 `json:"material_id"`
	Count        int   `json:"count"`
}

type GoodsMaterial struct {
	GoodsID    int64 `json:"goods_id"`
	MaterialID int64 `json:"material_id"`
	Count      int   `json:"count"`
}

func main() {
	updateTech()
}

func updateTech() {
	dsn := "root:palworld@admin123@tcp(120.78.196.38:3306)/palworld?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var techs []TechnologyTree
	db.Table("technology_tree").Find(&techs)
	for i := range techs {
		tech := techs[i]
		resp, err := http.Get("https://wiki.biligame.com/palworld/" + tech.Name)
		if err != nil {
			fmt.Println(err)
			return
		}
		dom, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		var desc string
		dom.Find(".palworld-textbox").First().Each(func(i int, selection *goquery.Selection) {
			desc = selection.Text()
			desc = strings.TrimSpace(desc)
			desc = strings.Replace(desc, "\n", "", -1)
			desc = strings.Replace(desc, " ", "", -1)
		})
		var materials []TechMaterial
		dom.Find(".palworld-textbox").Last().Each(func(i int, selection *goquery.Selection) {
			selection.Find("div").Each(func(j int, s *goquery.Selection) {
				k := strings.TrimSpace(s.Text())
				k = strings.Replace(k, "\n", "", -1)
				k = strings.Replace(k, " ", "", -1)
				k = strings.TrimSpace(k)
				if k != "" {
					t := strings.Split(k, "x")
					count, _ := strconv.Atoi(t[1])
					var g Goods
					db.Table("goods").Where("name = ?", t[0]).First(&g)
					materials = append(materials, TechMaterial{
						TechnologyID: tech.ID,
						MaterialID:   int64(g.Id),
						Count:        count,
					})
				}
			})
		})
		tech.Description = desc
		fmt.Println(materials)
		db.Table("technology_material").Create(&materials)
		//db.Table("technology_tree").Updates(&tech)
	}
}

func updateGoods() {
	//client := &http.Client{}
	//// 设置请求方法和URL
	//
	//req, err := http.NewRequest(http.MethodGet, "https://wiki.biligame.com/palworld/%E9%81%93%E5%85%B7%E4%B8%80%E8%A7%88", nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 发送请求
	//resp, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//dom, err := goquery.NewDocumentFromReader(resp.Body)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//dom.Find(".list ").Each(func(i int, s *goquery.Selection) {
	//	s.Find(".items .item").Each(func(j int, selection *goquery.Selection) {
	//		style, exists := selection.Attr("style")
	//		reg := regexp.MustCompile(`\((.*?)\)`)
	//		if exists {
	//			matches := reg.FindAllStringSubmatch(style, -1)
	//			for _, match := range matches {
	//				if len(match) == 2 {
	//					url := "https://palworld.gg" + match[1]
	//					fmt.Println(url)
	//					downloadImage(url, "."+match[1])
	//				}
	//			}
	//		}
	//	})
	//})
	//var goods []string
	//var names []string
	//dom.Find(".divsort td").Each(func(i int, s *goquery.Selection) {
	//	if _, ok := s.Find("img").Attr("src"); ok {
	//		key, _ := s.Find("img").Attr("alt")
	//		imageName := strings.Replace(key, " ", "_", -1)
	//		//downloadImage(url, "./images/goods/"+imageName)
	//		names = append(names, "/images/goods/"+imageName)
	//	}
	//	goods = append(goods, strings.TrimSpace(s.Text()))
	//})

	dsn := "root:123456@tcp(127.0.0.1:3306)/palworld?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var ret []string
	white := map[string]struct{}{}
	white["棉花糖"] = struct{}{}
	var goods []Goods
	mapGoods := make(map[string]int)
	db.Table("goods").Select("id,name").Where("quality = 0").Find(&goods)
	for _, val := range goods {
		mapGoods[val.Name] = val.Id
	}

	db.Table("goods").Group("name").Select("name").Find(&ret)
	for i := range ret {
		if _, ok := white[ret[i]]; ok {
			continue
		}
		url := "https://wiki.biligame.com/palworld/" + ret[i]
		materials := getGoodsMaterial(ret[i], url)

		if len(materials) > 0 {
			m := make(map[int]int)
			for _, val := range materials {
				arr := strings.Split(val, "x")
				v, _ := strconv.Atoi(arr[1])
				m[mapGoods[arr[0]]] = v

				gm := &GoodsMaterial{
					GoodsID:    int64(mapGoods[ret[i]]),
					MaterialID: int64(mapGoods[arr[0]]),
					Count:      v,
				}
				fmt.Println(gm)
				db.Table("goods_material").Create(&gm)
			}
		}
	}
}

func getGoodsMaterial(name, url string) []string {
	var str []string
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return str
	}
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	if dom.Find(".palworld-textbox").Length() > 1 {
		//fmt.Println(name)
		dom.Find(".palworld-textbox").Last().Each(func(i int, selection *goquery.Selection) {

			selection.Find("div").Each(func(j int, s *goquery.Selection) {
				k := strings.TrimSpace(s.Text())
				k = strings.Replace(k, "\n", "", -1)
				k = strings.Replace(k, " ", "", -1)
				k = strings.TrimSpace(k)
				if k != "" {
					str = append(str, k)
				}
			})

		})
	}
	return str
}

func insertPalSkillMap() {
	dsn := "root:palworld@admin123@tcp(120.78.196.38:3306)/palworld?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	pals := make([]*Pal, 0)
	db.Table("pal").Find(&pals)

	for i := range pals {
		pal := pals[i]
		newStr := strings.Replace(pal.Icon, "http://120.78.196.38", "https://ppcat.fun", 1)
		pal.Icon = newStr
		fmt.Println(pal.Icon)
		db.Table("pal").Save(&pal)
	}

	//goods := make([]*Goods, 0)
	//db.Table("goods").Find(&pals)
	//
	//for i := range goods {
	//	good := goods[i]
	//	newStr := strings.Replace(good.Image, "http://120.78.196.38", "https://ppcat.fun", 1)
	//	good.Image = newStr
	//	fmt.Println(good.Image)
	//	db.Table("goods").Save(&good)
	//}

	//skills := make([]*Skill, 0)
	//db.Table("skill").Find(&skills)
	//
	//skillMap := make(map[string]int64)
	//for _, val := range skills {
	//	skillMap[val.Name] = val.ID
	//}
	//
	//for i := range pals {
	//	pal := pals[i]
	//	resp, err := http.Get("https://wiki.biligame.com/palworld/" + pal.Name)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//
	//	dom, err := goquery.NewDocumentFromReader(resp.Body)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	regexp.MustCompile(`^\n`)
	//	dom.Find(".palworld-textbox").Each(func(i int, s *goquery.Selection) {
	//		if i > 1 {
	//			t := s.Text()
	//			reg := regexp.MustCompile(`\n`)
	//			// 替换所有空行为空字符串
	//			output := reg.ReplaceAllString(t, " ")
	//			strs := strings.Split(output, " ")
	//			sName := strs[1][3:]
	//			id := skillMap[sName]
	//			psm := &PalSkillMap{
	//				PalID:   pal.ID,
	//				SkillID: id,
	//			}
	//			db.Table("pal_skill_map").Create(&psm)
	//		}
	//	})
	//}
}

func updatePal() {
	f, err := excelize.OpenFile("scripts/data/palworld.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := f.GetRows("帕鲁工作技能表")

	dsn := "root:123456@tcp(127.0.0.1:3306)/palworld?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	pals := make([]*Pal, 0)
	db.Table("pal").Find(&pals)
	palMap := make(map[string]*Pal)
	for i, val := range pals {
		palMap[val.Name] = pals[i]
	}
	for i, row := range rows {
		if i > 1 {
			abilities := make([]Abillity, 0)
			for j := 7; j < 19; j++ {
				val, _ := strconv.Atoi(row[j])
				if val > 0 {
					abilities = append(abilities, Abillity{
						Name:  mapAbility[j],
						Level: val,
						Icon:  j - 6,
					})
				}
			}
			jsonA, _ := json.Marshal(abilities)
			eat, _ := strconv.Atoi(row[4])
			p := palMap[row[2]]
			p.Eat = eat
			p.Abilities = string(jsonA)
			db.Table("pal").Updates(&p)
		}
	}
}

func insertMap() {
	f, err := excelize.OpenFile("scripts/data/palworld.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := f.GetRows("配种组合")

	dsn := "root:123456@tcp(127.0.0.1:3306)/palworld?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	pals := make([]*Pal, 0)
	db.Table("pal").Find(&pals)
	palMap := make(map[string]int64)
	for _, val := range pals {
		palMap[val.Name] = val.ID
	}

	for i, row := range rows {
		if i > 1 && row[6] == "1" {
		}
	}
}

func insertGoods() {
	f, err := excelize.OpenFile("scripts/data/palworld.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	localPath := "./images/goods/"
	rows := f.GetRows("物品配方")

	//dsn := "root:123456@tcp(127.0.0.1:3306)/palworld?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	prefix := []string{"Material_"}
	for i, row := range rows {
		if i > 0 {
			//var goods []*Pal
			//_ = db.Table("pal").Find()
			//goods.Image = "http://120.78.196.38/palworld/images/goods/" + goods.EnName + ".png"
			//
			//db.Table("goods").Updates(goods)
			for _, val := range prefix {
				imageUrl := "https://data-resource.caimogu.cc/palworld/style/img/inventory/T_itemicon_" + val + row[1] + ".png"
				downloadImage(imageUrl, localPath+row[1])
			}
		}
	}
}

func checkSkill(name, desc string) bool {
	str := strings.Replace(desc, "\n", "", -1)
	i := strings.Index(str, "zh-Hans") < 0
	j := strings.Index(name, "zh-Hans") < 0
	k := strings.Index(name, "#") < 0 && strings.Index(str, "#") < 0
	return i && j && k
}

func changeDesc(desc string) string {
	var ret = desc
	if strings.Index(desc, "characterName") > 0 {
		re2 := regexp.MustCompile(`\|(.*?)\|`)
		match := re2.FindStringSubmatch(desc)
		if match != nil && len(match) > 1 {
			re := regexp.MustCompile(`\<(.*?)\>`)
			newStr := re.ReplaceAllString(desc, mapPal[match[1]].Name)
			ret = newStr
		}
	}
	return ret
}

func insertSkill() {
	f, err := excelize.OpenFile("scripts/data/palworld.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/palworld?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows := f.GetRows("主动技能表")
	for i, row := range rows {
		str := strings.Replace(row[2], "\n", "", -1)
		if i > 0 && checkSkill(row[1], row[2]) {
			ct, _ := strconv.Atoi(row[7])
			power, _ := strconv.Atoi(row[6])
			//fmt.Println(row[0], row[1], row[3])
			skill := Skill{
				Name:        row[1],
				Description: changeDesc(str),
				CT:          ct,
				Power:       power,
				AttributeId: mapAttribute[row[4]],
			}
			fmt.Println(skill)
			db.Table("skill").Create(&skill)
		}
	}
}

func insertPal() {
	f, err := excelize.OpenFile("scripts/data/palworld.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows := f.GetRows("参数")
	for _, row := range rows {
		if row[0] != "" {
			//fmt.Println(row[0], row[1], row[3])
			pal := Pal{
				Number: row[0],
				Name:   row[1],
				EnName: row[3],
			}
			mapPal[row[3]] = pal
		}
	}
	//dsn := "root:123456@tcp(127.0.0.1:3306)/palworld?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//palRows := f.GetRows("帕鲁战斗属性表")
	//for _, row := range palRows {
	//	pal, ok := mapPal[row[1]]
	//	if ok {
	//		attribute := strconv.Itoa(mapAttribute[row[13]])
	//		if row[14] != "" {
	//			attribute += "," + strconv.Itoa(mapAttribute[row[14]])
	//		}
	//		pal.HP, _ = strconv.Atoi(row[21])
	//		pal.Energy, _ = strconv.Atoi(row[22])
	//		pal.Defensively, _ = strconv.Atoi(row[24])
	//		pal.AttributeIds = attribute
	//		pal.Icon = "http://120.78.196.38/palworld/images/" + row[3] + ".png"
	//		mapPal[row[1]] = pal
	//		result := db.Table("pal").Where("name = ?", row[1]).Updates(&pal)
	//		if result.Error != nil {
	//			fmt.Println(err)
	//		}
	//	}
	//}
}

func getImage() {
	f, err := excelize.OpenFile("scripts/data/palworld.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	localPath := "../images/"
	os.MkdirAll(localPath, os.ModePerm)
	// Get all the rows in the Sheet1.
	rows := f.GetRows("参数")
	for _, row := range rows {
		if row[0] != "" {
			imageUrl := "https://data-resource.caimogu.cc/palworld/style/img/pal/T_" + row[3] + "_icon_normal.png"
			imagePath := localPath + row[3]
			fmt.Println(row[0], row[1], row[3], imageUrl)
			downloadImage(imageUrl, imagePath)
		}
	}

}

func getIcon() {
	localPath := "./images/icon/"
	for i := 0; i <= 8; i++ {
		str := strconv.Itoa(i)
		imageUrl := "https://data-resource.caimogu.cc/palworld/style/img/icon/T_Icon_element_s_0" + str + ".png"
		downloadImage(imageUrl, localPath+str)
	}
}

func downloadImage(imageUrl, localPath string) {
	resp, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("http get err:", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		return
	}
	fmt.Println(imageUrl)
	defer resp.Body.Close()

	//打开文件流
	picName := localPath
	f, err := os.Create(picName)
	if err != nil {
		fmt.Println("os create err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	//读httpbody数据写入文件流
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}

		f.Write(buf[:n])
	}
}
