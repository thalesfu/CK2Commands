package poznanskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PoznanskieCounty interface {
    feud.County
    BCzarnkow恰恩库夫() 	feud.Barony
    BKoscian科希强() 	feud.Barony
    BMiedzyrzecz缅济热奇() 	feud.Barony
    BOborniki奥博尔尼基() 	feud.Barony
    BPoznan波兹南() 	feud.Barony
    BWalcz瓦乌奇() 	feud.Barony
    BWschowa弗斯霍瓦() 	feud.Barony
}

type 波兹南PoznanskieCounty struct {
	feud.BaseCounty
	恰恩库夫Czarnkow 	feud.Barony
	科希强Koscian 	feud.Barony
	缅济热奇Miedzyrzecz 	feud.Barony
	奥博尔尼基Oborniki 	feud.Barony
	波兹南Poznan 	feud.Barony
	瓦乌奇Walcz 	feud.Barony
	弗斯霍瓦Wschowa 	feud.Barony
}

func (c *波兹南PoznanskieCounty) BCzarnkow恰恩库夫() feud.Barony {
	return c.恰恩库夫Czarnkow
}
    
func (c *波兹南PoznanskieCounty) BKoscian科希强() feud.Barony {
	return c.科希强Koscian
}
    
func (c *波兹南PoznanskieCounty) BMiedzyrzecz缅济热奇() feud.Barony {
	return c.缅济热奇Miedzyrzecz
}
    
func (c *波兹南PoznanskieCounty) BOborniki奥博尔尼基() feud.Barony {
	return c.奥博尔尼基Oborniki
}
    
func (c *波兹南PoznanskieCounty) BPoznan波兹南() feud.Barony {
	return c.波兹南Poznan
}
    
func (c *波兹南PoznanskieCounty) BWalcz瓦乌奇() feud.Barony {
	return c.瓦乌奇Walcz
}
    
func (c *波兹南PoznanskieCounty) BWschowa弗斯霍瓦() feud.Barony {
	return c.弗斯霍瓦Wschowa
}
    
var CPoznanskie波兹南 PoznanskieCounty = &波兹南PoznanskieCounty{}

func init() {
	f := CPoznanskie波兹南.(*波兹南PoznanskieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "431",
		Title:     "poznanskie",
		TitleName: "波兹南",
		TitleCode: "c_poznanskie",
		Baronies:  map[string]feud.Barony{},
	}

	f.恰恩库夫Czarnkow = BCzarnkow恰恩库夫
	f.恰恩库夫Czarnkow.SetParent(f)
	
	f.科希强Koscian = BKoscian科希强
	f.科希强Koscian.SetParent(f)
	
	f.缅济热奇Miedzyrzecz = BMiedzyrzecz缅济热奇
	f.缅济热奇Miedzyrzecz.SetParent(f)
	
	f.奥博尔尼基Oborniki = BOborniki奥博尔尼基
	f.奥博尔尼基Oborniki.SetParent(f)
	
	f.波兹南Poznan = BPoznan波兹南
	f.波兹南Poznan.SetParent(f)
	
	f.瓦乌奇Walcz = BWalcz瓦乌奇
	f.瓦乌奇Walcz.SetParent(f)
	
	f.弗斯霍瓦Wschowa = BWschowa弗斯霍瓦
	f.弗斯霍瓦Wschowa.SetParent(f)
	
}
