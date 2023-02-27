package kharkhiraa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KharkhiraaCounty interface {
    feud.County
    BBaatar巴特尔() 	feud.Barony
    BBumbag奔巴图() 	feud.Barony
    BDarvi_kharkhiraa达尔维() 	feud.Barony
    BKharkhiraa哈尔齐拉() 	feud.Barony
    BMyangan_ugalzat明安乌嘎勒扎特() 	feud.Barony
    BSutay苏泰() 	feud.Barony
    BTsetseg车车格() 	feud.Barony
}

type 哈尔齐拉KharkhiraaCounty struct {
	feud.BaseCounty
	巴特尔Baatar 	feud.Barony
	奔巴图Bumbag 	feud.Barony
	达尔维Darvi_kharkhiraa 	feud.Barony
	哈尔齐拉Kharkhiraa 	feud.Barony
	明安乌嘎勒扎特Myangan_ugalzat 	feud.Barony
	苏泰Sutay 	feud.Barony
	车车格Tsetseg 	feud.Barony
}

func (c *哈尔齐拉KharkhiraaCounty) BBaatar巴特尔() feud.Barony {
	return c.巴特尔Baatar
}
    
func (c *哈尔齐拉KharkhiraaCounty) BBumbag奔巴图() feud.Barony {
	return c.奔巴图Bumbag
}
    
func (c *哈尔齐拉KharkhiraaCounty) BDarvi_kharkhiraa达尔维() feud.Barony {
	return c.达尔维Darvi_kharkhiraa
}
    
func (c *哈尔齐拉KharkhiraaCounty) BKharkhiraa哈尔齐拉() feud.Barony {
	return c.哈尔齐拉Kharkhiraa
}
    
func (c *哈尔齐拉KharkhiraaCounty) BMyangan_ugalzat明安乌嘎勒扎特() feud.Barony {
	return c.明安乌嘎勒扎特Myangan_ugalzat
}
    
func (c *哈尔齐拉KharkhiraaCounty) BSutay苏泰() feud.Barony {
	return c.苏泰Sutay
}
    
func (c *哈尔齐拉KharkhiraaCounty) BTsetseg车车格() feud.Barony {
	return c.车车格Tsetseg
}
    
var CKharkhiraa哈尔齐拉 KharkhiraaCounty = &哈尔齐拉KharkhiraaCounty{}

func init() {
	f := CKharkhiraa哈尔齐拉.(*哈尔齐拉KharkhiraaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1907",
		Title:     "kharkhiraa",
		TitleName: "哈尔齐拉",
		TitleCode: "c_kharkhiraa",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴特尔Baatar = BBaatar巴特尔
	f.巴特尔Baatar.SetParent(f)
	
	f.奔巴图Bumbag = BBumbag奔巴图
	f.奔巴图Bumbag.SetParent(f)
	
	f.达尔维Darvi_kharkhiraa = BDarvi_kharkhiraa达尔维
	f.达尔维Darvi_kharkhiraa.SetParent(f)
	
	f.哈尔齐拉Kharkhiraa = BKharkhiraa哈尔齐拉
	f.哈尔齐拉Kharkhiraa.SetParent(f)
	
	f.明安乌嘎勒扎特Myangan_ugalzat = BMyangan_ugalzat明安乌嘎勒扎特
	f.明安乌嘎勒扎特Myangan_ugalzat.SetParent(f)
	
	f.苏泰Sutay = BSutay苏泰
	f.苏泰Sutay.SetParent(f)
	
	f.车车格Tsetseg = BTsetseg车车格
	f.车车格Tsetseg.SetParent(f)
	
}
