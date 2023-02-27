package melitene

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MeliteneCounty interface {
    feud.County
    BArca阿萨() 	feud.Barony
    BArguvan阿尔古万() 	feud.Barony
    BCorduene科都尼() 	feud.Barony
    BKigi吉基() 	feud.Barony
    BMelitene梅利蒂尼() 	feud.Barony
    BSamosata萨莫萨塔() 	feud.Barony
    BSeneqerim瑟尼格米() 	feud.Barony
    BYedisu亚迪苏() 	feud.Barony
}

type 梅利蒂尼MeliteneCounty struct {
	feud.BaseCounty
	阿萨Arca 	feud.Barony
	阿尔古万Arguvan 	feud.Barony
	科都尼Corduene 	feud.Barony
	吉基Kigi 	feud.Barony
	梅利蒂尼Melitene 	feud.Barony
	萨莫萨塔Samosata 	feud.Barony
	瑟尼格米Seneqerim 	feud.Barony
	亚迪苏Yedisu 	feud.Barony
}

func (c *梅利蒂尼MeliteneCounty) BArca阿萨() feud.Barony {
	return c.阿萨Arca
}
    
func (c *梅利蒂尼MeliteneCounty) BArguvan阿尔古万() feud.Barony {
	return c.阿尔古万Arguvan
}
    
func (c *梅利蒂尼MeliteneCounty) BCorduene科都尼() feud.Barony {
	return c.科都尼Corduene
}
    
func (c *梅利蒂尼MeliteneCounty) BKigi吉基() feud.Barony {
	return c.吉基Kigi
}
    
func (c *梅利蒂尼MeliteneCounty) BMelitene梅利蒂尼() feud.Barony {
	return c.梅利蒂尼Melitene
}
    
func (c *梅利蒂尼MeliteneCounty) BSamosata萨莫萨塔() feud.Barony {
	return c.萨莫萨塔Samosata
}
    
func (c *梅利蒂尼MeliteneCounty) BSeneqerim瑟尼格米() feud.Barony {
	return c.瑟尼格米Seneqerim
}
    
func (c *梅利蒂尼MeliteneCounty) BYedisu亚迪苏() feud.Barony {
	return c.亚迪苏Yedisu
}
    
var CMelitene梅利蒂尼 MeliteneCounty = &梅利蒂尼MeliteneCounty{}

func init() {
	f := CMelitene梅利蒂尼.(*梅利蒂尼MeliteneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "707",
		Title:     "melitene",
		TitleName: "梅利蒂尼",
		TitleCode: "c_melitene",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿萨Arca = BArca阿萨
	f.阿萨Arca.SetParent(f)
	
	f.阿尔古万Arguvan = BArguvan阿尔古万
	f.阿尔古万Arguvan.SetParent(f)
	
	f.科都尼Corduene = BCorduene科都尼
	f.科都尼Corduene.SetParent(f)
	
	f.吉基Kigi = BKigi吉基
	f.吉基Kigi.SetParent(f)
	
	f.梅利蒂尼Melitene = BMelitene梅利蒂尼
	f.梅利蒂尼Melitene.SetParent(f)
	
	f.萨莫萨塔Samosata = BSamosata萨莫萨塔
	f.萨莫萨塔Samosata.SetParent(f)
	
	f.瑟尼格米Seneqerim = BSeneqerim瑟尼格米
	f.瑟尼格米Seneqerim.SetParent(f)
	
	f.亚迪苏Yedisu = BYedisu亚迪苏
	f.亚迪苏Yedisu.SetParent(f)
	
}
