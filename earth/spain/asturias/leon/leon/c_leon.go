package leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeonCounty interface {
    feud.County
    BCistierna锡斯铁尔纳() 	feud.Barony
    BLarobla拉罗夫拉() 	feud.Barony
    BLeon莱昂() 	feud.Barony
    BSahagun萨阿贡() 	feud.Barony
    BSaldana萨尔达尼亚() 	feud.Barony
    BSanpedrodeperix圣佩德罗德佩里克斯() 	feud.Barony
    BValenciadecampos瓦伦西亚德坎波斯() 	feud.Barony
    BVillablino比利亚布利诺() 	feud.Barony
}

type 莱昂LeonCounty struct {
	feud.BaseCounty
	锡斯铁尔纳Cistierna 	feud.Barony
	拉罗夫拉Larobla 	feud.Barony
	莱昂Leon 	feud.Barony
	萨阿贡Sahagun 	feud.Barony
	萨尔达尼亚Saldana 	feud.Barony
	圣佩德罗德佩里克斯Sanpedrodeperix 	feud.Barony
	瓦伦西亚德坎波斯Valenciadecampos 	feud.Barony
	比利亚布利诺Villablino 	feud.Barony
}

func (c *莱昂LeonCounty) BCistierna锡斯铁尔纳() feud.Barony {
	return c.锡斯铁尔纳Cistierna
}
    
func (c *莱昂LeonCounty) BLarobla拉罗夫拉() feud.Barony {
	return c.拉罗夫拉Larobla
}
    
func (c *莱昂LeonCounty) BLeon莱昂() feud.Barony {
	return c.莱昂Leon
}
    
func (c *莱昂LeonCounty) BSahagun萨阿贡() feud.Barony {
	return c.萨阿贡Sahagun
}
    
func (c *莱昂LeonCounty) BSaldana萨尔达尼亚() feud.Barony {
	return c.萨尔达尼亚Saldana
}
    
func (c *莱昂LeonCounty) BSanpedrodeperix圣佩德罗德佩里克斯() feud.Barony {
	return c.圣佩德罗德佩里克斯Sanpedrodeperix
}
    
func (c *莱昂LeonCounty) BValenciadecampos瓦伦西亚德坎波斯() feud.Barony {
	return c.瓦伦西亚德坎波斯Valenciadecampos
}
    
func (c *莱昂LeonCounty) BVillablino比利亚布利诺() feud.Barony {
	return c.比利亚布利诺Villablino
}
    
var CLeon莱昂 LeonCounty = &莱昂LeonCounty{}

func init() {
	f := CLeon莱昂.(*莱昂LeonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "190",
		Title:     "leon",
		TitleName: "莱昂",
		TitleCode: "c_leon",
		Baronies:  map[string]feud.Barony{},
	}

	f.锡斯铁尔纳Cistierna = BCistierna锡斯铁尔纳
	f.锡斯铁尔纳Cistierna.SetParent(f)
	
	f.拉罗夫拉Larobla = BLarobla拉罗夫拉
	f.拉罗夫拉Larobla.SetParent(f)
	
	f.莱昂Leon = BLeon莱昂
	f.莱昂Leon.SetParent(f)
	
	f.萨阿贡Sahagun = BSahagun萨阿贡
	f.萨阿贡Sahagun.SetParent(f)
	
	f.萨尔达尼亚Saldana = BSaldana萨尔达尼亚
	f.萨尔达尼亚Saldana.SetParent(f)
	
	f.圣佩德罗德佩里克斯Sanpedrodeperix = BSanpedrodeperix圣佩德罗德佩里克斯
	f.圣佩德罗德佩里克斯Sanpedrodeperix.SetParent(f)
	
	f.瓦伦西亚德坎波斯Valenciadecampos = BValenciadecampos瓦伦西亚德坎波斯
	f.瓦伦西亚德坎波斯Valenciadecampos.SetParent(f)
	
	f.比利亚布利诺Villablino = BVillablino比利亚布利诺
	f.比利亚布利诺Villablino.SetParent(f)
	
}
