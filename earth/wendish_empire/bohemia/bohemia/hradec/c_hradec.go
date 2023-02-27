package hradec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HradecCounty interface {
    feud.County
    BCaslav恰斯拉夫() 	feud.Barony
    BChrudim赫鲁迪姆() 	feud.Barony
    BHradeckralove赫拉德茨克拉洛韦() 	feud.Barony
    BJaromer亚罗梅日() 	feud.Barony
    BKladsko克沃兹科() 	feud.Barony
    BLitomysl利托米什尔() 	feud.Barony
    BZleby日莱比() 	feud.Barony
}

type 赫拉德茨HradecCounty struct {
	feud.BaseCounty
	恰斯拉夫Caslav 	feud.Barony
	赫鲁迪姆Chrudim 	feud.Barony
	赫拉德茨克拉洛韦Hradeckralove 	feud.Barony
	亚罗梅日Jaromer 	feud.Barony
	克沃兹科Kladsko 	feud.Barony
	利托米什尔Litomysl 	feud.Barony
	日莱比Zleby 	feud.Barony
}

func (c *赫拉德茨HradecCounty) BCaslav恰斯拉夫() feud.Barony {
	return c.恰斯拉夫Caslav
}
    
func (c *赫拉德茨HradecCounty) BChrudim赫鲁迪姆() feud.Barony {
	return c.赫鲁迪姆Chrudim
}
    
func (c *赫拉德茨HradecCounty) BHradeckralove赫拉德茨克拉洛韦() feud.Barony {
	return c.赫拉德茨克拉洛韦Hradeckralove
}
    
func (c *赫拉德茨HradecCounty) BJaromer亚罗梅日() feud.Barony {
	return c.亚罗梅日Jaromer
}
    
func (c *赫拉德茨HradecCounty) BKladsko克沃兹科() feud.Barony {
	return c.克沃兹科Kladsko
}
    
func (c *赫拉德茨HradecCounty) BLitomysl利托米什尔() feud.Barony {
	return c.利托米什尔Litomysl
}
    
func (c *赫拉德茨HradecCounty) BZleby日莱比() feud.Barony {
	return c.日莱比Zleby
}
    
var CHradec赫拉德茨 HradecCounty = &赫拉德茨HradecCounty{}

func init() {
	f := CHradec赫拉德茨.(*赫拉德茨HradecCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "438",
		Title:     "hradec",
		TitleName: "赫拉德茨",
		TitleCode: "c_hradec",
		Baronies:  map[string]feud.Barony{},
	}

	f.恰斯拉夫Caslav = BCaslav恰斯拉夫
	f.恰斯拉夫Caslav.SetParent(f)
	
	f.赫鲁迪姆Chrudim = BChrudim赫鲁迪姆
	f.赫鲁迪姆Chrudim.SetParent(f)
	
	f.赫拉德茨克拉洛韦Hradeckralove = BHradeckralove赫拉德茨克拉洛韦
	f.赫拉德茨克拉洛韦Hradeckralove.SetParent(f)
	
	f.亚罗梅日Jaromer = BJaromer亚罗梅日
	f.亚罗梅日Jaromer.SetParent(f)
	
	f.克沃兹科Kladsko = BKladsko克沃兹科
	f.克沃兹科Kladsko.SetParent(f)
	
	f.利托米什尔Litomysl = BLitomysl利托米什尔
	f.利托米什尔Litomysl.SetParent(f)
	
	f.日莱比Zleby = BZleby日莱比
	f.日莱比Zleby.SetParent(f)
	
}
