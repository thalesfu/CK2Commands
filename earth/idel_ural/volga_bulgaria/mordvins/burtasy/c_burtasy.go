package burtasy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BurtasyCounty interface {
    feud.County
    BAtkarsk阿特卡尔斯克() 	feud.Barony
    BBalakova巴拉科沃() 	feud.Barony
    BGolykaramysh戈雷卡拉梅什() 	feud.Barony
    BMechetnaya梅切特纳亚() 	feud.Barony
    BPokrovsk波克罗夫斯克() 	feud.Barony
    BRtishchevo勒季谢沃() 	feud.Barony
    BSaratov萨拉托夫() 	feud.Barony
    BSosnovyostrov松岛() 	feud.Barony
}

type 布尔塔瑟BurtasyCounty struct {
	feud.BaseCounty
	阿特卡尔斯克Atkarsk 	feud.Barony
	巴拉科沃Balakova 	feud.Barony
	戈雷卡拉梅什Golykaramysh 	feud.Barony
	梅切特纳亚Mechetnaya 	feud.Barony
	波克罗夫斯克Pokrovsk 	feud.Barony
	勒季谢沃Rtishchevo 	feud.Barony
	萨拉托夫Saratov 	feud.Barony
	松岛Sosnovyostrov 	feud.Barony
}

func (c *布尔塔瑟BurtasyCounty) BAtkarsk阿特卡尔斯克() feud.Barony {
	return c.阿特卡尔斯克Atkarsk
}
    
func (c *布尔塔瑟BurtasyCounty) BBalakova巴拉科沃() feud.Barony {
	return c.巴拉科沃Balakova
}
    
func (c *布尔塔瑟BurtasyCounty) BGolykaramysh戈雷卡拉梅什() feud.Barony {
	return c.戈雷卡拉梅什Golykaramysh
}
    
func (c *布尔塔瑟BurtasyCounty) BMechetnaya梅切特纳亚() feud.Barony {
	return c.梅切特纳亚Mechetnaya
}
    
func (c *布尔塔瑟BurtasyCounty) BPokrovsk波克罗夫斯克() feud.Barony {
	return c.波克罗夫斯克Pokrovsk
}
    
func (c *布尔塔瑟BurtasyCounty) BRtishchevo勒季谢沃() feud.Barony {
	return c.勒季谢沃Rtishchevo
}
    
func (c *布尔塔瑟BurtasyCounty) BSaratov萨拉托夫() feud.Barony {
	return c.萨拉托夫Saratov
}
    
func (c *布尔塔瑟BurtasyCounty) BSosnovyostrov松岛() feud.Barony {
	return c.松岛Sosnovyostrov
}
    
var CBurtasy布尔塔瑟 BurtasyCounty = &布尔塔瑟BurtasyCounty{}

func init() {
	f := CBurtasy布尔塔瑟.(*布尔塔瑟BurtasyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "592",
		Title:     "burtasy",
		TitleName: "布尔塔瑟",
		TitleCode: "c_burtasy",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿特卡尔斯克Atkarsk = BAtkarsk阿特卡尔斯克
	f.阿特卡尔斯克Atkarsk.SetParent(f)
	
	f.巴拉科沃Balakova = BBalakova巴拉科沃
	f.巴拉科沃Balakova.SetParent(f)
	
	f.戈雷卡拉梅什Golykaramysh = BGolykaramysh戈雷卡拉梅什
	f.戈雷卡拉梅什Golykaramysh.SetParent(f)
	
	f.梅切特纳亚Mechetnaya = BMechetnaya梅切特纳亚
	f.梅切特纳亚Mechetnaya.SetParent(f)
	
	f.波克罗夫斯克Pokrovsk = BPokrovsk波克罗夫斯克
	f.波克罗夫斯克Pokrovsk.SetParent(f)
	
	f.勒季谢沃Rtishchevo = BRtishchevo勒季谢沃
	f.勒季谢沃Rtishchevo.SetParent(f)
	
	f.萨拉托夫Saratov = BSaratov萨拉托夫
	f.萨拉托夫Saratov.SetParent(f)
	
	f.松岛Sosnovyostrov = BSosnovyostrov松岛
	f.松岛Sosnovyostrov.SetParent(f)
	
}
