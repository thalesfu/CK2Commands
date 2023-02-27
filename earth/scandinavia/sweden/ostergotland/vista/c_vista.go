package vista

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VistaCounty interface {
    feud.County
    BAlvastra阿尔瓦斯特拉() 	feud.Barony
    BGranna格伦纳() 	feud.Barony
    BJonkoping延雪平() 	feud.Barony
    BNasborg内斯堡() 	feud.Barony
    BOdeshog厄德斯赫格() 	feud.Barony
    BRumlaborg鲁姆拉堡() 	feud.Barony
    BVireda维雷达() 	feud.Barony
}

type 维斯塔VistaCounty struct {
	feud.BaseCounty
	阿尔瓦斯特拉Alvastra 	feud.Barony
	格伦纳Granna 	feud.Barony
	延雪平Jonkoping 	feud.Barony
	内斯堡Nasborg 	feud.Barony
	厄德斯赫格Odeshog 	feud.Barony
	鲁姆拉堡Rumlaborg 	feud.Barony
	维雷达Vireda 	feud.Barony
}

func (c *维斯塔VistaCounty) BAlvastra阿尔瓦斯特拉() feud.Barony {
	return c.阿尔瓦斯特拉Alvastra
}
    
func (c *维斯塔VistaCounty) BGranna格伦纳() feud.Barony {
	return c.格伦纳Granna
}
    
func (c *维斯塔VistaCounty) BJonkoping延雪平() feud.Barony {
	return c.延雪平Jonkoping
}
    
func (c *维斯塔VistaCounty) BNasborg内斯堡() feud.Barony {
	return c.内斯堡Nasborg
}
    
func (c *维斯塔VistaCounty) BOdeshog厄德斯赫格() feud.Barony {
	return c.厄德斯赫格Odeshog
}
    
func (c *维斯塔VistaCounty) BRumlaborg鲁姆拉堡() feud.Barony {
	return c.鲁姆拉堡Rumlaborg
}
    
func (c *维斯塔VistaCounty) BVireda维雷达() feud.Barony {
	return c.维雷达Vireda
}
    
var CVista维斯塔 VistaCounty = &维斯塔VistaCounty{}

func init() {
	f := CVista维斯塔.(*维斯塔VistaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1701",
		Title:     "vista",
		TitleName: "维斯塔",
		TitleCode: "c_vista",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔瓦斯特拉Alvastra = BAlvastra阿尔瓦斯特拉
	f.阿尔瓦斯特拉Alvastra.SetParent(f)
	
	f.格伦纳Granna = BGranna格伦纳
	f.格伦纳Granna.SetParent(f)
	
	f.延雪平Jonkoping = BJonkoping延雪平
	f.延雪平Jonkoping.SetParent(f)
	
	f.内斯堡Nasborg = BNasborg内斯堡
	f.内斯堡Nasborg.SetParent(f)
	
	f.厄德斯赫格Odeshog = BOdeshog厄德斯赫格
	f.厄德斯赫格Odeshog.SetParent(f)
	
	f.鲁姆拉堡Rumlaborg = BRumlaborg鲁姆拉堡
	f.鲁姆拉堡Rumlaborg.SetParent(f)
	
	f.维雷达Vireda = BVireda维雷达
	f.维雷达Vireda.SetParent(f)
	
}
