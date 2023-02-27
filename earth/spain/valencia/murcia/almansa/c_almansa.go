package almansa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlmansaCounty interface {
    feud.County
    BAlbacete阿尔瓦塞特() 	feud.Barony
    BAlcaladeljucar胡卡尔堡() 	feud.Barony
    BAlmansa阿尔曼萨() 	feud.Barony
    BCaudete考德特() 	feud.Barony
    BHellin埃林() 	feud.Barony
    BPozocanada波索卡尼亚达() 	feud.Barony
    BTobarra托瓦拉() 	feud.Barony
    BVillarrobledo比利亚罗夫莱多() 	feud.Barony
}

type 阿尔曼萨AlmansaCounty struct {
	feud.BaseCounty
	阿尔瓦塞特Albacete 	feud.Barony
	胡卡尔堡Alcaladeljucar 	feud.Barony
	阿尔曼萨Almansa 	feud.Barony
	考德特Caudete 	feud.Barony
	埃林Hellin 	feud.Barony
	波索卡尼亚达Pozocanada 	feud.Barony
	托瓦拉Tobarra 	feud.Barony
	比利亚罗夫莱多Villarrobledo 	feud.Barony
}

func (c *阿尔曼萨AlmansaCounty) BAlbacete阿尔瓦塞特() feud.Barony {
	return c.阿尔瓦塞特Albacete
}
    
func (c *阿尔曼萨AlmansaCounty) BAlcaladeljucar胡卡尔堡() feud.Barony {
	return c.胡卡尔堡Alcaladeljucar
}
    
func (c *阿尔曼萨AlmansaCounty) BAlmansa阿尔曼萨() feud.Barony {
	return c.阿尔曼萨Almansa
}
    
func (c *阿尔曼萨AlmansaCounty) BCaudete考德特() feud.Barony {
	return c.考德特Caudete
}
    
func (c *阿尔曼萨AlmansaCounty) BHellin埃林() feud.Barony {
	return c.埃林Hellin
}
    
func (c *阿尔曼萨AlmansaCounty) BPozocanada波索卡尼亚达() feud.Barony {
	return c.波索卡尼亚达Pozocanada
}
    
func (c *阿尔曼萨AlmansaCounty) BTobarra托瓦拉() feud.Barony {
	return c.托瓦拉Tobarra
}
    
func (c *阿尔曼萨AlmansaCounty) BVillarrobledo比利亚罗夫莱多() feud.Barony {
	return c.比利亚罗夫莱多Villarrobledo
}
    
var CAlmansa阿尔曼萨 AlmansaCounty = &阿尔曼萨AlmansaCounty{}

func init() {
	f := CAlmansa阿尔曼萨.(*阿尔曼萨AlmansaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "179",
		Title:     "almansa",
		TitleName: "阿尔曼萨",
		TitleCode: "c_almansa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔瓦塞特Albacete = BAlbacete阿尔瓦塞特
	f.阿尔瓦塞特Albacete.SetParent(f)
	
	f.胡卡尔堡Alcaladeljucar = BAlcaladeljucar胡卡尔堡
	f.胡卡尔堡Alcaladeljucar.SetParent(f)
	
	f.阿尔曼萨Almansa = BAlmansa阿尔曼萨
	f.阿尔曼萨Almansa.SetParent(f)
	
	f.考德特Caudete = BCaudete考德特
	f.考德特Caudete.SetParent(f)
	
	f.埃林Hellin = BHellin埃林
	f.埃林Hellin.SetParent(f)
	
	f.波索卡尼亚达Pozocanada = BPozocanada波索卡尼亚达
	f.波索卡尼亚达Pozocanada.SetParent(f)
	
	f.托瓦拉Tobarra = BTobarra托瓦拉
	f.托瓦拉Tobarra.SetParent(f)
	
	f.比利亚罗夫莱多Villarrobledo = BVillarrobledo比利亚罗夫莱多
	f.比利亚罗夫莱多Villarrobledo.SetParent(f)
	
}
