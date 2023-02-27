package tsilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TsilmaCounty interface {
    feud.County
    BByk贝克() 	feud.Barony
    BMutnaya穆特纳亚() 	feud.Barony
    BMyla梅拉() 	feud.Barony
    BNonbur农布尔() 	feud.Barony
    BRudyanka鲁江卡() 	feud.Barony
    BTsilma齐利马() 	feud.Barony
    BUyeg乌耶格() 	feud.Barony
}

type 齐利马TsilmaCounty struct {
	feud.BaseCounty
	贝克Byk 	feud.Barony
	穆特纳亚Mutnaya 	feud.Barony
	梅拉Myla 	feud.Barony
	农布尔Nonbur 	feud.Barony
	鲁江卡Rudyanka 	feud.Barony
	齐利马Tsilma 	feud.Barony
	乌耶格Uyeg 	feud.Barony
}

func (c *齐利马TsilmaCounty) BByk贝克() feud.Barony {
	return c.贝克Byk
}
    
func (c *齐利马TsilmaCounty) BMutnaya穆特纳亚() feud.Barony {
	return c.穆特纳亚Mutnaya
}
    
func (c *齐利马TsilmaCounty) BMyla梅拉() feud.Barony {
	return c.梅拉Myla
}
    
func (c *齐利马TsilmaCounty) BNonbur农布尔() feud.Barony {
	return c.农布尔Nonbur
}
    
func (c *齐利马TsilmaCounty) BRudyanka鲁江卡() feud.Barony {
	return c.鲁江卡Rudyanka
}
    
func (c *齐利马TsilmaCounty) BTsilma齐利马() feud.Barony {
	return c.齐利马Tsilma
}
    
func (c *齐利马TsilmaCounty) BUyeg乌耶格() feud.Barony {
	return c.乌耶格Uyeg
}
    
var CTsilma齐利马 TsilmaCounty = &齐利马TsilmaCounty{}

func init() {
	f := CTsilma齐利马.(*齐利马TsilmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1833",
		Title:     "tsilma",
		TitleName: "齐利马",
		TitleCode: "c_tsilma",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝克Byk = BByk贝克
	f.贝克Byk.SetParent(f)
	
	f.穆特纳亚Mutnaya = BMutnaya穆特纳亚
	f.穆特纳亚Mutnaya.SetParent(f)
	
	f.梅拉Myla = BMyla梅拉
	f.梅拉Myla.SetParent(f)
	
	f.农布尔Nonbur = BNonbur农布尔
	f.农布尔Nonbur.SetParent(f)
	
	f.鲁江卡Rudyanka = BRudyanka鲁江卡
	f.鲁江卡Rudyanka.SetParent(f)
	
	f.齐利马Tsilma = BTsilma齐利马
	f.齐利马Tsilma.SetParent(f)
	
	f.乌耶格Uyeg = BUyeg乌耶格
	f.乌耶格Uyeg.SetParent(f)
	
}
