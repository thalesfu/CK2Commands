package dhofar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DhofarCounty interface {
    feud.County
    BAmal阿迈勒() 	feud.Barony
    BDawkah道卡() 	feud.Barony
    BDurrah杜拉赫() 	feud.Barony
    BHarzan哈赞() 	feud.Barony
    BRaysut赖苏特() 	feud.Barony
    BSalalah萨拉拉() 	feud.Barony
    BShisr什斯尔() 	feud.Barony
    BThamarit塞迈里特() 	feud.Barony
}

type 佐法尔DhofarCounty struct {
	feud.BaseCounty
	阿迈勒Amal 	feud.Barony
	道卡Dawkah 	feud.Barony
	杜拉赫Durrah 	feud.Barony
	哈赞Harzan 	feud.Barony
	赖苏特Raysut 	feud.Barony
	萨拉拉Salalah 	feud.Barony
	什斯尔Shisr 	feud.Barony
	塞迈里特Thamarit 	feud.Barony
}

func (c *佐法尔DhofarCounty) BAmal阿迈勒() feud.Barony {
	return c.阿迈勒Amal
}
    
func (c *佐法尔DhofarCounty) BDawkah道卡() feud.Barony {
	return c.道卡Dawkah
}
    
func (c *佐法尔DhofarCounty) BDurrah杜拉赫() feud.Barony {
	return c.杜拉赫Durrah
}
    
func (c *佐法尔DhofarCounty) BHarzan哈赞() feud.Barony {
	return c.哈赞Harzan
}
    
func (c *佐法尔DhofarCounty) BRaysut赖苏特() feud.Barony {
	return c.赖苏特Raysut
}
    
func (c *佐法尔DhofarCounty) BSalalah萨拉拉() feud.Barony {
	return c.萨拉拉Salalah
}
    
func (c *佐法尔DhofarCounty) BShisr什斯尔() feud.Barony {
	return c.什斯尔Shisr
}
    
func (c *佐法尔DhofarCounty) BThamarit塞迈里特() feud.Barony {
	return c.塞迈里特Thamarit
}
    
var CDhofar佐法尔 DhofarCounty = &佐法尔DhofarCounty{}

func init() {
	f := CDhofar佐法尔.(*佐法尔DhofarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "866",
		Title:     "dhofar",
		TitleName: "佐法尔",
		TitleCode: "c_dhofar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿迈勒Amal = BAmal阿迈勒
	f.阿迈勒Amal.SetParent(f)
	
	f.道卡Dawkah = BDawkah道卡
	f.道卡Dawkah.SetParent(f)
	
	f.杜拉赫Durrah = BDurrah杜拉赫
	f.杜拉赫Durrah.SetParent(f)
	
	f.哈赞Harzan = BHarzan哈赞
	f.哈赞Harzan.SetParent(f)
	
	f.赖苏特Raysut = BRaysut赖苏特
	f.赖苏特Raysut.SetParent(f)
	
	f.萨拉拉Salalah = BSalalah萨拉拉
	f.萨拉拉Salalah.SetParent(f)
	
	f.什斯尔Shisr = BShisr什斯尔
	f.什斯尔Shisr.SetParent(f)
	
	f.塞迈里特Thamarit = BThamarit塞迈里特
	f.塞迈里特Thamarit.SetParent(f)
	
}
