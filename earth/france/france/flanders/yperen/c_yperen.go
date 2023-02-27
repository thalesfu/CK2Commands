package yperen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YperenCounty interface {
    feud.County
    BCassel卡塞勒() 	feud.Barony
    BDiksmuide迪克斯迈德() 	feud.Barony
    BMenen梅嫩() 	feud.Barony
    BPoperinge波珀灵厄() 	feud.Barony
    BRoeselare鲁瑟拉勒() 	feud.Barony
    BRosebeke罗泽贝克() 	feud.Barony
    BWervik韦尔菲克() 	feud.Barony
    BYpres伊普尔() 	feud.Barony
}

type 伊普尔YperenCounty struct {
	feud.BaseCounty
	卡塞勒Cassel 	feud.Barony
	迪克斯迈德Diksmuide 	feud.Barony
	梅嫩Menen 	feud.Barony
	波珀灵厄Poperinge 	feud.Barony
	鲁瑟拉勒Roeselare 	feud.Barony
	罗泽贝克Rosebeke 	feud.Barony
	韦尔菲克Wervik 	feud.Barony
	伊普尔Ypres 	feud.Barony
}

func (c *伊普尔YperenCounty) BCassel卡塞勒() feud.Barony {
	return c.卡塞勒Cassel
}
    
func (c *伊普尔YperenCounty) BDiksmuide迪克斯迈德() feud.Barony {
	return c.迪克斯迈德Diksmuide
}
    
func (c *伊普尔YperenCounty) BMenen梅嫩() feud.Barony {
	return c.梅嫩Menen
}
    
func (c *伊普尔YperenCounty) BPoperinge波珀灵厄() feud.Barony {
	return c.波珀灵厄Poperinge
}
    
func (c *伊普尔YperenCounty) BRoeselare鲁瑟拉勒() feud.Barony {
	return c.鲁瑟拉勒Roeselare
}
    
func (c *伊普尔YperenCounty) BRosebeke罗泽贝克() feud.Barony {
	return c.罗泽贝克Rosebeke
}
    
func (c *伊普尔YperenCounty) BWervik韦尔菲克() feud.Barony {
	return c.韦尔菲克Wervik
}
    
func (c *伊普尔YperenCounty) BYpres伊普尔() feud.Barony {
	return c.伊普尔Ypres
}
    
var CYperen伊普尔 YperenCounty = &伊普尔YperenCounty{}

func init() {
	f := CYperen伊普尔.(*伊普尔YperenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "76",
		Title:     "yperen",
		TitleName: "伊普尔",
		TitleCode: "c_yperen",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡塞勒Cassel = BCassel卡塞勒
	f.卡塞勒Cassel.SetParent(f)
	
	f.迪克斯迈德Diksmuide = BDiksmuide迪克斯迈德
	f.迪克斯迈德Diksmuide.SetParent(f)
	
	f.梅嫩Menen = BMenen梅嫩
	f.梅嫩Menen.SetParent(f)
	
	f.波珀灵厄Poperinge = BPoperinge波珀灵厄
	f.波珀灵厄Poperinge.SetParent(f)
	
	f.鲁瑟拉勒Roeselare = BRoeselare鲁瑟拉勒
	f.鲁瑟拉勒Roeselare.SetParent(f)
	
	f.罗泽贝克Rosebeke = BRosebeke罗泽贝克
	f.罗泽贝克Rosebeke.SetParent(f)
	
	f.韦尔菲克Wervik = BWervik韦尔菲克
	f.韦尔菲克Wervik.SetParent(f)
	
	f.伊普尔Ypres = BYpres伊普尔
	f.伊普尔Ypres.SetParent(f)
	
}
