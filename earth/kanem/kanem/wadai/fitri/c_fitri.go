package fitri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FitriCounty interface {
    feud.County
    BAtourda阿图尔达() 	feud.Barony
    BGaska加斯卡() 	feud.Barony
    BKuka库卡() 	feud.Barony
    BManouachi马努瓦希() 	feud.Barony
    BNdoye恩多耶() 	feud.Barony
    BSambourou桑布鲁() 	feud.Barony
    BTroan特罗昂() 	feud.Barony
}

type 菲特里FitriCounty struct {
	feud.BaseCounty
	阿图尔达Atourda 	feud.Barony
	加斯卡Gaska 	feud.Barony
	库卡Kuka 	feud.Barony
	马努瓦希Manouachi 	feud.Barony
	恩多耶Ndoye 	feud.Barony
	桑布鲁Sambourou 	feud.Barony
	特罗昂Troan 	feud.Barony
}

func (c *菲特里FitriCounty) BAtourda阿图尔达() feud.Barony {
	return c.阿图尔达Atourda
}
    
func (c *菲特里FitriCounty) BGaska加斯卡() feud.Barony {
	return c.加斯卡Gaska
}
    
func (c *菲特里FitriCounty) BKuka库卡() feud.Barony {
	return c.库卡Kuka
}
    
func (c *菲特里FitriCounty) BManouachi马努瓦希() feud.Barony {
	return c.马努瓦希Manouachi
}
    
func (c *菲特里FitriCounty) BNdoye恩多耶() feud.Barony {
	return c.恩多耶Ndoye
}
    
func (c *菲特里FitriCounty) BSambourou桑布鲁() feud.Barony {
	return c.桑布鲁Sambourou
}
    
func (c *菲特里FitriCounty) BTroan特罗昂() feud.Barony {
	return c.特罗昂Troan
}
    
var CFitri菲特里 FitriCounty = &菲特里FitriCounty{}

func init() {
	f := CFitri菲特里.(*菲特里FitriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1769",
		Title:     "fitri",
		TitleName: "菲特里",
		TitleCode: "c_fitri",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿图尔达Atourda = BAtourda阿图尔达
	f.阿图尔达Atourda.SetParent(f)
	
	f.加斯卡Gaska = BGaska加斯卡
	f.加斯卡Gaska.SetParent(f)
	
	f.库卡Kuka = BKuka库卡
	f.库卡Kuka.SetParent(f)
	
	f.马努瓦希Manouachi = BManouachi马努瓦希
	f.马努瓦希Manouachi.SetParent(f)
	
	f.恩多耶Ndoye = BNdoye恩多耶
	f.恩多耶Ndoye.SetParent(f)
	
	f.桑布鲁Sambourou = BSambourou桑布鲁
	f.桑布鲁Sambourou.SetParent(f)
	
	f.特罗昂Troan = BTroan特罗昂
	f.特罗昂Troan.SetParent(f)
	
}
