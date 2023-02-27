package gottingen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GottingenCounty interface {
    feud.County
    BCorvey柯维() 	feud.Barony
    BDetmold代特莫尔德() 	feud.Barony
    BEisenach艾森纳赫() 	feud.Barony
    BGoslar戈斯拉尔() 	feud.Barony
    BGottingen哥廷根() 	feud.Barony
    BKassel卡塞尔() 	feud.Barony
    BLippe利珀() 	feud.Barony
    BNortheim诺特海姆() 	feud.Barony
    BPaderborn帕德博恩() 	feud.Barony
}

type 帕德博恩GottingenCounty struct {
	feud.BaseCounty
	柯维Corvey 	feud.Barony
	代特莫尔德Detmold 	feud.Barony
	艾森纳赫Eisenach 	feud.Barony
	戈斯拉尔Goslar 	feud.Barony
	哥廷根Gottingen 	feud.Barony
	卡塞尔Kassel 	feud.Barony
	利珀Lippe 	feud.Barony
	诺特海姆Northeim 	feud.Barony
	帕德博恩Paderborn 	feud.Barony
}

func (c *帕德博恩GottingenCounty) BCorvey柯维() feud.Barony {
	return c.柯维Corvey
}
    
func (c *帕德博恩GottingenCounty) BDetmold代特莫尔德() feud.Barony {
	return c.代特莫尔德Detmold
}
    
func (c *帕德博恩GottingenCounty) BEisenach艾森纳赫() feud.Barony {
	return c.艾森纳赫Eisenach
}
    
func (c *帕德博恩GottingenCounty) BGoslar戈斯拉尔() feud.Barony {
	return c.戈斯拉尔Goslar
}
    
func (c *帕德博恩GottingenCounty) BGottingen哥廷根() feud.Barony {
	return c.哥廷根Gottingen
}
    
func (c *帕德博恩GottingenCounty) BKassel卡塞尔() feud.Barony {
	return c.卡塞尔Kassel
}
    
func (c *帕德博恩GottingenCounty) BLippe利珀() feud.Barony {
	return c.利珀Lippe
}
    
func (c *帕德博恩GottingenCounty) BNortheim诺特海姆() feud.Barony {
	return c.诺特海姆Northeim
}
    
func (c *帕德博恩GottingenCounty) BPaderborn帕德博恩() feud.Barony {
	return c.帕德博恩Paderborn
}
    
var CGottingen帕德博恩 GottingenCounty = &帕德博恩GottingenCounty{}

func init() {
	f := CGottingen帕德博恩.(*帕德博恩GottingenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "120",
		Title:     "gottingen",
		TitleName: "帕德博恩",
		TitleCode: "c_gottingen",
		Baronies:  map[string]feud.Barony{},
	}

	f.柯维Corvey = BCorvey柯维
	f.柯维Corvey.SetParent(f)
	
	f.代特莫尔德Detmold = BDetmold代特莫尔德
	f.代特莫尔德Detmold.SetParent(f)
	
	f.艾森纳赫Eisenach = BEisenach艾森纳赫
	f.艾森纳赫Eisenach.SetParent(f)
	
	f.戈斯拉尔Goslar = BGoslar戈斯拉尔
	f.戈斯拉尔Goslar.SetParent(f)
	
	f.哥廷根Gottingen = BGottingen哥廷根
	f.哥廷根Gottingen.SetParent(f)
	
	f.卡塞尔Kassel = BKassel卡塞尔
	f.卡塞尔Kassel.SetParent(f)
	
	f.利珀Lippe = BLippe利珀
	f.利珀Lippe.SetParent(f)
	
	f.诺特海姆Northeim = BNortheim诺特海姆
	f.诺特海姆Northeim.SetParent(f)
	
	f.帕德博恩Paderborn = BPaderborn帕德博恩
	f.帕德博恩Paderborn.SetParent(f)
	
}
