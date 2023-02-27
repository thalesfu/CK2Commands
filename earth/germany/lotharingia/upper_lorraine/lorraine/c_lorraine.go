package lorraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LorraineCounty interface {
    feud.County
    BCharpeigne沙尔佩涅() 	feud.Barony
    BEpinal埃皮纳勒() 	feud.Barony
    BLuneville吕内维尔() 	feud.Barony
    BMortagnedevosges莫尔塔涅() 	feud.Barony
    BNancy南锡() 	feud.Barony
    BSaintavoid圣阿沃尔德() 	feud.Barony
    BSarrebourg萨尔堡() 	feud.Barony
    BToul图勒() 	feud.Barony
}

type 洛林LorraineCounty struct {
	feud.BaseCounty
	沙尔佩涅Charpeigne 	feud.Barony
	埃皮纳勒Epinal 	feud.Barony
	吕内维尔Luneville 	feud.Barony
	莫尔塔涅Mortagnedevosges 	feud.Barony
	南锡Nancy 	feud.Barony
	圣阿沃尔德Saintavoid 	feud.Barony
	萨尔堡Sarrebourg 	feud.Barony
	图勒Toul 	feud.Barony
}

func (c *洛林LorraineCounty) BCharpeigne沙尔佩涅() feud.Barony {
	return c.沙尔佩涅Charpeigne
}
    
func (c *洛林LorraineCounty) BEpinal埃皮纳勒() feud.Barony {
	return c.埃皮纳勒Epinal
}
    
func (c *洛林LorraineCounty) BLuneville吕内维尔() feud.Barony {
	return c.吕内维尔Luneville
}
    
func (c *洛林LorraineCounty) BMortagnedevosges莫尔塔涅() feud.Barony {
	return c.莫尔塔涅Mortagnedevosges
}
    
func (c *洛林LorraineCounty) BNancy南锡() feud.Barony {
	return c.南锡Nancy
}
    
func (c *洛林LorraineCounty) BSaintavoid圣阿沃尔德() feud.Barony {
	return c.圣阿沃尔德Saintavoid
}
    
func (c *洛林LorraineCounty) BSarrebourg萨尔堡() feud.Barony {
	return c.萨尔堡Sarrebourg
}
    
func (c *洛林LorraineCounty) BToul图勒() feud.Barony {
	return c.图勒Toul
}
    
var CLorraine洛林 LorraineCounty = &洛林LorraineCounty{}

func init() {
	f := CLorraine洛林.(*洛林LorraineCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "127",
		Title:     "lorraine",
		TitleName: "洛林",
		TitleCode: "c_lorraine",
		Baronies:  map[string]feud.Barony{},
	}

	f.沙尔佩涅Charpeigne = BCharpeigne沙尔佩涅
	f.沙尔佩涅Charpeigne.SetParent(f)
	
	f.埃皮纳勒Epinal = BEpinal埃皮纳勒
	f.埃皮纳勒Epinal.SetParent(f)
	
	f.吕内维尔Luneville = BLuneville吕内维尔
	f.吕内维尔Luneville.SetParent(f)
	
	f.莫尔塔涅Mortagnedevosges = BMortagnedevosges莫尔塔涅
	f.莫尔塔涅Mortagnedevosges.SetParent(f)
	
	f.南锡Nancy = BNancy南锡
	f.南锡Nancy.SetParent(f)
	
	f.圣阿沃尔德Saintavoid = BSaintavoid圣阿沃尔德
	f.圣阿沃尔德Saintavoid.SetParent(f)
	
	f.萨尔堡Sarrebourg = BSarrebourg萨尔堡
	f.萨尔堡Sarrebourg.SetParent(f)
	
	f.图勒Toul = BToul图勒
	f.图勒Toul.SetParent(f)
	
}
