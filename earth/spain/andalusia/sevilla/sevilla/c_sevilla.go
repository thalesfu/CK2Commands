package sevilla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SevillaCounty interface {
    feud.County
    BAlcaladeguadaira瓜代拉堡() 	feud.Barony
    BCarmona卡莫纳() 	feud.Barony
    BDoshermanas两姊妹镇() 	feud.Barony
    BEcija埃西哈() 	feud.Barony
    BLaalgaba拉阿尔加瓦() 	feud.Barony
    BSevilla塞维利亚() 	feud.Barony
    BSevimoron莫龙() 	feud.Barony
    BUtrera乌特雷拉() 	feud.Barony
}

type 塞维利亚SevillaCounty struct {
	feud.BaseCounty
	瓜代拉堡Alcaladeguadaira 	feud.Barony
	卡莫纳Carmona 	feud.Barony
	两姊妹镇Doshermanas 	feud.Barony
	埃西哈Ecija 	feud.Barony
	拉阿尔加瓦Laalgaba 	feud.Barony
	塞维利亚Sevilla 	feud.Barony
	莫龙Sevimoron 	feud.Barony
	乌特雷拉Utrera 	feud.Barony
}

func (c *塞维利亚SevillaCounty) BAlcaladeguadaira瓜代拉堡() feud.Barony {
	return c.瓜代拉堡Alcaladeguadaira
}
    
func (c *塞维利亚SevillaCounty) BCarmona卡莫纳() feud.Barony {
	return c.卡莫纳Carmona
}
    
func (c *塞维利亚SevillaCounty) BDoshermanas两姊妹镇() feud.Barony {
	return c.两姊妹镇Doshermanas
}
    
func (c *塞维利亚SevillaCounty) BEcija埃西哈() feud.Barony {
	return c.埃西哈Ecija
}
    
func (c *塞维利亚SevillaCounty) BLaalgaba拉阿尔加瓦() feud.Barony {
	return c.拉阿尔加瓦Laalgaba
}
    
func (c *塞维利亚SevillaCounty) BSevilla塞维利亚() feud.Barony {
	return c.塞维利亚Sevilla
}
    
func (c *塞维利亚SevillaCounty) BSevimoron莫龙() feud.Barony {
	return c.莫龙Sevimoron
}
    
func (c *塞维利亚SevillaCounty) BUtrera乌特雷拉() feud.Barony {
	return c.乌特雷拉Utrera
}
    
var CSevilla塞维利亚 SevillaCounty = &塞维利亚SevillaCounty{}

func init() {
	f := CSevilla塞维利亚.(*塞维利亚SevillaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "182",
		Title:     "sevilla",
		TitleName: "塞维利亚",
		TitleCode: "c_sevilla",
		Baronies:  map[string]feud.Barony{},
	}

	f.瓜代拉堡Alcaladeguadaira = BAlcaladeguadaira瓜代拉堡
	f.瓜代拉堡Alcaladeguadaira.SetParent(f)
	
	f.卡莫纳Carmona = BCarmona卡莫纳
	f.卡莫纳Carmona.SetParent(f)
	
	f.两姊妹镇Doshermanas = BDoshermanas两姊妹镇
	f.两姊妹镇Doshermanas.SetParent(f)
	
	f.埃西哈Ecija = BEcija埃西哈
	f.埃西哈Ecija.SetParent(f)
	
	f.拉阿尔加瓦Laalgaba = BLaalgaba拉阿尔加瓦
	f.拉阿尔加瓦Laalgaba.SetParent(f)
	
	f.塞维利亚Sevilla = BSevilla塞维利亚
	f.塞维利亚Sevilla.SetParent(f)
	
	f.莫龙Sevimoron = BSevimoron莫龙
	f.莫龙Sevimoron.SetParent(f)
	
	f.乌特雷拉Utrera = BUtrera乌特雷拉
	f.乌特雷拉Utrera.SetParent(f)
	
}
