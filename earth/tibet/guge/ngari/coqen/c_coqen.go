package coqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CoqenCounty interface {
    feud.County
    BAsog阿索() 	feud.Barony
    BCoqen措勤() 	feud.Barony
    BDawaxung达雄() 	feud.Barony
    BGyangrang江让() 	feud.Barony
    BMaindong门东() 	feud.Barony
    BTarungha特荣海() 	feud.Barony
    BZhuglung珠龙() 	feud.Barony
}

type 措勤CoqenCounty struct {
	feud.BaseCounty
	阿索Asog 	feud.Barony
	措勤Coqen 	feud.Barony
	达雄Dawaxung 	feud.Barony
	江让Gyangrang 	feud.Barony
	门东Maindong 	feud.Barony
	特荣海Tarungha 	feud.Barony
	珠龙Zhuglung 	feud.Barony
}

func (c *措勤CoqenCounty) BAsog阿索() feud.Barony {
	return c.阿索Asog
}
    
func (c *措勤CoqenCounty) BCoqen措勤() feud.Barony {
	return c.措勤Coqen
}
    
func (c *措勤CoqenCounty) BDawaxung达雄() feud.Barony {
	return c.达雄Dawaxung
}
    
func (c *措勤CoqenCounty) BGyangrang江让() feud.Barony {
	return c.江让Gyangrang
}
    
func (c *措勤CoqenCounty) BMaindong门东() feud.Barony {
	return c.门东Maindong
}
    
func (c *措勤CoqenCounty) BTarungha特荣海() feud.Barony {
	return c.特荣海Tarungha
}
    
func (c *措勤CoqenCounty) BZhuglung珠龙() feud.Barony {
	return c.珠龙Zhuglung
}
    
var CCoqen措勤 CoqenCounty = &措勤CoqenCounty{}

func init() {
	f := CCoqen措勤.(*措勤CoqenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1493",
		Title:     "coqen",
		TitleName: "措勤",
		TitleCode: "c_coqen",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿索Asog = BAsog阿索
	f.阿索Asog.SetParent(f)
	
	f.措勤Coqen = BCoqen措勤
	f.措勤Coqen.SetParent(f)
	
	f.达雄Dawaxung = BDawaxung达雄
	f.达雄Dawaxung.SetParent(f)
	
	f.江让Gyangrang = BGyangrang江让
	f.江让Gyangrang.SetParent(f)
	
	f.门东Maindong = BMaindong门东
	f.门东Maindong.SetParent(f)
	
	f.特荣海Tarungha = BTarungha特荣海
	f.特荣海Tarungha.SetParent(f)
	
	f.珠龙Zhuglung = BZhuglung珠龙
	f.珠龙Zhuglung.SetParent(f)
	
}
