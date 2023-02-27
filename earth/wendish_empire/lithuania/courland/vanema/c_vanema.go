package vanema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VanemaCounty interface {
    feud.County
    BLagzdene拉格兹德内() 	feud.Barony
    BMatkule马特库莱() 	feud.Barony
    BPiltyn皮尔滕() 	feud.Barony
    BRumen鲁梅内() 	feud.Barony
    BTalsen塔尔森() 	feud.Barony
    BVanema瓦内马() 	feud.Barony
    BVentspils文茨皮尔斯() 	feud.Barony
}

type 瓦内马VanemaCounty struct {
	feud.BaseCounty
	拉格兹德内Lagzdene 	feud.Barony
	马特库莱Matkule 	feud.Barony
	皮尔滕Piltyn 	feud.Barony
	鲁梅内Rumen 	feud.Barony
	塔尔森Talsen 	feud.Barony
	瓦内马Vanema 	feud.Barony
	文茨皮尔斯Ventspils 	feud.Barony
}

func (c *瓦内马VanemaCounty) BLagzdene拉格兹德内() feud.Barony {
	return c.拉格兹德内Lagzdene
}
    
func (c *瓦内马VanemaCounty) BMatkule马特库莱() feud.Barony {
	return c.马特库莱Matkule
}
    
func (c *瓦内马VanemaCounty) BPiltyn皮尔滕() feud.Barony {
	return c.皮尔滕Piltyn
}
    
func (c *瓦内马VanemaCounty) BRumen鲁梅内() feud.Barony {
	return c.鲁梅内Rumen
}
    
func (c *瓦内马VanemaCounty) BTalsen塔尔森() feud.Barony {
	return c.塔尔森Talsen
}
    
func (c *瓦内马VanemaCounty) BVanema瓦内马() feud.Barony {
	return c.瓦内马Vanema
}
    
func (c *瓦内马VanemaCounty) BVentspils文茨皮尔斯() feud.Barony {
	return c.文茨皮尔斯Ventspils
}
    
var CVanema瓦内马 VanemaCounty = &瓦内马VanemaCounty{}

func init() {
	f := CVanema瓦内马.(*瓦内马VanemaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1593",
		Title:     "vanema",
		TitleName: "瓦内马",
		TitleCode: "c_vanema",
		Baronies:  map[string]feud.Barony{},
	}

	f.拉格兹德内Lagzdene = BLagzdene拉格兹德内
	f.拉格兹德内Lagzdene.SetParent(f)
	
	f.马特库莱Matkule = BMatkule马特库莱
	f.马特库莱Matkule.SetParent(f)
	
	f.皮尔滕Piltyn = BPiltyn皮尔滕
	f.皮尔滕Piltyn.SetParent(f)
	
	f.鲁梅内Rumen = BRumen鲁梅内
	f.鲁梅内Rumen.SetParent(f)
	
	f.塔尔森Talsen = BTalsen塔尔森
	f.塔尔森Talsen.SetParent(f)
	
	f.瓦内马Vanema = BVanema瓦内马
	f.瓦内马Vanema.SetParent(f)
	
	f.文茨皮尔斯Ventspils = BVentspils文茨皮尔斯
	f.文茨皮尔斯Ventspils.SetParent(f)
	
}
