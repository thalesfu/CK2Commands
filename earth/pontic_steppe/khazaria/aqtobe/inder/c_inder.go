package inder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type InderCounty interface {
    feud.County
    BBogunbay博贡拜() 	feud.Barony
    BChausay乔_赛() 	feud.Barony
    BDzhanatalap贾纳塔拉普() 	feud.Barony
    BEbita埃比德() 	feud.Barony
    BKaratogay卡拉托盖() 	feud.Barony
    BKaumetey考梅捷伊() 	feud.Barony
    BKemer凯梅尔() 	feud.Barony
    BOrsk因德博尔() 	feud.Barony
}

type 伊列克InderCounty struct {
	feud.BaseCounty
	博贡拜Bogunbay 	feud.Barony
	乔_赛Chausay 	feud.Barony
	贾纳塔拉普Dzhanatalap 	feud.Barony
	埃比德Ebita 	feud.Barony
	卡拉托盖Karatogay 	feud.Barony
	考梅捷伊Kaumetey 	feud.Barony
	凯梅尔Kemer 	feud.Barony
	因德博尔Orsk 	feud.Barony
}

func (c *伊列克InderCounty) BBogunbay博贡拜() feud.Barony {
	return c.博贡拜Bogunbay
}
    
func (c *伊列克InderCounty) BChausay乔_赛() feud.Barony {
	return c.乔_赛Chausay
}
    
func (c *伊列克InderCounty) BDzhanatalap贾纳塔拉普() feud.Barony {
	return c.贾纳塔拉普Dzhanatalap
}
    
func (c *伊列克InderCounty) BEbita埃比德() feud.Barony {
	return c.埃比德Ebita
}
    
func (c *伊列克InderCounty) BKaratogay卡拉托盖() feud.Barony {
	return c.卡拉托盖Karatogay
}
    
func (c *伊列克InderCounty) BKaumetey考梅捷伊() feud.Barony {
	return c.考梅捷伊Kaumetey
}
    
func (c *伊列克InderCounty) BKemer凯梅尔() feud.Barony {
	return c.凯梅尔Kemer
}
    
func (c *伊列克InderCounty) BOrsk因德博尔() feud.Barony {
	return c.因德博尔Orsk
}
    
var CInder伊列克 InderCounty = &伊列克InderCounty{}

func init() {
	f := CInder伊列克.(*伊列克InderCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "897",
		Title:     "inder",
		TitleName: "伊列克",
		TitleCode: "c_inder",
		Baronies:  map[string]feud.Barony{},
	}

	f.博贡拜Bogunbay = BBogunbay博贡拜
	f.博贡拜Bogunbay.SetParent(f)
	
	f.乔_赛Chausay = BChausay乔_赛
	f.乔_赛Chausay.SetParent(f)
	
	f.贾纳塔拉普Dzhanatalap = BDzhanatalap贾纳塔拉普
	f.贾纳塔拉普Dzhanatalap.SetParent(f)
	
	f.埃比德Ebita = BEbita埃比德
	f.埃比德Ebita.SetParent(f)
	
	f.卡拉托盖Karatogay = BKaratogay卡拉托盖
	f.卡拉托盖Karatogay.SetParent(f)
	
	f.考梅捷伊Kaumetey = BKaumetey考梅捷伊
	f.考梅捷伊Kaumetey.SetParent(f)
	
	f.凯梅尔Kemer = BKemer凯梅尔
	f.凯梅尔Kemer.SetParent(f)
	
	f.因德博尔Orsk = BOrsk因德博尔
	f.因德博尔Orsk.SetParent(f)
	
}
