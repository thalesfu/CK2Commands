package dauphine_viennois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Dauphine_viennoisCounty interface {
    feud.County
    BAlbon阿尔邦() 	feud.Barony
    BChartreuse沙特勒斯() 	feud.Barony
    BGrenoble格勒诺布尔() 	feud.Barony
    BMontelimar蒙特利马尔() 	feud.Barony
    BStantoine圣昂图万() 	feud.Barony
    BValence瓦朗斯() 	feud.Barony
    BValreas瓦尔雷阿斯() 	feud.Barony
    BVienne维埃纳() 	feud.Barony
}

type 多菲内Dauphine_viennoisCounty struct {
	feud.BaseCounty
	阿尔邦Albon 	feud.Barony
	沙特勒斯Chartreuse 	feud.Barony
	格勒诺布尔Grenoble 	feud.Barony
	蒙特利马尔Montelimar 	feud.Barony
	圣昂图万Stantoine 	feud.Barony
	瓦朗斯Valence 	feud.Barony
	瓦尔雷阿斯Valreas 	feud.Barony
	维埃纳Vienne 	feud.Barony
}

func (c *多菲内Dauphine_viennoisCounty) BAlbon阿尔邦() feud.Barony {
	return c.阿尔邦Albon
}
    
func (c *多菲内Dauphine_viennoisCounty) BChartreuse沙特勒斯() feud.Barony {
	return c.沙特勒斯Chartreuse
}
    
func (c *多菲内Dauphine_viennoisCounty) BGrenoble格勒诺布尔() feud.Barony {
	return c.格勒诺布尔Grenoble
}
    
func (c *多菲内Dauphine_viennoisCounty) BMontelimar蒙特利马尔() feud.Barony {
	return c.蒙特利马尔Montelimar
}
    
func (c *多菲内Dauphine_viennoisCounty) BStantoine圣昂图万() feud.Barony {
	return c.圣昂图万Stantoine
}
    
func (c *多菲内Dauphine_viennoisCounty) BValence瓦朗斯() feud.Barony {
	return c.瓦朗斯Valence
}
    
func (c *多菲内Dauphine_viennoisCounty) BValreas瓦尔雷阿斯() feud.Barony {
	return c.瓦尔雷阿斯Valreas
}
    
func (c *多菲内Dauphine_viennoisCounty) BVienne维埃纳() feud.Barony {
	return c.维埃纳Vienne
}
    
var CDauphine_viennois多菲内 Dauphine_viennoisCounty = &多菲内Dauphine_viennoisCounty{}

func init() {
	f := CDauphine_viennois多菲内.(*多菲内Dauphine_viennoisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "228",
		Title:     "dauphine_viennois",
		TitleName: "多菲内",
		TitleCode: "c_dauphine_viennois",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔邦Albon = BAlbon阿尔邦
	f.阿尔邦Albon.SetParent(f)
	
	f.沙特勒斯Chartreuse = BChartreuse沙特勒斯
	f.沙特勒斯Chartreuse.SetParent(f)
	
	f.格勒诺布尔Grenoble = BGrenoble格勒诺布尔
	f.格勒诺布尔Grenoble.SetParent(f)
	
	f.蒙特利马尔Montelimar = BMontelimar蒙特利马尔
	f.蒙特利马尔Montelimar.SetParent(f)
	
	f.圣昂图万Stantoine = BStantoine圣昂图万
	f.圣昂图万Stantoine.SetParent(f)
	
	f.瓦朗斯Valence = BValence瓦朗斯
	f.瓦朗斯Valence.SetParent(f)
	
	f.瓦尔雷阿斯Valreas = BValreas瓦尔雷阿斯
	f.瓦尔雷阿斯Valreas.SetParent(f)
	
	f.维埃纳Vienne = BVienne维埃纳
	f.维埃纳Vienne.SetParent(f)
	
}
