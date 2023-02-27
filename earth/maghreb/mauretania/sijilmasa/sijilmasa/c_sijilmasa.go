package sijilmasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SijilmasaCounty interface {
    feud.County
    BErtoud伊尔富德() 	feud.Barony
    BHannabou汉纳布() 	feud.Barony
    BMerzane梅尔赞() 	feud.Barony
    BMerzouga梅尔祖卡() 	feud.Barony
    BSijilmasa锡吉勒马萨() 	feud.Barony
    BTadaout塔达乌特() 	feud.Barony
    BTanamouste塔纳穆斯特() 	feud.Barony
    BTisserdmine提瑟尔德明() 	feud.Barony
}

type 锡吉勒马萨SijilmasaCounty struct {
	feud.BaseCounty
	伊尔富德Ertoud 	feud.Barony
	汉纳布Hannabou 	feud.Barony
	梅尔赞Merzane 	feud.Barony
	梅尔祖卡Merzouga 	feud.Barony
	锡吉勒马萨Sijilmasa 	feud.Barony
	塔达乌特Tadaout 	feud.Barony
	塔纳穆斯特Tanamouste 	feud.Barony
	提瑟尔德明Tisserdmine 	feud.Barony
}

func (c *锡吉勒马萨SijilmasaCounty) BErtoud伊尔富德() feud.Barony {
	return c.伊尔富德Ertoud
}
    
func (c *锡吉勒马萨SijilmasaCounty) BHannabou汉纳布() feud.Barony {
	return c.汉纳布Hannabou
}
    
func (c *锡吉勒马萨SijilmasaCounty) BMerzane梅尔赞() feud.Barony {
	return c.梅尔赞Merzane
}
    
func (c *锡吉勒马萨SijilmasaCounty) BMerzouga梅尔祖卡() feud.Barony {
	return c.梅尔祖卡Merzouga
}
    
func (c *锡吉勒马萨SijilmasaCounty) BSijilmasa锡吉勒马萨() feud.Barony {
	return c.锡吉勒马萨Sijilmasa
}
    
func (c *锡吉勒马萨SijilmasaCounty) BTadaout塔达乌特() feud.Barony {
	return c.塔达乌特Tadaout
}
    
func (c *锡吉勒马萨SijilmasaCounty) BTanamouste塔纳穆斯特() feud.Barony {
	return c.塔纳穆斯特Tanamouste
}
    
func (c *锡吉勒马萨SijilmasaCounty) BTisserdmine提瑟尔德明() feud.Barony {
	return c.提瑟尔德明Tisserdmine
}
    
var CSijilmasa锡吉勒马萨 SijilmasaCounty = &锡吉勒马萨SijilmasaCounty{}

func init() {
	f := CSijilmasa锡吉勒马萨.(*锡吉勒马萨SijilmasaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "918",
		Title:     "sijilmasa",
		TitleName: "锡吉勒马萨",
		TitleCode: "c_sijilmasa",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊尔富德Ertoud = BErtoud伊尔富德
	f.伊尔富德Ertoud.SetParent(f)
	
	f.汉纳布Hannabou = BHannabou汉纳布
	f.汉纳布Hannabou.SetParent(f)
	
	f.梅尔赞Merzane = BMerzane梅尔赞
	f.梅尔赞Merzane.SetParent(f)
	
	f.梅尔祖卡Merzouga = BMerzouga梅尔祖卡
	f.梅尔祖卡Merzouga.SetParent(f)
	
	f.锡吉勒马萨Sijilmasa = BSijilmasa锡吉勒马萨
	f.锡吉勒马萨Sijilmasa.SetParent(f)
	
	f.塔达乌特Tadaout = BTadaout塔达乌特
	f.塔达乌特Tadaout.SetParent(f)
	
	f.塔纳穆斯特Tanamouste = BTanamouste塔纳穆斯特
	f.塔纳穆斯特Tanamouste.SetParent(f)
	
	f.提瑟尔德明Tisserdmine = BTisserdmine提瑟尔德明
	f.提瑟尔德明Tisserdmine.SetParent(f)
	
}
