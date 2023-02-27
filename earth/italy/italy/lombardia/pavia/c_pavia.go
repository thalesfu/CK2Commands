package pavia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PaviaCounty interface {
    feud.County
    BAlessandria亚历山德里亚() 	feud.Barony
    BBobbio博比奥() 	feud.Barony
    BCasteggio卡斯泰焦() 	feud.Barony
    BColorno科洛尔诺() 	feud.Barony
    BMontebello蒙特贝洛() 	feud.Barony
    BPavia帕维亚() 	feud.Barony
    BPiacenza皮亚琴察() 	feud.Barony
    BVoghera沃盖拉() 	feud.Barony
}

type 帕维亚PaviaCounty struct {
	feud.BaseCounty
	亚历山德里亚Alessandria 	feud.Barony
	博比奥Bobbio 	feud.Barony
	卡斯泰焦Casteggio 	feud.Barony
	科洛尔诺Colorno 	feud.Barony
	蒙特贝洛Montebello 	feud.Barony
	帕维亚Pavia 	feud.Barony
	皮亚琴察Piacenza 	feud.Barony
	沃盖拉Voghera 	feud.Barony
}

func (c *帕维亚PaviaCounty) BAlessandria亚历山德里亚() feud.Barony {
	return c.亚历山德里亚Alessandria
}
    
func (c *帕维亚PaviaCounty) BBobbio博比奥() feud.Barony {
	return c.博比奥Bobbio
}
    
func (c *帕维亚PaviaCounty) BCasteggio卡斯泰焦() feud.Barony {
	return c.卡斯泰焦Casteggio
}
    
func (c *帕维亚PaviaCounty) BColorno科洛尔诺() feud.Barony {
	return c.科洛尔诺Colorno
}
    
func (c *帕维亚PaviaCounty) BMontebello蒙特贝洛() feud.Barony {
	return c.蒙特贝洛Montebello
}
    
func (c *帕维亚PaviaCounty) BPavia帕维亚() feud.Barony {
	return c.帕维亚Pavia
}
    
func (c *帕维亚PaviaCounty) BPiacenza皮亚琴察() feud.Barony {
	return c.皮亚琴察Piacenza
}
    
func (c *帕维亚PaviaCounty) BVoghera沃盖拉() feud.Barony {
	return c.沃盖拉Voghera
}
    
var CPavia帕维亚 PaviaCounty = &帕维亚PaviaCounty{}

func init() {
	f := CPavia帕维亚.(*帕维亚PaviaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "234",
		Title:     "pavia",
		TitleName: "帕维亚",
		TitleCode: "c_pavia",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚历山德里亚Alessandria = BAlessandria亚历山德里亚
	f.亚历山德里亚Alessandria.SetParent(f)
	
	f.博比奥Bobbio = BBobbio博比奥
	f.博比奥Bobbio.SetParent(f)
	
	f.卡斯泰焦Casteggio = BCasteggio卡斯泰焦
	f.卡斯泰焦Casteggio.SetParent(f)
	
	f.科洛尔诺Colorno = BColorno科洛尔诺
	f.科洛尔诺Colorno.SetParent(f)
	
	f.蒙特贝洛Montebello = BMontebello蒙特贝洛
	f.蒙特贝洛Montebello.SetParent(f)
	
	f.帕维亚Pavia = BPavia帕维亚
	f.帕维亚Pavia.SetParent(f)
	
	f.皮亚琴察Piacenza = BPiacenza皮亚琴察
	f.皮亚琴察Piacenza.SetParent(f)
	
	f.沃盖拉Voghera = BVoghera沃盖拉
	f.沃盖拉Voghera.SetParent(f)
	
}
