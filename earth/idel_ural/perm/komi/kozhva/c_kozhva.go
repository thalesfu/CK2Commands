package kozhva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KozhvaCounty interface {
    feud.County
    BBerozovka别罗佐夫卡() 	feud.Barony
    BByzovaya贝佐瓦亚() 	feud.Barony
    BChikshino奇克希诺() 	feud.Barony
    BIz_yayu伊兹亚尤() 	feud.Barony
    BKozhva科日瓦() 	feud.Barony
    BOzornyy奥佐尔内() 	feud.Barony
    BPuteyets普捷耶茨() 	feud.Barony
}

type 科日瓦KozhvaCounty struct {
	feud.BaseCounty
	别罗佐夫卡Berozovka 	feud.Barony
	贝佐瓦亚Byzovaya 	feud.Barony
	奇克希诺Chikshino 	feud.Barony
	伊兹亚尤Iz_yayu 	feud.Barony
	科日瓦Kozhva 	feud.Barony
	奥佐尔内Ozornyy 	feud.Barony
	普捷耶茨Puteyets 	feud.Barony
}

func (c *科日瓦KozhvaCounty) BBerozovka别罗佐夫卡() feud.Barony {
	return c.别罗佐夫卡Berozovka
}
    
func (c *科日瓦KozhvaCounty) BByzovaya贝佐瓦亚() feud.Barony {
	return c.贝佐瓦亚Byzovaya
}
    
func (c *科日瓦KozhvaCounty) BChikshino奇克希诺() feud.Barony {
	return c.奇克希诺Chikshino
}
    
func (c *科日瓦KozhvaCounty) BIz_yayu伊兹亚尤() feud.Barony {
	return c.伊兹亚尤Iz_yayu
}
    
func (c *科日瓦KozhvaCounty) BKozhva科日瓦() feud.Barony {
	return c.科日瓦Kozhva
}
    
func (c *科日瓦KozhvaCounty) BOzornyy奥佐尔内() feud.Barony {
	return c.奥佐尔内Ozornyy
}
    
func (c *科日瓦KozhvaCounty) BPuteyets普捷耶茨() feud.Barony {
	return c.普捷耶茨Puteyets
}
    
var CKozhva科日瓦 KozhvaCounty = &科日瓦KozhvaCounty{}

func init() {
	f := CKozhva科日瓦.(*科日瓦KozhvaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1834",
		Title:     "kozhva",
		TitleName: "科日瓦",
		TitleCode: "c_kozhva",
		Baronies:  map[string]feud.Barony{},
	}

	f.别罗佐夫卡Berozovka = BBerozovka别罗佐夫卡
	f.别罗佐夫卡Berozovka.SetParent(f)
	
	f.贝佐瓦亚Byzovaya = BByzovaya贝佐瓦亚
	f.贝佐瓦亚Byzovaya.SetParent(f)
	
	f.奇克希诺Chikshino = BChikshino奇克希诺
	f.奇克希诺Chikshino.SetParent(f)
	
	f.伊兹亚尤Iz_yayu = BIz_yayu伊兹亚尤
	f.伊兹亚尤Iz_yayu.SetParent(f)
	
	f.科日瓦Kozhva = BKozhva科日瓦
	f.科日瓦Kozhva.SetParent(f)
	
	f.奥佐尔内Ozornyy = BOzornyy奥佐尔内
	f.奥佐尔内Ozornyy.SetParent(f)
	
	f.普捷耶茨Puteyets = BPuteyets普捷耶茨
	f.普捷耶茨Puteyets.SetParent(f)
	
}
