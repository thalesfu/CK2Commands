package daman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DamanCounty interface {
    feud.County
    BBulsar补萨尔() 	feud.Barony
    BDaman陀曼那() 	feud.Barony
    BNagar_haveli那迦诃伐尼() 	feud.Barony
    BPardi波尔提() 	feud.Barony
    BRamnagar拉姆那迦() 	feud.Barony
    BSanjan纳加尔() 	feud.Barony
    BSurparaka首波罗() 	feud.Barony
}

type 陀曼那DamanCounty struct {
	feud.BaseCounty
	补萨尔Bulsar 	feud.Barony
	陀曼那Daman 	feud.Barony
	那迦诃伐尼Nagar_haveli 	feud.Barony
	波尔提Pardi 	feud.Barony
	拉姆那迦Ramnagar 	feud.Barony
	纳加尔Sanjan 	feud.Barony
	首波罗Surparaka 	feud.Barony
}

func (c *陀曼那DamanCounty) BBulsar补萨尔() feud.Barony {
	return c.补萨尔Bulsar
}
    
func (c *陀曼那DamanCounty) BDaman陀曼那() feud.Barony {
	return c.陀曼那Daman
}
    
func (c *陀曼那DamanCounty) BNagar_haveli那迦诃伐尼() feud.Barony {
	return c.那迦诃伐尼Nagar_haveli
}
    
func (c *陀曼那DamanCounty) BPardi波尔提() feud.Barony {
	return c.波尔提Pardi
}
    
func (c *陀曼那DamanCounty) BRamnagar拉姆那迦() feud.Barony {
	return c.拉姆那迦Ramnagar
}
    
func (c *陀曼那DamanCounty) BSanjan纳加尔() feud.Barony {
	return c.纳加尔Sanjan
}
    
func (c *陀曼那DamanCounty) BSurparaka首波罗() feud.Barony {
	return c.首波罗Surparaka
}
    
var CDaman陀曼那 DamanCounty = &陀曼那DamanCounty{}

func init() {
	f := CDaman陀曼那.(*陀曼那DamanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1126",
		Title:     "daman",
		TitleName: "陀曼那",
		TitleCode: "c_daman",
		Baronies:  map[string]feud.Barony{},
	}

	f.补萨尔Bulsar = BBulsar补萨尔
	f.补萨尔Bulsar.SetParent(f)
	
	f.陀曼那Daman = BDaman陀曼那
	f.陀曼那Daman.SetParent(f)
	
	f.那迦诃伐尼Nagar_haveli = BNagar_haveli那迦诃伐尼
	f.那迦诃伐尼Nagar_haveli.SetParent(f)
	
	f.波尔提Pardi = BPardi波尔提
	f.波尔提Pardi.SetParent(f)
	
	f.拉姆那迦Ramnagar = BRamnagar拉姆那迦
	f.拉姆那迦Ramnagar.SetParent(f)
	
	f.纳加尔Sanjan = BSanjan纳加尔
	f.纳加尔Sanjan.SetParent(f)
	
	f.首波罗Surparaka = BSurparaka首波罗
	f.首波罗Surparaka.SetParent(f)
	
}
