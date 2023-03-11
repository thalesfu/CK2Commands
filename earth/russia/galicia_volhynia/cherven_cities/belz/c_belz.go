package belz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BelzCounty interface {
    feud.County
    BBelz贝尔兹() 	feud.Barony
    BGrabowiec格拉博维茨() 	feud.Barony
    BHorodlo霍罗德沃() 	feud.Barony
    BLopatyn洛帕京() 	feud.Barony
    BMycow梅楚夫() 	feud.Barony
    BSokal索卡尔() 	feud.Barony
    BTchernogrov切尔诺格罗夫() 	feud.Barony
}

type 贝尔兹BelzCounty struct {
	feud.BaseCounty
	贝尔兹Belz 	feud.Barony
	格拉博维茨Grabowiec 	feud.Barony
	霍罗德沃Horodlo 	feud.Barony
	洛帕京Lopatyn 	feud.Barony
	梅楚夫Mycow 	feud.Barony
	索卡尔Sokal 	feud.Barony
	切尔诺格罗夫Tchernogrov 	feud.Barony
}

func (c *贝尔兹BelzCounty) BBelz贝尔兹() feud.Barony {
	return c.贝尔兹Belz
}
    
func (c *贝尔兹BelzCounty) BGrabowiec格拉博维茨() feud.Barony {
	return c.格拉博维茨Grabowiec
}
    
func (c *贝尔兹BelzCounty) BHorodlo霍罗德沃() feud.Barony {
	return c.霍罗德沃Horodlo
}
    
func (c *贝尔兹BelzCounty) BLopatyn洛帕京() feud.Barony {
	return c.洛帕京Lopatyn
}
    
func (c *贝尔兹BelzCounty) BMycow梅楚夫() feud.Barony {
	return c.梅楚夫Mycow
}
    
func (c *贝尔兹BelzCounty) BSokal索卡尔() feud.Barony {
	return c.索卡尔Sokal
}
    
func (c *贝尔兹BelzCounty) BTchernogrov切尔诺格罗夫() feud.Barony {
	return c.切尔诺格罗夫Tchernogrov
}
    
var CBelz贝尔兹 BelzCounty = &贝尔兹BelzCounty{}

func init() {
	f := CBelz贝尔兹.(*贝尔兹BelzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1632",
		Title:     "belz",
		TitleName: "贝尔兹",
		TitleCode: "c_belz",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝尔兹Belz = BBelz贝尔兹
	f.贝尔兹Belz.SetParent(f)
	
	f.格拉博维茨Grabowiec = BGrabowiec格拉博维茨
	f.格拉博维茨Grabowiec.SetParent(f)
	
	f.霍罗德沃Horodlo = BHorodlo霍罗德沃
	f.霍罗德沃Horodlo.SetParent(f)
	
	f.洛帕京Lopatyn = BLopatyn洛帕京
	f.洛帕京Lopatyn.SetParent(f)
	
	f.梅楚夫Mycow = BMycow梅楚夫
	f.梅楚夫Mycow.SetParent(f)
	
	f.索卡尔Sokal = BSokal索卡尔
	f.索卡尔Sokal.SetParent(f)
	
	f.切尔诺格罗夫Tchernogrov = BTchernogrov切尔诺格罗夫
	f.切尔诺格罗夫Tchernogrov.SetParent(f)
	
}
