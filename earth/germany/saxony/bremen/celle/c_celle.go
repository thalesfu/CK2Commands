package celle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CelleCounty interface {
    feud.County
    BBremen不来梅() 	feud.Barony
    BCelle策勒() 	feud.Barony
    BHannover汉诺威() 	feud.Barony
    BHerford黑尔福德() 	feud.Barony
    BHermannsburg黑尔曼斯堡() 	feud.Barony
    BNienburg宁堡() 	feud.Barony
    BRavensberg拉芬斯堡() 	feud.Barony
    BWedemark韦德马克() 	feud.Barony
    BWittingen维廷根() 	feud.Barony
}

type 不来梅CelleCounty struct {
	feud.BaseCounty
	不来梅Bremen 	feud.Barony
	策勒Celle 	feud.Barony
	汉诺威Hannover 	feud.Barony
	黑尔福德Herford 	feud.Barony
	黑尔曼斯堡Hermannsburg 	feud.Barony
	宁堡Nienburg 	feud.Barony
	拉芬斯堡Ravensberg 	feud.Barony
	韦德马克Wedemark 	feud.Barony
	维廷根Wittingen 	feud.Barony
}

func (c *不来梅CelleCounty) BBremen不来梅() feud.Barony {
	return c.不来梅Bremen
}
    
func (c *不来梅CelleCounty) BCelle策勒() feud.Barony {
	return c.策勒Celle
}
    
func (c *不来梅CelleCounty) BHannover汉诺威() feud.Barony {
	return c.汉诺威Hannover
}
    
func (c *不来梅CelleCounty) BHerford黑尔福德() feud.Barony {
	return c.黑尔福德Herford
}
    
func (c *不来梅CelleCounty) BHermannsburg黑尔曼斯堡() feud.Barony {
	return c.黑尔曼斯堡Hermannsburg
}
    
func (c *不来梅CelleCounty) BNienburg宁堡() feud.Barony {
	return c.宁堡Nienburg
}
    
func (c *不来梅CelleCounty) BRavensberg拉芬斯堡() feud.Barony {
	return c.拉芬斯堡Ravensberg
}
    
func (c *不来梅CelleCounty) BWedemark韦德马克() feud.Barony {
	return c.韦德马克Wedemark
}
    
func (c *不来梅CelleCounty) BWittingen维廷根() feud.Barony {
	return c.维廷根Wittingen
}
    
var CCelle不来梅 CelleCounty = &不来梅CelleCounty{}

func init() {
	f := CCelle不来梅.(*不来梅CelleCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "259",
		Title:     "celle",
		TitleName: "不来梅",
		TitleCode: "c_celle",
		Baronies:  map[string]feud.Barony{},
	}

	f.不来梅Bremen = BBremen不来梅
	f.不来梅Bremen.SetParent(f)
	
	f.策勒Celle = BCelle策勒
	f.策勒Celle.SetParent(f)
	
	f.汉诺威Hannover = BHannover汉诺威
	f.汉诺威Hannover.SetParent(f)
	
	f.黑尔福德Herford = BHerford黑尔福德
	f.黑尔福德Herford.SetParent(f)
	
	f.黑尔曼斯堡Hermannsburg = BHermannsburg黑尔曼斯堡
	f.黑尔曼斯堡Hermannsburg.SetParent(f)
	
	f.宁堡Nienburg = BNienburg宁堡
	f.宁堡Nienburg.SetParent(f)
	
	f.拉芬斯堡Ravensberg = BRavensberg拉芬斯堡
	f.拉芬斯堡Ravensberg.SetParent(f)
	
	f.韦德马克Wedemark = BWedemark韦德马克
	f.韦德马克Wedemark.SetParent(f)
	
	f.维廷根Wittingen = BWittingen维廷根
	f.维廷根Wittingen.SetParent(f)
	
}
