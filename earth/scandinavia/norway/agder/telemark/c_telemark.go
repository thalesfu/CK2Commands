package telemark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TelemarkCounty interface {
    feud.County
    BEidsborg埃兹堡() 	feud.Barony
    BFredriksten费德克斯腾() 	feud.Barony
    BFyresdal菲勒斯达尔() 	feud.Barony
    BGimsoy伊姆绥() 	feud.Barony
    BGrenland格伦兰() 	feud.Barony
    BHitterdals希特达拉斯() 	feud.Barony
    BSeljord塞尔尤尔() 	feud.Barony
    BSkien希恩() 	feud.Barony
}

type 泰勒马克TelemarkCounty struct {
	feud.BaseCounty
	埃兹堡Eidsborg 	feud.Barony
	费德克斯腾Fredriksten 	feud.Barony
	菲勒斯达尔Fyresdal 	feud.Barony
	伊姆绥Gimsoy 	feud.Barony
	格伦兰Grenland 	feud.Barony
	希特达拉斯Hitterdals 	feud.Barony
	塞尔尤尔Seljord 	feud.Barony
	希恩Skien 	feud.Barony
}

func (c *泰勒马克TelemarkCounty) BEidsborg埃兹堡() feud.Barony {
	return c.埃兹堡Eidsborg
}
    
func (c *泰勒马克TelemarkCounty) BFredriksten费德克斯腾() feud.Barony {
	return c.费德克斯腾Fredriksten
}
    
func (c *泰勒马克TelemarkCounty) BFyresdal菲勒斯达尔() feud.Barony {
	return c.菲勒斯达尔Fyresdal
}
    
func (c *泰勒马克TelemarkCounty) BGimsoy伊姆绥() feud.Barony {
	return c.伊姆绥Gimsoy
}
    
func (c *泰勒马克TelemarkCounty) BGrenland格伦兰() feud.Barony {
	return c.格伦兰Grenland
}
    
func (c *泰勒马克TelemarkCounty) BHitterdals希特达拉斯() feud.Barony {
	return c.希特达拉斯Hitterdals
}
    
func (c *泰勒马克TelemarkCounty) BSeljord塞尔尤尔() feud.Barony {
	return c.塞尔尤尔Seljord
}
    
func (c *泰勒马克TelemarkCounty) BSkien希恩() feud.Barony {
	return c.希恩Skien
}
    
var CTelemark泰勒马克 TelemarkCounty = &泰勒马克TelemarkCounty{}

func init() {
	f := CTelemark泰勒马克.(*泰勒马克TelemarkCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "270",
		Title:     "telemark",
		TitleName: "泰勒马克",
		TitleCode: "c_telemark",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃兹堡Eidsborg = BEidsborg埃兹堡
	f.埃兹堡Eidsborg.SetParent(f)
	
	f.费德克斯腾Fredriksten = BFredriksten费德克斯腾
	f.费德克斯腾Fredriksten.SetParent(f)
	
	f.菲勒斯达尔Fyresdal = BFyresdal菲勒斯达尔
	f.菲勒斯达尔Fyresdal.SetParent(f)
	
	f.伊姆绥Gimsoy = BGimsoy伊姆绥
	f.伊姆绥Gimsoy.SetParent(f)
	
	f.格伦兰Grenland = BGrenland格伦兰
	f.格伦兰Grenland.SetParent(f)
	
	f.希特达拉斯Hitterdals = BHitterdals希特达拉斯
	f.希特达拉斯Hitterdals.SetParent(f)
	
	f.塞尔尤尔Seljord = BSeljord塞尔尤尔
	f.塞尔尤尔Seljord.SetParent(f)
	
	f.希恩Skien = BSkien希恩
	f.希恩Skien.SetParent(f)
	
}
