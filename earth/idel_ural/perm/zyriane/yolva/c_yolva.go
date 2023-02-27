package yolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YolvaCounty interface {
    feud.County
    BChub丘布() 	feud.Barony
    BEdva埃德瓦() 	feud.Barony
    BKoin科因() 	feud.Barony
    BPizma皮兹马() 	feud.Barony
    BVorykva沃雷克瓦() 	feud.Barony
    BVym维姆() 	feud.Barony
    BYolva约尔瓦() 	feud.Barony
}

type 维姆YolvaCounty struct {
	feud.BaseCounty
	丘布Chub 	feud.Barony
	埃德瓦Edva 	feud.Barony
	科因Koin 	feud.Barony
	皮兹马Pizma 	feud.Barony
	沃雷克瓦Vorykva 	feud.Barony
	维姆Vym 	feud.Barony
	约尔瓦Yolva 	feud.Barony
}

func (c *维姆YolvaCounty) BChub丘布() feud.Barony {
	return c.丘布Chub
}
    
func (c *维姆YolvaCounty) BEdva埃德瓦() feud.Barony {
	return c.埃德瓦Edva
}
    
func (c *维姆YolvaCounty) BKoin科因() feud.Barony {
	return c.科因Koin
}
    
func (c *维姆YolvaCounty) BPizma皮兹马() feud.Barony {
	return c.皮兹马Pizma
}
    
func (c *维姆YolvaCounty) BVorykva沃雷克瓦() feud.Barony {
	return c.沃雷克瓦Vorykva
}
    
func (c *维姆YolvaCounty) BVym维姆() feud.Barony {
	return c.维姆Vym
}
    
func (c *维姆YolvaCounty) BYolva约尔瓦() feud.Barony {
	return c.约尔瓦Yolva
}
    
var CYolva维姆 YolvaCounty = &维姆YolvaCounty{}

func init() {
	f := CYolva维姆.(*维姆YolvaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1839",
		Title:     "yolva",
		TitleName: "维姆",
		TitleCode: "c_yolva",
		Baronies:  map[string]feud.Barony{},
	}

	f.丘布Chub = BChub丘布
	f.丘布Chub.SetParent(f)
	
	f.埃德瓦Edva = BEdva埃德瓦
	f.埃德瓦Edva.SetParent(f)
	
	f.科因Koin = BKoin科因
	f.科因Koin.SetParent(f)
	
	f.皮兹马Pizma = BPizma皮兹马
	f.皮兹马Pizma.SetParent(f)
	
	f.沃雷克瓦Vorykva = BVorykva沃雷克瓦
	f.沃雷克瓦Vorykva.SetParent(f)
	
	f.维姆Vym = BVym维姆
	f.维姆Vym.SetParent(f)
	
	f.约尔瓦Yolva = BYolva约尔瓦
	f.约尔瓦Yolva.SetParent(f)
	
}
