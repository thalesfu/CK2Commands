package ziz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZizCounty interface {
    feud.County
    BAghoud阿古德() 	feud.Barony
    BBechar贝沙尔() 	feud.Barony
    BBoudnib布德尼卜() 	feud.Barony
    BDouar_chaouia杜瓦尔沙威亚() 	feud.Barony
    BMsanne姆桑() 	feud.Barony
    BTiskouit蒂斯库伊特() 	feud.Barony
    BZiz济兹() 	feud.Barony
}

type 济兹ZizCounty struct {
	feud.BaseCounty
	阿古德Aghoud 	feud.Barony
	贝沙尔Bechar 	feud.Barony
	布德尼卜Boudnib 	feud.Barony
	杜瓦尔沙威亚Douar_chaouia 	feud.Barony
	姆桑Msanne 	feud.Barony
	蒂斯库伊特Tiskouit 	feud.Barony
	济兹Ziz 	feud.Barony
}

func (c *济兹ZizCounty) BAghoud阿古德() feud.Barony {
	return c.阿古德Aghoud
}
    
func (c *济兹ZizCounty) BBechar贝沙尔() feud.Barony {
	return c.贝沙尔Bechar
}
    
func (c *济兹ZizCounty) BBoudnib布德尼卜() feud.Barony {
	return c.布德尼卜Boudnib
}
    
func (c *济兹ZizCounty) BDouar_chaouia杜瓦尔沙威亚() feud.Barony {
	return c.杜瓦尔沙威亚Douar_chaouia
}
    
func (c *济兹ZizCounty) BMsanne姆桑() feud.Barony {
	return c.姆桑Msanne
}
    
func (c *济兹ZizCounty) BTiskouit蒂斯库伊特() feud.Barony {
	return c.蒂斯库伊特Tiskouit
}
    
func (c *济兹ZizCounty) BZiz济兹() feud.Barony {
	return c.济兹Ziz
}
    
var CZiz济兹 ZizCounty = &济兹ZizCounty{}

func init() {
	f := CZiz济兹.(*济兹ZizCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1780",
		Title:     "ziz",
		TitleName: "济兹",
		TitleCode: "c_ziz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿古德Aghoud = BAghoud阿古德
	f.阿古德Aghoud.SetParent(f)
	
	f.贝沙尔Bechar = BBechar贝沙尔
	f.贝沙尔Bechar.SetParent(f)
	
	f.布德尼卜Boudnib = BBoudnib布德尼卜
	f.布德尼卜Boudnib.SetParent(f)
	
	f.杜瓦尔沙威亚Douar_chaouia = BDouar_chaouia杜瓦尔沙威亚
	f.杜瓦尔沙威亚Douar_chaouia.SetParent(f)
	
	f.姆桑Msanne = BMsanne姆桑
	f.姆桑Msanne.SetParent(f)
	
	f.蒂斯库伊特Tiskouit = BTiskouit蒂斯库伊特
	f.蒂斯库伊特Tiskouit.SetParent(f)
	
	f.济兹Ziz = BZiz济兹
	f.济兹Ziz.SetParent(f)
	
}
