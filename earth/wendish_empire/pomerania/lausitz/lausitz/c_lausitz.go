package lausitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LausitzCounty interface {
    feud.County
    BCottbus科特布斯() 	feud.Barony
    BDobrilugk多伯卢格() 	feud.Barony
    BForst福斯特() 	feud.Barony
    BLebus莱布斯() 	feud.Barony
    BLebusa莱布萨() 	feud.Barony
    BLuckau卢考() 	feud.Barony
    BRothenburg罗滕堡() 	feud.Barony
    BZittau齐陶() 	feud.Barony
}

type 劳西茨LausitzCounty struct {
	feud.BaseCounty
	科特布斯Cottbus 	feud.Barony
	多伯卢格Dobrilugk 	feud.Barony
	福斯特Forst 	feud.Barony
	莱布斯Lebus 	feud.Barony
	莱布萨Lebusa 	feud.Barony
	卢考Luckau 	feud.Barony
	罗滕堡Rothenburg 	feud.Barony
	齐陶Zittau 	feud.Barony
}

func (c *劳西茨LausitzCounty) BCottbus科特布斯() feud.Barony {
	return c.科特布斯Cottbus
}
    
func (c *劳西茨LausitzCounty) BDobrilugk多伯卢格() feud.Barony {
	return c.多伯卢格Dobrilugk
}
    
func (c *劳西茨LausitzCounty) BForst福斯特() feud.Barony {
	return c.福斯特Forst
}
    
func (c *劳西茨LausitzCounty) BLebus莱布斯() feud.Barony {
	return c.莱布斯Lebus
}
    
func (c *劳西茨LausitzCounty) BLebusa莱布萨() feud.Barony {
	return c.莱布萨Lebusa
}
    
func (c *劳西茨LausitzCounty) BLuckau卢考() feud.Barony {
	return c.卢考Luckau
}
    
func (c *劳西茨LausitzCounty) BRothenburg罗滕堡() feud.Barony {
	return c.罗滕堡Rothenburg
}
    
func (c *劳西茨LausitzCounty) BZittau齐陶() feud.Barony {
	return c.齐陶Zittau
}
    
var CLausitz劳西茨 LausitzCounty = &劳西茨LausitzCounty{}

func init() {
	f := CLausitz劳西茨.(*劳西茨LausitzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "364",
		Title:     "lausitz",
		TitleName: "劳西茨",
		TitleCode: "c_lausitz",
		Baronies:  map[string]feud.Barony{},
	}

	f.科特布斯Cottbus = BCottbus科特布斯
	f.科特布斯Cottbus.SetParent(f)
	
	f.多伯卢格Dobrilugk = BDobrilugk多伯卢格
	f.多伯卢格Dobrilugk.SetParent(f)
	
	f.福斯特Forst = BForst福斯特
	f.福斯特Forst.SetParent(f)
	
	f.莱布斯Lebus = BLebus莱布斯
	f.莱布斯Lebus.SetParent(f)
	
	f.莱布萨Lebusa = BLebusa莱布萨
	f.莱布萨Lebusa.SetParent(f)
	
	f.卢考Luckau = BLuckau卢考
	f.卢考Luckau.SetParent(f)
	
	f.罗滕堡Rothenburg = BRothenburg罗滕堡
	f.罗滕堡Rothenburg.SetParent(f)
	
	f.齐陶Zittau = BZittau齐陶
	f.齐陶Zittau.SetParent(f)
	
}
