package semender

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SemenderCounty interface {
    feud.County
    BBurgaikala布尔格卡拉() 	feud.Barony
    BKhannalkala卡哈纳拉卡拉() 	feud.Barony
    BKhattibaku卡哈提巴库() 	feud.Barony
    BKumukh库穆赫() 	feud.Barony
    BSemender谢缅杰尔() 	feud.Barony
    BTarki塔尔基() 	feud.Barony
    BUrtseki乌斯特基() 	feud.Barony
}

type 谢缅杰尔SemenderCounty struct {
	feud.BaseCounty
	布尔格卡拉Burgaikala 	feud.Barony
	卡哈纳拉卡拉Khannalkala 	feud.Barony
	卡哈提巴库Khattibaku 	feud.Barony
	库穆赫Kumukh 	feud.Barony
	谢缅杰尔Semender 	feud.Barony
	塔尔基Tarki 	feud.Barony
	乌斯特基Urtseki 	feud.Barony
}

func (c *谢缅杰尔SemenderCounty) BBurgaikala布尔格卡拉() feud.Barony {
	return c.布尔格卡拉Burgaikala
}
    
func (c *谢缅杰尔SemenderCounty) BKhannalkala卡哈纳拉卡拉() feud.Barony {
	return c.卡哈纳拉卡拉Khannalkala
}
    
func (c *谢缅杰尔SemenderCounty) BKhattibaku卡哈提巴库() feud.Barony {
	return c.卡哈提巴库Khattibaku
}
    
func (c *谢缅杰尔SemenderCounty) BKumukh库穆赫() feud.Barony {
	return c.库穆赫Kumukh
}
    
func (c *谢缅杰尔SemenderCounty) BSemender谢缅杰尔() feud.Barony {
	return c.谢缅杰尔Semender
}
    
func (c *谢缅杰尔SemenderCounty) BTarki塔尔基() feud.Barony {
	return c.塔尔基Tarki
}
    
func (c *谢缅杰尔SemenderCounty) BUrtseki乌斯特基() feud.Barony {
	return c.乌斯特基Urtseki
}
    
var CSemender谢缅杰尔 SemenderCounty = &谢缅杰尔SemenderCounty{}

func init() {
	f := CSemender谢缅杰尔.(*谢缅杰尔SemenderCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "675",
		Title:     "semender",
		TitleName: "谢缅杰尔",
		TitleCode: "c_semender",
		Baronies:  map[string]feud.Barony{},
	}

	f.布尔格卡拉Burgaikala = BBurgaikala布尔格卡拉
	f.布尔格卡拉Burgaikala.SetParent(f)
	
	f.卡哈纳拉卡拉Khannalkala = BKhannalkala卡哈纳拉卡拉
	f.卡哈纳拉卡拉Khannalkala.SetParent(f)
	
	f.卡哈提巴库Khattibaku = BKhattibaku卡哈提巴库
	f.卡哈提巴库Khattibaku.SetParent(f)
	
	f.库穆赫Kumukh = BKumukh库穆赫
	f.库穆赫Kumukh.SetParent(f)
	
	f.谢缅杰尔Semender = BSemender谢缅杰尔
	f.谢缅杰尔Semender.SetParent(f)
	
	f.塔尔基Tarki = BTarki塔尔基
	f.塔尔基Tarki.SetParent(f)
	
	f.乌斯特基Urtseki = BUrtseki乌斯特基
	f.乌斯特基Urtseki.SetParent(f)
	
}
