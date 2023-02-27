package ouadane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OuadaneCounty interface {
    feud.County
    BAzuggi阿祖吉() 	feud.Barony
    BGag加格() 	feud.Barony
    BIdjil伊吉勒() 	feud.Barony
    BKaliki卡利基() 	feud.Barony
    BKsaralkiali基阿里堡() 	feud.Barony
    BOuadane瓦丹() 	feud.Barony
    BWadan瓦丹() 	feud.Barony
    BWadane瓦达内() 	feud.Barony
}

type 瓦丹OuadaneCounty struct {
	feud.BaseCounty
	阿祖吉Azuggi 	feud.Barony
	加格Gag 	feud.Barony
	伊吉勒Idjil 	feud.Barony
	卡利基Kaliki 	feud.Barony
	基阿里堡Ksaralkiali 	feud.Barony
	瓦丹Ouadane 	feud.Barony
	瓦丹Wadan 	feud.Barony
	瓦达内Wadane 	feud.Barony
}

func (c *瓦丹OuadaneCounty) BAzuggi阿祖吉() feud.Barony {
	return c.阿祖吉Azuggi
}
    
func (c *瓦丹OuadaneCounty) BGag加格() feud.Barony {
	return c.加格Gag
}
    
func (c *瓦丹OuadaneCounty) BIdjil伊吉勒() feud.Barony {
	return c.伊吉勒Idjil
}
    
func (c *瓦丹OuadaneCounty) BKaliki卡利基() feud.Barony {
	return c.卡利基Kaliki
}
    
func (c *瓦丹OuadaneCounty) BKsaralkiali基阿里堡() feud.Barony {
	return c.基阿里堡Ksaralkiali
}
    
func (c *瓦丹OuadaneCounty) BOuadane瓦丹() feud.Barony {
	return c.瓦丹Ouadane
}
    
func (c *瓦丹OuadaneCounty) BWadan瓦丹() feud.Barony {
	return c.瓦丹Wadan
}
    
func (c *瓦丹OuadaneCounty) BWadane瓦达内() feud.Barony {
	return c.瓦达内Wadane
}
    
var COuadane瓦丹 OuadaneCounty = &瓦丹OuadaneCounty{}

func init() {
	f := COuadane瓦丹.(*瓦丹OuadaneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "920",
		Title:     "ouadane",
		TitleName: "瓦丹",
		TitleCode: "c_ouadane",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿祖吉Azuggi = BAzuggi阿祖吉
	f.阿祖吉Azuggi.SetParent(f)
	
	f.加格Gag = BGag加格
	f.加格Gag.SetParent(f)
	
	f.伊吉勒Idjil = BIdjil伊吉勒
	f.伊吉勒Idjil.SetParent(f)
	
	f.卡利基Kaliki = BKaliki卡利基
	f.卡利基Kaliki.SetParent(f)
	
	f.基阿里堡Ksaralkiali = BKsaralkiali基阿里堡
	f.基阿里堡Ksaralkiali.SetParent(f)
	
	f.瓦丹Ouadane = BOuadane瓦丹
	f.瓦丹Ouadane.SetParent(f)
	
	f.瓦丹Wadan = BWadan瓦丹
	f.瓦丹Wadan.SetParent(f)
	
	f.瓦达内Wadane = BWadane瓦达内
	f.瓦达内Wadane.SetParent(f)
	
}
