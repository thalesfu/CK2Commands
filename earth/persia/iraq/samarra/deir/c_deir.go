package deir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DeirCounty interface {
    feud.County
    BAnbar安巴尔() 	feud.Barony
    BHit希特() 	feud.Barony
    BKasra卡撕拉() 	feud.Barony
    BNehardea尼赫底() 	feud.Barony
    BRamadi拉马迪() 	feud.Barony
    BRawa雷瓦() 	feud.Barony
    BRutbah鲁特巴赫() 	feud.Barony
    BTakrit提克里特() 	feud.Barony
}

type 代尔DeirCounty struct {
	feud.BaseCounty
	安巴尔Anbar 	feud.Barony
	希特Hit 	feud.Barony
	卡撕拉Kasra 	feud.Barony
	尼赫底Nehardea 	feud.Barony
	拉马迪Ramadi 	feud.Barony
	雷瓦Rawa 	feud.Barony
	鲁特巴赫Rutbah 	feud.Barony
	提克里特Takrit 	feud.Barony
}

func (c *代尔DeirCounty) BAnbar安巴尔() feud.Barony {
	return c.安巴尔Anbar
}
    
func (c *代尔DeirCounty) BHit希特() feud.Barony {
	return c.希特Hit
}
    
func (c *代尔DeirCounty) BKasra卡撕拉() feud.Barony {
	return c.卡撕拉Kasra
}
    
func (c *代尔DeirCounty) BNehardea尼赫底() feud.Barony {
	return c.尼赫底Nehardea
}
    
func (c *代尔DeirCounty) BRamadi拉马迪() feud.Barony {
	return c.拉马迪Ramadi
}
    
func (c *代尔DeirCounty) BRawa雷瓦() feud.Barony {
	return c.雷瓦Rawa
}
    
func (c *代尔DeirCounty) BRutbah鲁特巴赫() feud.Barony {
	return c.鲁特巴赫Rutbah
}
    
func (c *代尔DeirCounty) BTakrit提克里特() feud.Barony {
	return c.提克里特Takrit
}
    
var CDeir代尔 DeirCounty = &代尔DeirCounty{}

func init() {
	f := CDeir代尔.(*代尔DeirCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "695",
		Title:     "deir",
		TitleName: "代尔",
		TitleCode: "c_deir",
		Baronies:  map[string]feud.Barony{},
	}

	f.安巴尔Anbar = BAnbar安巴尔
	f.安巴尔Anbar.SetParent(f)
	
	f.希特Hit = BHit希特
	f.希特Hit.SetParent(f)
	
	f.卡撕拉Kasra = BKasra卡撕拉
	f.卡撕拉Kasra.SetParent(f)
	
	f.尼赫底Nehardea = BNehardea尼赫底
	f.尼赫底Nehardea.SetParent(f)
	
	f.拉马迪Ramadi = BRamadi拉马迪
	f.拉马迪Ramadi.SetParent(f)
	
	f.雷瓦Rawa = BRawa雷瓦
	f.雷瓦Rawa.SetParent(f)
	
	f.鲁特巴赫Rutbah = BRutbah鲁特巴赫
	f.鲁特巴赫Rutbah.SetParent(f)
	
	f.提克里特Takrit = BTakrit提克里特
	f.提克里特Takrit.SetParent(f)
	
}
