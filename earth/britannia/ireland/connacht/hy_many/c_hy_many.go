package hy_many

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Hy_manyCounty interface {
    feud.County
    BCruachu克鲁胡() 	feud.Barony
    BDelbna德尔布纳() 	feud.Barony
    BOweynagat乌埃夫纳加特() 	feud.Barony
    BRathbeg拉比亚格() 	feud.Barony
    BRathmore拉沃尔() 	feud.Barony
    BRathnadarve拉纳达尔夫() 	feud.Barony
    BRelignaree雷利格纳里() 	feud.Barony
}

type 伊_马尼Hy_manyCounty struct {
	feud.BaseCounty
	克鲁胡Cruachu 	feud.Barony
	德尔布纳Delbna 	feud.Barony
	乌埃夫纳加特Oweynagat 	feud.Barony
	拉比亚格Rathbeg 	feud.Barony
	拉沃尔Rathmore 	feud.Barony
	拉纳达尔夫Rathnadarve 	feud.Barony
	雷利格纳里Relignaree 	feud.Barony
}

func (c *伊_马尼Hy_manyCounty) BCruachu克鲁胡() feud.Barony {
	return c.克鲁胡Cruachu
}
    
func (c *伊_马尼Hy_manyCounty) BDelbna德尔布纳() feud.Barony {
	return c.德尔布纳Delbna
}
    
func (c *伊_马尼Hy_manyCounty) BOweynagat乌埃夫纳加特() feud.Barony {
	return c.乌埃夫纳加特Oweynagat
}
    
func (c *伊_马尼Hy_manyCounty) BRathbeg拉比亚格() feud.Barony {
	return c.拉比亚格Rathbeg
}
    
func (c *伊_马尼Hy_manyCounty) BRathmore拉沃尔() feud.Barony {
	return c.拉沃尔Rathmore
}
    
func (c *伊_马尼Hy_manyCounty) BRathnadarve拉纳达尔夫() feud.Barony {
	return c.拉纳达尔夫Rathnadarve
}
    
func (c *伊_马尼Hy_manyCounty) BRelignaree雷利格纳里() feud.Barony {
	return c.雷利格纳里Relignaree
}
    
var CHy_many伊_马尼 Hy_manyCounty = &伊_马尼Hy_manyCounty{}

func init() {
	f := CHy_many伊_马尼.(*伊_马尼Hy_manyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1955",
		Title:     "hy_many",
		TitleName: "伊_马尼",
		TitleCode: "c_hy_many",
		Baronies:  map[string]feud.Barony{},
	}

	f.克鲁胡Cruachu = BCruachu克鲁胡
	f.克鲁胡Cruachu.SetParent(f)
	
	f.德尔布纳Delbna = BDelbna德尔布纳
	f.德尔布纳Delbna.SetParent(f)
	
	f.乌埃夫纳加特Oweynagat = BOweynagat乌埃夫纳加特
	f.乌埃夫纳加特Oweynagat.SetParent(f)
	
	f.拉比亚格Rathbeg = BRathbeg拉比亚格
	f.拉比亚格Rathbeg.SetParent(f)
	
	f.拉沃尔Rathmore = BRathmore拉沃尔
	f.拉沃尔Rathmore.SetParent(f)
	
	f.拉纳达尔夫Rathnadarve = BRathnadarve拉纳达尔夫
	f.拉纳达尔夫Rathnadarve.SetParent(f)
	
	f.雷利格纳里Relignaree = BRelignaree雷利格纳里
	f.雷利格纳里Relignaree.SetParent(f)
	
}
