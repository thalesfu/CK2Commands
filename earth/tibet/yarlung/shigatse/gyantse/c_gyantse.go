package gyantse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GyantseCounty interface {
    feud.County
    BBainang白朗() 	feud.Barony
    BGyantse江孜() 	feud.Barony
    BKangmar康马() 	feud.Barony
    BKarmai卡麦() 	feud.Barony
    BPalcho白居() 	feud.Barony
    BRinbung仁布() 	feud.Barony
    BYatung亚东() 	feud.Barony
}

type 江孜GyantseCounty struct {
	feud.BaseCounty
	白朗Bainang 	feud.Barony
	江孜Gyantse 	feud.Barony
	康马Kangmar 	feud.Barony
	卡麦Karmai 	feud.Barony
	白居Palcho 	feud.Barony
	仁布Rinbung 	feud.Barony
	亚东Yatung 	feud.Barony
}

func (c *江孜GyantseCounty) BBainang白朗() feud.Barony {
	return c.白朗Bainang
}
    
func (c *江孜GyantseCounty) BGyantse江孜() feud.Barony {
	return c.江孜Gyantse
}
    
func (c *江孜GyantseCounty) BKangmar康马() feud.Barony {
	return c.康马Kangmar
}
    
func (c *江孜GyantseCounty) BKarmai卡麦() feud.Barony {
	return c.卡麦Karmai
}
    
func (c *江孜GyantseCounty) BPalcho白居() feud.Barony {
	return c.白居Palcho
}
    
func (c *江孜GyantseCounty) BRinbung仁布() feud.Barony {
	return c.仁布Rinbung
}
    
func (c *江孜GyantseCounty) BYatung亚东() feud.Barony {
	return c.亚东Yatung
}
    
var CGyantse江孜 GyantseCounty = &江孜GyantseCounty{}

func init() {
	f := CGyantse江孜.(*江孜GyantseCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1497",
		Title:     "gyantse",
		TitleName: "江孜",
		TitleCode: "c_gyantse",
		Baronies:  map[string]feud.Barony{},
	}

	f.白朗Bainang = BBainang白朗
	f.白朗Bainang.SetParent(f)
	
	f.江孜Gyantse = BGyantse江孜
	f.江孜Gyantse.SetParent(f)
	
	f.康马Kangmar = BKangmar康马
	f.康马Kangmar.SetParent(f)
	
	f.卡麦Karmai = BKarmai卡麦
	f.卡麦Karmai.SetParent(f)
	
	f.白居Palcho = BPalcho白居
	f.白居Palcho.SetParent(f)
	
	f.仁布Rinbung = BRinbung仁布
	f.仁布Rinbung.SetParent(f)
	
	f.亚东Yatung = BYatung亚东
	f.亚东Yatung.SetParent(f)
	
}
