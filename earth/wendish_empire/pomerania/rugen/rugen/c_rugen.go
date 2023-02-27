package rugen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RugenCounty interface {
    feud.County
    BArkona阿科纳() 	feud.Barony
    BBarth巴尔特() 	feud.Barony
    BCharenza卡伦察() 	feud.Barony
    BHiddensee希登塞() 	feud.Barony
    BPutbus普特布斯() 	feud.Barony
    BRalswiek拉尔斯维克() 	feud.Barony
    BRugard卢加德() 	feud.Barony
    BTribuzin特里布钦() 	feud.Barony
}

type 吕根RugenCounty struct {
	feud.BaseCounty
	阿科纳Arkona 	feud.Barony
	巴尔特Barth 	feud.Barony
	卡伦察Charenza 	feud.Barony
	希登塞Hiddensee 	feud.Barony
	普特布斯Putbus 	feud.Barony
	拉尔斯维克Ralswiek 	feud.Barony
	卢加德Rugard 	feud.Barony
	特里布钦Tribuzin 	feud.Barony
}

func (c *吕根RugenCounty) BArkona阿科纳() feud.Barony {
	return c.阿科纳Arkona
}
    
func (c *吕根RugenCounty) BBarth巴尔特() feud.Barony {
	return c.巴尔特Barth
}
    
func (c *吕根RugenCounty) BCharenza卡伦察() feud.Barony {
	return c.卡伦察Charenza
}
    
func (c *吕根RugenCounty) BHiddensee希登塞() feud.Barony {
	return c.希登塞Hiddensee
}
    
func (c *吕根RugenCounty) BPutbus普特布斯() feud.Barony {
	return c.普特布斯Putbus
}
    
func (c *吕根RugenCounty) BRalswiek拉尔斯维克() feud.Barony {
	return c.拉尔斯维克Ralswiek
}
    
func (c *吕根RugenCounty) BRugard卢加德() feud.Barony {
	return c.卢加德Rugard
}
    
func (c *吕根RugenCounty) BTribuzin特里布钦() feud.Barony {
	return c.特里布钦Tribuzin
}
    
var CRugen吕根 RugenCounty = &吕根RugenCounty{}

func init() {
	f := CRugen吕根.(*吕根RugenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "304",
		Title:     "rugen",
		TitleName: "吕根",
		TitleCode: "c_rugen",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿科纳Arkona = BArkona阿科纳
	f.阿科纳Arkona.SetParent(f)
	
	f.巴尔特Barth = BBarth巴尔特
	f.巴尔特Barth.SetParent(f)
	
	f.卡伦察Charenza = BCharenza卡伦察
	f.卡伦察Charenza.SetParent(f)
	
	f.希登塞Hiddensee = BHiddensee希登塞
	f.希登塞Hiddensee.SetParent(f)
	
	f.普特布斯Putbus = BPutbus普特布斯
	f.普特布斯Putbus.SetParent(f)
	
	f.拉尔斯维克Ralswiek = BRalswiek拉尔斯维克
	f.拉尔斯维克Ralswiek.SetParent(f)
	
	f.卢加德Rugard = BRugard卢加德
	f.卢加德Rugard.SetParent(f)
	
	f.特里布钦Tribuzin = BTribuzin特里布钦
	f.特里布钦Tribuzin.SetParent(f)
	
}
