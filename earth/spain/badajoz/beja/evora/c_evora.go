package evora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EvoraCounty interface {
    feud.County
    BAvis阿维什() 	feud.Barony
    BCastelodevide维迪堡() 	feud.Barony
    BCrato克拉图() 	feud.Barony
    BEvora埃武拉() 	feud.Barony
    BMarvao马尔旺() 	feud.Barony
    BMonforte蒙福蒂() 	feud.Barony
    BOuguela欧格拉() 	feud.Barony
    BPortalegre波塔莱格雷() 	feud.Barony
}

type 埃武拉EvoraCounty struct {
	feud.BaseCounty
	阿维什Avis 	feud.Barony
	维迪堡Castelodevide 	feud.Barony
	克拉图Crato 	feud.Barony
	埃武拉Evora 	feud.Barony
	马尔旺Marvao 	feud.Barony
	蒙福蒂Monforte 	feud.Barony
	欧格拉Ouguela 	feud.Barony
	波塔莱格雷Portalegre 	feud.Barony
}

func (c *埃武拉EvoraCounty) BAvis阿维什() feud.Barony {
	return c.阿维什Avis
}
    
func (c *埃武拉EvoraCounty) BCastelodevide维迪堡() feud.Barony {
	return c.维迪堡Castelodevide
}
    
func (c *埃武拉EvoraCounty) BCrato克拉图() feud.Barony {
	return c.克拉图Crato
}
    
func (c *埃武拉EvoraCounty) BEvora埃武拉() feud.Barony {
	return c.埃武拉Evora
}
    
func (c *埃武拉EvoraCounty) BMarvao马尔旺() feud.Barony {
	return c.马尔旺Marvao
}
    
func (c *埃武拉EvoraCounty) BMonforte蒙福蒂() feud.Barony {
	return c.蒙福蒂Monforte
}
    
func (c *埃武拉EvoraCounty) BOuguela欧格拉() feud.Barony {
	return c.欧格拉Ouguela
}
    
func (c *埃武拉EvoraCounty) BPortalegre波塔莱格雷() feud.Barony {
	return c.波塔莱格雷Portalegre
}
    
var CEvora埃武拉 EvoraCounty = &埃武拉EvoraCounty{}

func init() {
	f := CEvora埃武拉.(*埃武拉EvoraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "186",
		Title:     "evora",
		TitleName: "埃武拉",
		TitleCode: "c_evora",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿维什Avis = BAvis阿维什
	f.阿维什Avis.SetParent(f)
	
	f.维迪堡Castelodevide = BCastelodevide维迪堡
	f.维迪堡Castelodevide.SetParent(f)
	
	f.克拉图Crato = BCrato克拉图
	f.克拉图Crato.SetParent(f)
	
	f.埃武拉Evora = BEvora埃武拉
	f.埃武拉Evora.SetParent(f)
	
	f.马尔旺Marvao = BMarvao马尔旺
	f.马尔旺Marvao.SetParent(f)
	
	f.蒙福蒂Monforte = BMonforte蒙福蒂
	f.蒙福蒂Monforte.SetParent(f)
	
	f.欧格拉Ouguela = BOuguela欧格拉
	f.欧格拉Ouguela.SetParent(f)
	
	f.波塔莱格雷Portalegre = BPortalegre波塔莱格雷
	f.波塔莱格雷Portalegre.SetParent(f)
	
}
