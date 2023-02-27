package french_leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type French_leonCounty interface {
    feud.County
    BBreles布雷莱() 	feud.Barony
    BBrest布雷斯特() 	feud.Barony
    BLanderneau朗代诺() 	feud.Barony
    BMorlaix莫尔莱() 	feud.Barony
    BPlouescat普卢埃斯卡() 	feud.Barony
    BRoscoff罗斯科夫() 	feud.Barony
    BSt_pol_de_leon圣波勒德莱昂() 	feud.Barony
}

type 莱昂French_leonCounty struct {
	feud.BaseCounty
	布雷莱Breles 	feud.Barony
	布雷斯特Brest 	feud.Barony
	朗代诺Landerneau 	feud.Barony
	莫尔莱Morlaix 	feud.Barony
	普卢埃斯卡Plouescat 	feud.Barony
	罗斯科夫Roscoff 	feud.Barony
	圣波勒德莱昂St_pol_de_leon 	feud.Barony
}

func (c *莱昂French_leonCounty) BBreles布雷莱() feud.Barony {
	return c.布雷莱Breles
}
    
func (c *莱昂French_leonCounty) BBrest布雷斯特() feud.Barony {
	return c.布雷斯特Brest
}
    
func (c *莱昂French_leonCounty) BLanderneau朗代诺() feud.Barony {
	return c.朗代诺Landerneau
}
    
func (c *莱昂French_leonCounty) BMorlaix莫尔莱() feud.Barony {
	return c.莫尔莱Morlaix
}
    
func (c *莱昂French_leonCounty) BPlouescat普卢埃斯卡() feud.Barony {
	return c.普卢埃斯卡Plouescat
}
    
func (c *莱昂French_leonCounty) BRoscoff罗斯科夫() feud.Barony {
	return c.罗斯科夫Roscoff
}
    
func (c *莱昂French_leonCounty) BSt_pol_de_leon圣波勒德莱昂() feud.Barony {
	return c.圣波勒德莱昂St_pol_de_leon
}
    
var CFrench_leon莱昂 French_leonCounty = &莱昂French_leonCounty{}

func init() {
	f := CFrench_leon莱昂.(*莱昂French_leonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "103",
		Title:     "french_leon",
		TitleName: "莱昂",
		TitleCode: "c_french_leon",
		Baronies:  map[string]feud.Barony{},
	}

	f.布雷莱Breles = BBreles布雷莱
	f.布雷莱Breles.SetParent(f)
	
	f.布雷斯特Brest = BBrest布雷斯特
	f.布雷斯特Brest.SetParent(f)
	
	f.朗代诺Landerneau = BLanderneau朗代诺
	f.朗代诺Landerneau.SetParent(f)
	
	f.莫尔莱Morlaix = BMorlaix莫尔莱
	f.莫尔莱Morlaix.SetParent(f)
	
	f.普卢埃斯卡Plouescat = BPlouescat普卢埃斯卡
	f.普卢埃斯卡Plouescat.SetParent(f)
	
	f.罗斯科夫Roscoff = BRoscoff罗斯科夫
	f.罗斯科夫Roscoff.SetParent(f)
	
	f.圣波勒德莱昂St_pol_de_leon = BSt_pol_de_leon圣波勒德莱昂
	f.圣波勒德莱昂St_pol_de_leon.SetParent(f)
	
}
