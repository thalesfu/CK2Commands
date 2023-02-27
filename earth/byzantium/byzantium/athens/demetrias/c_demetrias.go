package demetrias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DemetriasCounty interface {
    feud.County
    BBoudonitza波多尼察() 	feud.Barony
    BDemetrias德米特里阿斯() 	feud.Barony
    BGravia格拉维亚() 	feud.Barony
    BLebadea莱瓦贾() 	feud.Barony
    BNeopatras尼奥帕特拉斯() 	feud.Barony
    BPharsalus法尔萨卢斯() 	feud.Barony
    BRavennika拉文尼卡() 	feud.Barony
}

type 德米特里阿斯DemetriasCounty struct {
	feud.BaseCounty
	波多尼察Boudonitza 	feud.Barony
	德米特里阿斯Demetrias 	feud.Barony
	格拉维亚Gravia 	feud.Barony
	莱瓦贾Lebadea 	feud.Barony
	尼奥帕特拉斯Neopatras 	feud.Barony
	法尔萨卢斯Pharsalus 	feud.Barony
	拉文尼卡Ravennika 	feud.Barony
}

func (c *德米特里阿斯DemetriasCounty) BBoudonitza波多尼察() feud.Barony {
	return c.波多尼察Boudonitza
}
    
func (c *德米特里阿斯DemetriasCounty) BDemetrias德米特里阿斯() feud.Barony {
	return c.德米特里阿斯Demetrias
}
    
func (c *德米特里阿斯DemetriasCounty) BGravia格拉维亚() feud.Barony {
	return c.格拉维亚Gravia
}
    
func (c *德米特里阿斯DemetriasCounty) BLebadea莱瓦贾() feud.Barony {
	return c.莱瓦贾Lebadea
}
    
func (c *德米特里阿斯DemetriasCounty) BNeopatras尼奥帕特拉斯() feud.Barony {
	return c.尼奥帕特拉斯Neopatras
}
    
func (c *德米特里阿斯DemetriasCounty) BPharsalus法尔萨卢斯() feud.Barony {
	return c.法尔萨卢斯Pharsalus
}
    
func (c *德米特里阿斯DemetriasCounty) BRavennika拉文尼卡() feud.Barony {
	return c.拉文尼卡Ravennika
}
    
var CDemetrias德米特里阿斯 DemetriasCounty = &德米特里阿斯DemetriasCounty{}

func init() {
	f := CDemetrias德米特里阿斯.(*德米特里阿斯DemetriasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "488",
		Title:     "demetrias",
		TitleName: "德米特里阿斯",
		TitleCode: "c_demetrias",
		Baronies:  map[string]feud.Barony{},
	}

	f.波多尼察Boudonitza = BBoudonitza波多尼察
	f.波多尼察Boudonitza.SetParent(f)
	
	f.德米特里阿斯Demetrias = BDemetrias德米特里阿斯
	f.德米特里阿斯Demetrias.SetParent(f)
	
	f.格拉维亚Gravia = BGravia格拉维亚
	f.格拉维亚Gravia.SetParent(f)
	
	f.莱瓦贾Lebadea = BLebadea莱瓦贾
	f.莱瓦贾Lebadea.SetParent(f)
	
	f.尼奥帕特拉斯Neopatras = BNeopatras尼奥帕特拉斯
	f.尼奥帕特拉斯Neopatras.SetParent(f)
	
	f.法尔萨卢斯Pharsalus = BPharsalus法尔萨卢斯
	f.法尔萨卢斯Pharsalus.SetParent(f)
	
	f.拉文尼卡Ravennika = BRavennika拉文尼卡
	f.拉文尼卡Ravennika.SetParent(f)
	
}
