package dunois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DunoisCounty interface {
    feud.County
    BBonneval博讷瓦勒() 	feud.Barony
    BChateaudun沙托丹() 	feud.Barony
    BCloyes克卢瓦() 	feud.Barony
    BFreteval弗雷特瓦勒() 	feud.Barony
    BMarchenoir马什努瓦() 	feud.Barony
    BPatay帕泰() 	feud.Barony
    BTiron蒂龙() 	feud.Barony
}

type 迪努瓦DunoisCounty struct {
	feud.BaseCounty
	博讷瓦勒Bonneval 	feud.Barony
	沙托丹Chateaudun 	feud.Barony
	克卢瓦Cloyes 	feud.Barony
	弗雷特瓦勒Freteval 	feud.Barony
	马什努瓦Marchenoir 	feud.Barony
	帕泰Patay 	feud.Barony
	蒂龙Tiron 	feud.Barony
}

func (c *迪努瓦DunoisCounty) BBonneval博讷瓦勒() feud.Barony {
	return c.博讷瓦勒Bonneval
}
    
func (c *迪努瓦DunoisCounty) BChateaudun沙托丹() feud.Barony {
	return c.沙托丹Chateaudun
}
    
func (c *迪努瓦DunoisCounty) BCloyes克卢瓦() feud.Barony {
	return c.克卢瓦Cloyes
}
    
func (c *迪努瓦DunoisCounty) BFreteval弗雷特瓦勒() feud.Barony {
	return c.弗雷特瓦勒Freteval
}
    
func (c *迪努瓦DunoisCounty) BMarchenoir马什努瓦() feud.Barony {
	return c.马什努瓦Marchenoir
}
    
func (c *迪努瓦DunoisCounty) BPatay帕泰() feud.Barony {
	return c.帕泰Patay
}
    
func (c *迪努瓦DunoisCounty) BTiron蒂龙() feud.Barony {
	return c.蒂龙Tiron
}
    
var CDunois迪努瓦 DunoisCounty = &迪努瓦DunoisCounty{}

func init() {
	f := CDunois迪努瓦.(*迪努瓦DunoisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1965",
		Title:     "dunois",
		TitleName: "迪努瓦",
		TitleCode: "c_dunois",
		Baronies:  map[string]feud.Barony{},
	}

	f.博讷瓦勒Bonneval = BBonneval博讷瓦勒
	f.博讷瓦勒Bonneval.SetParent(f)
	
	f.沙托丹Chateaudun = BChateaudun沙托丹
	f.沙托丹Chateaudun.SetParent(f)
	
	f.克卢瓦Cloyes = BCloyes克卢瓦
	f.克卢瓦Cloyes.SetParent(f)
	
	f.弗雷特瓦勒Freteval = BFreteval弗雷特瓦勒
	f.弗雷特瓦勒Freteval.SetParent(f)
	
	f.马什努瓦Marchenoir = BMarchenoir马什努瓦
	f.马什努瓦Marchenoir.SetParent(f)
	
	f.帕泰Patay = BPatay帕泰
	f.帕泰Patay.SetParent(f)
	
	f.蒂龙Tiron = BTiron蒂龙
	f.蒂龙Tiron.SetParent(f)
	
}
