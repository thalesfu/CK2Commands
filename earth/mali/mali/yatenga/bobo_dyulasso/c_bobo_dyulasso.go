package bobo_dyulasso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Bobo_dyulassoCounty interface {
    feud.County
    BAwadi阿瓦迪() 	feud.Barony
    BBobo_dyulasso博博迪乌拉索() 	feud.Barony
    BBountoumba本通巴() 	feud.Barony
    BKikele基凯莱() 	feud.Barony
    BKongola孔戈拉() 	feud.Barony
    BSikasso西卡索() 	feud.Barony
    BTaba塔巴() 	feud.Barony
}

type 博博迪乌拉索Bobo_dyulassoCounty struct {
	feud.BaseCounty
	阿瓦迪Awadi 	feud.Barony
	博博迪乌拉索Bobo_dyulasso 	feud.Barony
	本通巴Bountoumba 	feud.Barony
	基凯莱Kikele 	feud.Barony
	孔戈拉Kongola 	feud.Barony
	西卡索Sikasso 	feud.Barony
	塔巴Taba 	feud.Barony
}

func (c *博博迪乌拉索Bobo_dyulassoCounty) BAwadi阿瓦迪() feud.Barony {
	return c.阿瓦迪Awadi
}
    
func (c *博博迪乌拉索Bobo_dyulassoCounty) BBobo_dyulasso博博迪乌拉索() feud.Barony {
	return c.博博迪乌拉索Bobo_dyulasso
}
    
func (c *博博迪乌拉索Bobo_dyulassoCounty) BBountoumba本通巴() feud.Barony {
	return c.本通巴Bountoumba
}
    
func (c *博博迪乌拉索Bobo_dyulassoCounty) BKikele基凯莱() feud.Barony {
	return c.基凯莱Kikele
}
    
func (c *博博迪乌拉索Bobo_dyulassoCounty) BKongola孔戈拉() feud.Barony {
	return c.孔戈拉Kongola
}
    
func (c *博博迪乌拉索Bobo_dyulassoCounty) BSikasso西卡索() feud.Barony {
	return c.西卡索Sikasso
}
    
func (c *博博迪乌拉索Bobo_dyulassoCounty) BTaba塔巴() feud.Barony {
	return c.塔巴Taba
}
    
var CBobo_dyulasso博博迪乌拉索 Bobo_dyulassoCounty = &博博迪乌拉索Bobo_dyulassoCounty{}

func init() {
	f := CBobo_dyulasso博博迪乌拉索.(*博博迪乌拉索Bobo_dyulassoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1758",
		Title:     "bobo_dyulasso",
		TitleName: "博博迪乌拉索",
		TitleCode: "c_bobo_dyulasso",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿瓦迪Awadi = BAwadi阿瓦迪
	f.阿瓦迪Awadi.SetParent(f)
	
	f.博博迪乌拉索Bobo_dyulasso = BBobo_dyulasso博博迪乌拉索
	f.博博迪乌拉索Bobo_dyulasso.SetParent(f)
	
	f.本通巴Bountoumba = BBountoumba本通巴
	f.本通巴Bountoumba.SetParent(f)
	
	f.基凯莱Kikele = BKikele基凯莱
	f.基凯莱Kikele.SetParent(f)
	
	f.孔戈拉Kongola = BKongola孔戈拉
	f.孔戈拉Kongola.SetParent(f)
	
	f.西卡索Sikasso = BSikasso西卡索
	f.西卡索Sikasso.SetParent(f)
	
	f.塔巴Taba = BTaba塔巴
	f.塔巴Taba.SetParent(f)
	
}
