package tortosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TortosaCounty interface {
    feud.County
    BAlkhawabi哈瓦比() 	feud.Barony
    BBalemia巴勒米亚() 	feud.Barony
    BJabala杰卜莱() 	feud.Barony
    BMaraclea马拉克利() 	feud.Barony
    BMargat马加特() 	feud.Barony
    BRuad鲁阿德() 	feud.Barony
    BTortosa托尔托萨() 	feud.Barony
    BValania瓦拉尼亚() 	feud.Barony
}

type 托尔托萨TortosaCounty struct {
	feud.BaseCounty
	哈瓦比Alkhawabi 	feud.Barony
	巴勒米亚Balemia 	feud.Barony
	杰卜莱Jabala 	feud.Barony
	马拉克利Maraclea 	feud.Barony
	马加特Margat 	feud.Barony
	鲁阿德Ruad 	feud.Barony
	托尔托萨Tortosa 	feud.Barony
	瓦拉尼亚Valania 	feud.Barony
}

func (c *托尔托萨TortosaCounty) BAlkhawabi哈瓦比() feud.Barony {
	return c.哈瓦比Alkhawabi
}
    
func (c *托尔托萨TortosaCounty) BBalemia巴勒米亚() feud.Barony {
	return c.巴勒米亚Balemia
}
    
func (c *托尔托萨TortosaCounty) BJabala杰卜莱() feud.Barony {
	return c.杰卜莱Jabala
}
    
func (c *托尔托萨TortosaCounty) BMaraclea马拉克利() feud.Barony {
	return c.马拉克利Maraclea
}
    
func (c *托尔托萨TortosaCounty) BMargat马加特() feud.Barony {
	return c.马加特Margat
}
    
func (c *托尔托萨TortosaCounty) BRuad鲁阿德() feud.Barony {
	return c.鲁阿德Ruad
}
    
func (c *托尔托萨TortosaCounty) BTortosa托尔托萨() feud.Barony {
	return c.托尔托萨Tortosa
}
    
func (c *托尔托萨TortosaCounty) BValania瓦拉尼亚() feud.Barony {
	return c.瓦拉尼亚Valania
}
    
var CTortosa托尔托萨 TortosaCounty = &托尔托萨TortosaCounty{}

func init() {
	f := CTortosa托尔托萨.(*托尔托萨TortosaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "766",
		Title:     "tortosa",
		TitleName: "托尔托萨",
		TitleCode: "c_tortosa",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈瓦比Alkhawabi = BAlkhawabi哈瓦比
	f.哈瓦比Alkhawabi.SetParent(f)
	
	f.巴勒米亚Balemia = BBalemia巴勒米亚
	f.巴勒米亚Balemia.SetParent(f)
	
	f.杰卜莱Jabala = BJabala杰卜莱
	f.杰卜莱Jabala.SetParent(f)
	
	f.马拉克利Maraclea = BMaraclea马拉克利
	f.马拉克利Maraclea.SetParent(f)
	
	f.马加特Margat = BMargat马加特
	f.马加特Margat.SetParent(f)
	
	f.鲁阿德Ruad = BRuad鲁阿德
	f.鲁阿德Ruad.SetParent(f)
	
	f.托尔托萨Tortosa = BTortosa托尔托萨
	f.托尔托萨Tortosa.SetParent(f)
	
	f.瓦拉尼亚Valania = BValania瓦拉尼亚
	f.瓦拉尼亚Valania.SetParent(f)
	
}
