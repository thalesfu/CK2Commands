package otrar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OtrarCounty interface {
    feud.County
    BBirlik比尔利克() 	feud.Barony
    BChernak车尔纳克() 	feud.Barony
    BKentau肯套() 	feud.Barony
    BOtrar讹答剌() 	feud.Barony
    BShaulder绍利杰尔() 	feud.Barony
    BShoshkakoi朔什卡凯() 	feud.Barony
    BTurtkul图尔特库尔() 	feud.Barony
}

type 讹答剌OtrarCounty struct {
	feud.BaseCounty
	比尔利克Birlik 	feud.Barony
	车尔纳克Chernak 	feud.Barony
	肯套Kentau 	feud.Barony
	讹答剌Otrar 	feud.Barony
	绍利杰尔Shaulder 	feud.Barony
	朔什卡凯Shoshkakoi 	feud.Barony
	图尔特库尔Turtkul 	feud.Barony
}

func (c *讹答剌OtrarCounty) BBirlik比尔利克() feud.Barony {
	return c.比尔利克Birlik
}
    
func (c *讹答剌OtrarCounty) BChernak车尔纳克() feud.Barony {
	return c.车尔纳克Chernak
}
    
func (c *讹答剌OtrarCounty) BKentau肯套() feud.Barony {
	return c.肯套Kentau
}
    
func (c *讹答剌OtrarCounty) BOtrar讹答剌() feud.Barony {
	return c.讹答剌Otrar
}
    
func (c *讹答剌OtrarCounty) BShaulder绍利杰尔() feud.Barony {
	return c.绍利杰尔Shaulder
}
    
func (c *讹答剌OtrarCounty) BShoshkakoi朔什卡凯() feud.Barony {
	return c.朔什卡凯Shoshkakoi
}
    
func (c *讹答剌OtrarCounty) BTurtkul图尔特库尔() feud.Barony {
	return c.图尔特库尔Turtkul
}
    
var COtrar讹答剌 OtrarCounty = &讹答剌OtrarCounty{}

func init() {
	f := COtrar讹答剌.(*讹答剌OtrarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1431",
		Title:     "otrar",
		TitleName: "讹答剌",
		TitleCode: "c_otrar",
		Baronies:  map[string]feud.Barony{},
	}

	f.比尔利克Birlik = BBirlik比尔利克
	f.比尔利克Birlik.SetParent(f)
	
	f.车尔纳克Chernak = BChernak车尔纳克
	f.车尔纳克Chernak.SetParent(f)
	
	f.肯套Kentau = BKentau肯套
	f.肯套Kentau.SetParent(f)
	
	f.讹答剌Otrar = BOtrar讹答剌
	f.讹答剌Otrar.SetParent(f)
	
	f.绍利杰尔Shaulder = BShaulder绍利杰尔
	f.绍利杰尔Shaulder.SetParent(f)
	
	f.朔什卡凯Shoshkakoi = BShoshkakoi朔什卡凯
	f.朔什卡凯Shoshkakoi.SetParent(f)
	
	f.图尔特库尔Turtkul = BTurtkul图尔特库尔
	f.图尔特库尔Turtkul.SetParent(f)
	
}
