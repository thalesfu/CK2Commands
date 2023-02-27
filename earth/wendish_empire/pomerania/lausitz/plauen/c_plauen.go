package plauen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PlauenCounty interface {
    feud.County
    BHalle哈雷() 	feud.Barony
    BKnobelsdorf克诺贝尔斯多夫() 	feud.Barony
    BLeipzig莱比锡() 	feud.Barony
    BMerseburg梅泽堡() 	feud.Barony
    BNaumburg瑙姆堡() 	feud.Barony
    BPlauen普劳恩() 	feud.Barony
    BZeitz蔡茨() 	feud.Barony
    BZwickau茨维考() 	feud.Barony
}

type 梅泽堡PlauenCounty struct {
	feud.BaseCounty
	哈雷Halle 	feud.Barony
	克诺贝尔斯多夫Knobelsdorf 	feud.Barony
	莱比锡Leipzig 	feud.Barony
	梅泽堡Merseburg 	feud.Barony
	瑙姆堡Naumburg 	feud.Barony
	普劳恩Plauen 	feud.Barony
	蔡茨Zeitz 	feud.Barony
	茨维考Zwickau 	feud.Barony
}

func (c *梅泽堡PlauenCounty) BHalle哈雷() feud.Barony {
	return c.哈雷Halle
}
    
func (c *梅泽堡PlauenCounty) BKnobelsdorf克诺贝尔斯多夫() feud.Barony {
	return c.克诺贝尔斯多夫Knobelsdorf
}
    
func (c *梅泽堡PlauenCounty) BLeipzig莱比锡() feud.Barony {
	return c.莱比锡Leipzig
}
    
func (c *梅泽堡PlauenCounty) BMerseburg梅泽堡() feud.Barony {
	return c.梅泽堡Merseburg
}
    
func (c *梅泽堡PlauenCounty) BNaumburg瑙姆堡() feud.Barony {
	return c.瑙姆堡Naumburg
}
    
func (c *梅泽堡PlauenCounty) BPlauen普劳恩() feud.Barony {
	return c.普劳恩Plauen
}
    
func (c *梅泽堡PlauenCounty) BZeitz蔡茨() feud.Barony {
	return c.蔡茨Zeitz
}
    
func (c *梅泽堡PlauenCounty) BZwickau茨维考() feud.Barony {
	return c.茨维考Zwickau
}
    
var CPlauen梅泽堡 PlauenCounty = &梅泽堡PlauenCounty{}

func init() {
	f := CPlauen梅泽堡.(*梅泽堡PlauenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "311",
		Title:     "plauen",
		TitleName: "梅泽堡",
		TitleCode: "c_plauen",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈雷Halle = BHalle哈雷
	f.哈雷Halle.SetParent(f)
	
	f.克诺贝尔斯多夫Knobelsdorf = BKnobelsdorf克诺贝尔斯多夫
	f.克诺贝尔斯多夫Knobelsdorf.SetParent(f)
	
	f.莱比锡Leipzig = BLeipzig莱比锡
	f.莱比锡Leipzig.SetParent(f)
	
	f.梅泽堡Merseburg = BMerseburg梅泽堡
	f.梅泽堡Merseburg.SetParent(f)
	
	f.瑙姆堡Naumburg = BNaumburg瑙姆堡
	f.瑙姆堡Naumburg.SetParent(f)
	
	f.普劳恩Plauen = BPlauen普劳恩
	f.普劳恩Plauen.SetParent(f)
	
	f.蔡茨Zeitz = BZeitz蔡茨
	f.蔡茨Zeitz.SetParent(f)
	
	f.茨维考Zwickau = BZwickau茨维考
	f.茨维考Zwickau.SetParent(f)
	
}
