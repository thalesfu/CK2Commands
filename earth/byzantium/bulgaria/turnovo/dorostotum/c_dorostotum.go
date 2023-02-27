package dorostotum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DorostotumCounty interface {
    feud.County
    BBorovo博罗沃() 	feud.Barony
    BByaladve比亚拉德韦() 	feud.Barony
    BDorosturum杜罗斯托鲁姆() 	feud.Barony
    BRusi鲁希() 	feud.Barony
    BSamuil萨穆伊尔() 	feud.Barony
    BShumen舒门() 	feud.Barony
    BSlivopole斯利沃波莱() 	feud.Barony
    BTsenovo采诺沃() 	feud.Barony
}

type 杜罗斯托鲁姆DorostotumCounty struct {
	feud.BaseCounty
	博罗沃Borovo 	feud.Barony
	比亚拉德韦Byaladve 	feud.Barony
	杜罗斯托鲁姆Dorosturum 	feud.Barony
	鲁希Rusi 	feud.Barony
	萨穆伊尔Samuil 	feud.Barony
	舒门Shumen 	feud.Barony
	斯利沃波莱Slivopole 	feud.Barony
	采诺沃Tsenovo 	feud.Barony
}

func (c *杜罗斯托鲁姆DorostotumCounty) BBorovo博罗沃() feud.Barony {
	return c.博罗沃Borovo
}
    
func (c *杜罗斯托鲁姆DorostotumCounty) BByaladve比亚拉德韦() feud.Barony {
	return c.比亚拉德韦Byaladve
}
    
func (c *杜罗斯托鲁姆DorostotumCounty) BDorosturum杜罗斯托鲁姆() feud.Barony {
	return c.杜罗斯托鲁姆Dorosturum
}
    
func (c *杜罗斯托鲁姆DorostotumCounty) BRusi鲁希() feud.Barony {
	return c.鲁希Rusi
}
    
func (c *杜罗斯托鲁姆DorostotumCounty) BSamuil萨穆伊尔() feud.Barony {
	return c.萨穆伊尔Samuil
}
    
func (c *杜罗斯托鲁姆DorostotumCounty) BShumen舒门() feud.Barony {
	return c.舒门Shumen
}
    
func (c *杜罗斯托鲁姆DorostotumCounty) BSlivopole斯利沃波莱() feud.Barony {
	return c.斯利沃波莱Slivopole
}
    
func (c *杜罗斯托鲁姆DorostotumCounty) BTsenovo采诺沃() feud.Barony {
	return c.采诺沃Tsenovo
}
    
var CDorostotum杜罗斯托鲁姆 DorostotumCounty = &杜罗斯托鲁姆DorostotumCounty{}

func init() {
	f := CDorostotum杜罗斯托鲁姆.(*杜罗斯托鲁姆DorostotumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "508",
		Title:     "dorostotum",
		TitleName: "杜罗斯托鲁姆",
		TitleCode: "c_dorostotum",
		Baronies:  map[string]feud.Barony{},
	}

	f.博罗沃Borovo = BBorovo博罗沃
	f.博罗沃Borovo.SetParent(f)
	
	f.比亚拉德韦Byaladve = BByaladve比亚拉德韦
	f.比亚拉德韦Byaladve.SetParent(f)
	
	f.杜罗斯托鲁姆Dorosturum = BDorosturum杜罗斯托鲁姆
	f.杜罗斯托鲁姆Dorosturum.SetParent(f)
	
	f.鲁希Rusi = BRusi鲁希
	f.鲁希Rusi.SetParent(f)
	
	f.萨穆伊尔Samuil = BSamuil萨穆伊尔
	f.萨穆伊尔Samuil.SetParent(f)
	
	f.舒门Shumen = BShumen舒门
	f.舒门Shumen.SetParent(f)
	
	f.斯利沃波莱Slivopole = BSlivopole斯利沃波莱
	f.斯利沃波莱Slivopole.SetParent(f)
	
	f.采诺沃Tsenovo = BTsenovo采诺沃
	f.采诺沃Tsenovo.SetParent(f)
	
}
