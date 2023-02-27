package munda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MundaCounty interface {
    feud.County
    BChutia周底耶() 	feud.Barony
    BMarang_buru摩烂迦补卢() 	feud.Barony
    BMelmarudun摩摩卢顿() 	feud.Barony
    BNathorl那兜() 	feud.Barony
    BPatera波谛罗() 	feud.Barony
    BPatho波豆() 	feud.Barony
    BRatu罗兜() 	feud.Barony
}

type 文荼MundaCounty struct {
	feud.BaseCounty
	周底耶Chutia 	feud.Barony
	摩烂迦补卢Marang_buru 	feud.Barony
	摩摩卢顿Melmarudun 	feud.Barony
	那兜Nathorl 	feud.Barony
	波谛罗Patera 	feud.Barony
	波豆Patho 	feud.Barony
	罗兜Ratu 	feud.Barony
}

func (c *文荼MundaCounty) BChutia周底耶() feud.Barony {
	return c.周底耶Chutia
}
    
func (c *文荼MundaCounty) BMarang_buru摩烂迦补卢() feud.Barony {
	return c.摩烂迦补卢Marang_buru
}
    
func (c *文荼MundaCounty) BMelmarudun摩摩卢顿() feud.Barony {
	return c.摩摩卢顿Melmarudun
}
    
func (c *文荼MundaCounty) BNathorl那兜() feud.Barony {
	return c.那兜Nathorl
}
    
func (c *文荼MundaCounty) BPatera波谛罗() feud.Barony {
	return c.波谛罗Patera
}
    
func (c *文荼MundaCounty) BPatho波豆() feud.Barony {
	return c.波豆Patho
}
    
func (c *文荼MundaCounty) BRatu罗兜() feud.Barony {
	return c.罗兜Ratu
}
    
var CMunda文荼 MundaCounty = &文荼MundaCounty{}

func init() {
	f := CMunda文荼.(*文荼MundaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1248",
		Title:     "munda",
		TitleName: "文荼",
		TitleCode: "c_munda",
		Baronies:  map[string]feud.Barony{},
	}

	f.周底耶Chutia = BChutia周底耶
	f.周底耶Chutia.SetParent(f)
	
	f.摩烂迦补卢Marang_buru = BMarang_buru摩烂迦补卢
	f.摩烂迦补卢Marang_buru.SetParent(f)
	
	f.摩摩卢顿Melmarudun = BMelmarudun摩摩卢顿
	f.摩摩卢顿Melmarudun.SetParent(f)
	
	f.那兜Nathorl = BNathorl那兜
	f.那兜Nathorl.SetParent(f)
	
	f.波谛罗Patera = BPatera波谛罗
	f.波谛罗Patera.SetParent(f)
	
	f.波豆Patho = BPatho波豆
	f.波豆Patho.SetParent(f)
	
	f.罗兜Ratu = BRatu罗兜
	f.罗兜Ratu.SetParent(f)
	
}
