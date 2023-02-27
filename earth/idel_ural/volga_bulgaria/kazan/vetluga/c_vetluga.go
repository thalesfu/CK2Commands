package vetluga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VetlugaCounty interface {
    feud.County
    BChakhunia察胡尼亚() 	feud.Barony
    BCharia沙里亚() 	feud.Barony
    BKajirv卡伊尔夫() 	feud.Barony
    BOuren乌连() 	feud.Barony
    BSyava夏瓦() 	feud.Barony
    BVetluga韦特卢加() 	feud.Barony
    BVokhma沃赫马() 	feud.Barony
}

type 韦特卢加VetlugaCounty struct {
	feud.BaseCounty
	察胡尼亚Chakhunia 	feud.Barony
	沙里亚Charia 	feud.Barony
	卡伊尔夫Kajirv 	feud.Barony
	乌连Ouren 	feud.Barony
	夏瓦Syava 	feud.Barony
	韦特卢加Vetluga 	feud.Barony
	沃赫马Vokhma 	feud.Barony
}

func (c *韦特卢加VetlugaCounty) BChakhunia察胡尼亚() feud.Barony {
	return c.察胡尼亚Chakhunia
}
    
func (c *韦特卢加VetlugaCounty) BCharia沙里亚() feud.Barony {
	return c.沙里亚Charia
}
    
func (c *韦特卢加VetlugaCounty) BKajirv卡伊尔夫() feud.Barony {
	return c.卡伊尔夫Kajirv
}
    
func (c *韦特卢加VetlugaCounty) BOuren乌连() feud.Barony {
	return c.乌连Ouren
}
    
func (c *韦特卢加VetlugaCounty) BSyava夏瓦() feud.Barony {
	return c.夏瓦Syava
}
    
func (c *韦特卢加VetlugaCounty) BVetluga韦特卢加() feud.Barony {
	return c.韦特卢加Vetluga
}
    
func (c *韦特卢加VetlugaCounty) BVokhma沃赫马() feud.Barony {
	return c.沃赫马Vokhma
}
    
var CVetluga韦特卢加 VetlugaCounty = &韦特卢加VetlugaCounty{}

func init() {
	f := CVetluga韦特卢加.(*韦特卢加VetlugaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1715",
		Title:     "vetluga",
		TitleName: "韦特卢加",
		TitleCode: "c_vetluga",
		Baronies:  map[string]feud.Barony{},
	}

	f.察胡尼亚Chakhunia = BChakhunia察胡尼亚
	f.察胡尼亚Chakhunia.SetParent(f)
	
	f.沙里亚Charia = BCharia沙里亚
	f.沙里亚Charia.SetParent(f)
	
	f.卡伊尔夫Kajirv = BKajirv卡伊尔夫
	f.卡伊尔夫Kajirv.SetParent(f)
	
	f.乌连Ouren = BOuren乌连
	f.乌连Ouren.SetParent(f)
	
	f.夏瓦Syava = BSyava夏瓦
	f.夏瓦Syava.SetParent(f)
	
	f.韦特卢加Vetluga = BVetluga韦特卢加
	f.韦特卢加Vetluga.SetParent(f)
	
	f.沃赫马Vokhma = BVokhma沃赫马
	f.沃赫马Vokhma.SetParent(f)
	
}
