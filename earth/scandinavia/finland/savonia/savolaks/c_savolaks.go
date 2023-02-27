package savolaks

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SavolaksCounty interface {
    feud.County
    BBrahelinna布拉海林纳() 	feud.Barony
    BMikkeli米凯利() 	feud.Barony
    BOlavinlinna奥拉维堡() 	feud.Barony
    BPuula普拉() 	feud.Barony
    BSavitaipale萨维泰帕莱() 	feud.Barony
    BSavonlinna萨翁林纳() 	feud.Barony
    BSysma叙斯迈() 	feud.Barony
}

type 萨沃SavolaksCounty struct {
	feud.BaseCounty
	布拉海林纳Brahelinna 	feud.Barony
	米凯利Mikkeli 	feud.Barony
	奥拉维堡Olavinlinna 	feud.Barony
	普拉Puula 	feud.Barony
	萨维泰帕莱Savitaipale 	feud.Barony
	萨翁林纳Savonlinna 	feud.Barony
	叙斯迈Sysma 	feud.Barony
}

func (c *萨沃SavolaksCounty) BBrahelinna布拉海林纳() feud.Barony {
	return c.布拉海林纳Brahelinna
}
    
func (c *萨沃SavolaksCounty) BMikkeli米凯利() feud.Barony {
	return c.米凯利Mikkeli
}
    
func (c *萨沃SavolaksCounty) BOlavinlinna奥拉维堡() feud.Barony {
	return c.奥拉维堡Olavinlinna
}
    
func (c *萨沃SavolaksCounty) BPuula普拉() feud.Barony {
	return c.普拉Puula
}
    
func (c *萨沃SavolaksCounty) BSavitaipale萨维泰帕莱() feud.Barony {
	return c.萨维泰帕莱Savitaipale
}
    
func (c *萨沃SavolaksCounty) BSavonlinna萨翁林纳() feud.Barony {
	return c.萨翁林纳Savonlinna
}
    
func (c *萨沃SavolaksCounty) BSysma叙斯迈() feud.Barony {
	return c.叙斯迈Sysma
}
    
var CSavolaks萨沃 SavolaksCounty = &萨沃SavolaksCounty{}

func init() {
	f := CSavolaks萨沃.(*萨沃SavolaksCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "390",
		Title:     "savolaks",
		TitleName: "萨沃",
		TitleCode: "c_savolaks",
		Baronies:  map[string]feud.Barony{},
	}

	f.布拉海林纳Brahelinna = BBrahelinna布拉海林纳
	f.布拉海林纳Brahelinna.SetParent(f)
	
	f.米凯利Mikkeli = BMikkeli米凯利
	f.米凯利Mikkeli.SetParent(f)
	
	f.奥拉维堡Olavinlinna = BOlavinlinna奥拉维堡
	f.奥拉维堡Olavinlinna.SetParent(f)
	
	f.普拉Puula = BPuula普拉
	f.普拉Puula.SetParent(f)
	
	f.萨维泰帕莱Savitaipale = BSavitaipale萨维泰帕莱
	f.萨维泰帕莱Savitaipale.SetParent(f)
	
	f.萨翁林纳Savonlinna = BSavonlinna萨翁林纳
	f.萨翁林纳Savonlinna.SetParent(f)
	
	f.叙斯迈Sysma = BSysma叙斯迈
	f.叙斯迈Sysma.SetParent(f)
	
}
