package paro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ParoCounty interface {
    feud.County
    BDamthang当塘() 	feud.Barony
    BGasa加萨() 	feud.Barony
    BHungrel洪热() 	feud.Barony
    BKyichu祈楚() 	feud.Barony
    BParo帕罗() 	feud.Barony
    BTaktshang虎穴() 	feud.Barony
    BTshalunang涉鲁囊() 	feud.Barony
}

type 帕罗ParoCounty struct {
	feud.BaseCounty
	当塘Damthang 	feud.Barony
	加萨Gasa 	feud.Barony
	洪热Hungrel 	feud.Barony
	祈楚Kyichu 	feud.Barony
	帕罗Paro 	feud.Barony
	虎穴Taktshang 	feud.Barony
	涉鲁囊Tshalunang 	feud.Barony
}

func (c *帕罗ParoCounty) BDamthang当塘() feud.Barony {
	return c.当塘Damthang
}
    
func (c *帕罗ParoCounty) BGasa加萨() feud.Barony {
	return c.加萨Gasa
}
    
func (c *帕罗ParoCounty) BHungrel洪热() feud.Barony {
	return c.洪热Hungrel
}
    
func (c *帕罗ParoCounty) BKyichu祈楚() feud.Barony {
	return c.祈楚Kyichu
}
    
func (c *帕罗ParoCounty) BParo帕罗() feud.Barony {
	return c.帕罗Paro
}
    
func (c *帕罗ParoCounty) BTaktshang虎穴() feud.Barony {
	return c.虎穴Taktshang
}
    
func (c *帕罗ParoCounty) BTshalunang涉鲁囊() feud.Barony {
	return c.涉鲁囊Tshalunang
}
    
var CParo帕罗 ParoCounty = &帕罗ParoCounty{}

func init() {
	f := CParo帕罗.(*帕罗ParoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1481",
		Title:     "paro",
		TitleName: "帕罗",
		TitleCode: "c_paro",
		Baronies:  map[string]feud.Barony{},
	}

	f.当塘Damthang = BDamthang当塘
	f.当塘Damthang.SetParent(f)
	
	f.加萨Gasa = BGasa加萨
	f.加萨Gasa.SetParent(f)
	
	f.洪热Hungrel = BHungrel洪热
	f.洪热Hungrel.SetParent(f)
	
	f.祈楚Kyichu = BKyichu祈楚
	f.祈楚Kyichu.SetParent(f)
	
	f.帕罗Paro = BParo帕罗
	f.帕罗Paro.SetParent(f)
	
	f.虎穴Taktshang = BTaktshang虎穴
	f.虎穴Taktshang.SetParent(f)
	
	f.涉鲁囊Tshalunang = BTshalunang涉鲁囊
	f.涉鲁囊Tshalunang.SetParent(f)
	
}
