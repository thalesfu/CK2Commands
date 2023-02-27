package syrmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SyrmiaCounty interface {
    feud.County
    BBassiana巴西亚纳() 	feud.Barony
    BBononia博诺尼亚() 	feud.Barony
    BBurgenae布尔格奈() 	feud.Barony
    BRittium里蒂乌姆() 	feud.Barony
    BZemlen泽姆伦() 	feud.Barony
    BZenthdemeter圣戴迈泰尔() 	feud.Barony
    BZvecan_syrmia兹韦钱() 	feud.Barony
}

type 叙尔米亚SyrmiaCounty struct {
	feud.BaseCounty
	巴西亚纳Bassiana 	feud.Barony
	博诺尼亚Bononia 	feud.Barony
	布尔格奈Burgenae 	feud.Barony
	里蒂乌姆Rittium 	feud.Barony
	泽姆伦Zemlen 	feud.Barony
	圣戴迈泰尔Zenthdemeter 	feud.Barony
	兹韦钱Zvecan_syrmia 	feud.Barony
}

func (c *叙尔米亚SyrmiaCounty) BBassiana巴西亚纳() feud.Barony {
	return c.巴西亚纳Bassiana
}
    
func (c *叙尔米亚SyrmiaCounty) BBononia博诺尼亚() feud.Barony {
	return c.博诺尼亚Bononia
}
    
func (c *叙尔米亚SyrmiaCounty) BBurgenae布尔格奈() feud.Barony {
	return c.布尔格奈Burgenae
}
    
func (c *叙尔米亚SyrmiaCounty) BRittium里蒂乌姆() feud.Barony {
	return c.里蒂乌姆Rittium
}
    
func (c *叙尔米亚SyrmiaCounty) BZemlen泽姆伦() feud.Barony {
	return c.泽姆伦Zemlen
}
    
func (c *叙尔米亚SyrmiaCounty) BZenthdemeter圣戴迈泰尔() feud.Barony {
	return c.圣戴迈泰尔Zenthdemeter
}
    
func (c *叙尔米亚SyrmiaCounty) BZvecan_syrmia兹韦钱() feud.Barony {
	return c.兹韦钱Zvecan_syrmia
}
    
var CSyrmia叙尔米亚 SyrmiaCounty = &叙尔米亚SyrmiaCounty{}

func init() {
	f := CSyrmia叙尔米亚.(*叙尔米亚SyrmiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1968",
		Title:     "syrmia",
		TitleName: "叙尔米亚",
		TitleCode: "c_syrmia",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴西亚纳Bassiana = BBassiana巴西亚纳
	f.巴西亚纳Bassiana.SetParent(f)
	
	f.博诺尼亚Bononia = BBononia博诺尼亚
	f.博诺尼亚Bononia.SetParent(f)
	
	f.布尔格奈Burgenae = BBurgenae布尔格奈
	f.布尔格奈Burgenae.SetParent(f)
	
	f.里蒂乌姆Rittium = BRittium里蒂乌姆
	f.里蒂乌姆Rittium.SetParent(f)
	
	f.泽姆伦Zemlen = BZemlen泽姆伦
	f.泽姆伦Zemlen.SetParent(f)
	
	f.圣戴迈泰尔Zenthdemeter = BZenthdemeter圣戴迈泰尔
	f.圣戴迈泰尔Zenthdemeter.SetParent(f)
	
	f.兹韦钱Zvecan_syrmia = BZvecan_syrmia兹韦钱
	f.兹韦钱Zvecan_syrmia.SetParent(f)
	
}
