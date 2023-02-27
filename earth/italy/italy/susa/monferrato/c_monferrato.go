package monferrato

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MonferratoCounty interface {
    feud.County
    BAcqui阿奎() 	feud.Barony
    BAlba阿尔巴() 	feud.Barony
    BAsti阿斯蒂() 	feud.Barony
    BCanelli卡内利() 	feud.Barony
    BCasale卡萨莱() 	feud.Barony
    BMonferrato蒙费拉托() 	feud.Barony
    BTortona托尔托纳() 	feud.Barony
}

type 蒙费拉托MonferratoCounty struct {
	feud.BaseCounty
	阿奎Acqui 	feud.Barony
	阿尔巴Alba 	feud.Barony
	阿斯蒂Asti 	feud.Barony
	卡内利Canelli 	feud.Barony
	卡萨莱Casale 	feud.Barony
	蒙费拉托Monferrato 	feud.Barony
	托尔托纳Tortona 	feud.Barony
}

func (c *蒙费拉托MonferratoCounty) BAcqui阿奎() feud.Barony {
	return c.阿奎Acqui
}
    
func (c *蒙费拉托MonferratoCounty) BAlba阿尔巴() feud.Barony {
	return c.阿尔巴Alba
}
    
func (c *蒙费拉托MonferratoCounty) BAsti阿斯蒂() feud.Barony {
	return c.阿斯蒂Asti
}
    
func (c *蒙费拉托MonferratoCounty) BCanelli卡内利() feud.Barony {
	return c.卡内利Canelli
}
    
func (c *蒙费拉托MonferratoCounty) BCasale卡萨莱() feud.Barony {
	return c.卡萨莱Casale
}
    
func (c *蒙费拉托MonferratoCounty) BMonferrato蒙费拉托() feud.Barony {
	return c.蒙费拉托Monferrato
}
    
func (c *蒙费拉托MonferratoCounty) BTortona托尔托纳() feud.Barony {
	return c.托尔托纳Tortona
}
    
var CMonferrato蒙费拉托 MonferratoCounty = &蒙费拉托MonferratoCounty{}

func init() {
	f := CMonferrato蒙费拉托.(*蒙费拉托MonferratoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "232",
		Title:     "monferrato",
		TitleName: "蒙费拉托",
		TitleCode: "c_monferrato",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿奎Acqui = BAcqui阿奎
	f.阿奎Acqui.SetParent(f)
	
	f.阿尔巴Alba = BAlba阿尔巴
	f.阿尔巴Alba.SetParent(f)
	
	f.阿斯蒂Asti = BAsti阿斯蒂
	f.阿斯蒂Asti.SetParent(f)
	
	f.卡内利Canelli = BCanelli卡内利
	f.卡内利Canelli.SetParent(f)
	
	f.卡萨莱Casale = BCasale卡萨莱
	f.卡萨莱Casale.SetParent(f)
	
	f.蒙费拉托Monferrato = BMonferrato蒙费拉托
	f.蒙费拉托Monferrato.SetParent(f)
	
	f.托尔托纳Tortona = BTortona托尔托纳
	f.托尔托纳Tortona.SetParent(f)
	
}
