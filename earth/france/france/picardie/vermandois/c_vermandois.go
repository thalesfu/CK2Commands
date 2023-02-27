package vermandois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VermandoisCounty interface {
    feud.County
    BCoucy库西() 	feud.Barony
    BGauchy戈希() 	feud.Barony
    BGuise吉斯() 	feud.Barony
    BLaon拉昂() 	feud.Barony
    BMontaigu蒙泰居() 	feud.Barony
    BSigny锡尼() 	feud.Barony
    BStquentin圣康坦() 	feud.Barony
}

type 韦尔芒杜瓦VermandoisCounty struct {
	feud.BaseCounty
	库西Coucy 	feud.Barony
	戈希Gauchy 	feud.Barony
	吉斯Guise 	feud.Barony
	拉昂Laon 	feud.Barony
	蒙泰居Montaigu 	feud.Barony
	锡尼Signy 	feud.Barony
	圣康坦Stquentin 	feud.Barony
}

func (c *韦尔芒杜瓦VermandoisCounty) BCoucy库西() feud.Barony {
	return c.库西Coucy
}
    
func (c *韦尔芒杜瓦VermandoisCounty) BGauchy戈希() feud.Barony {
	return c.戈希Gauchy
}
    
func (c *韦尔芒杜瓦VermandoisCounty) BGuise吉斯() feud.Barony {
	return c.吉斯Guise
}
    
func (c *韦尔芒杜瓦VermandoisCounty) BLaon拉昂() feud.Barony {
	return c.拉昂Laon
}
    
func (c *韦尔芒杜瓦VermandoisCounty) BMontaigu蒙泰居() feud.Barony {
	return c.蒙泰居Montaigu
}
    
func (c *韦尔芒杜瓦VermandoisCounty) BSigny锡尼() feud.Barony {
	return c.锡尼Signy
}
    
func (c *韦尔芒杜瓦VermandoisCounty) BStquentin圣康坦() feud.Barony {
	return c.圣康坦Stquentin
}
    
var CVermandois韦尔芒杜瓦 VermandoisCounty = &韦尔芒杜瓦VermandoisCounty{}

func init() {
	f := CVermandois韦尔芒杜瓦.(*韦尔芒杜瓦VermandoisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "113",
		Title:     "vermandois",
		TitleName: "韦尔芒杜瓦",
		TitleCode: "c_vermandois",
		Baronies:  map[string]feud.Barony{},
	}

	f.库西Coucy = BCoucy库西
	f.库西Coucy.SetParent(f)
	
	f.戈希Gauchy = BGauchy戈希
	f.戈希Gauchy.SetParent(f)
	
	f.吉斯Guise = BGuise吉斯
	f.吉斯Guise.SetParent(f)
	
	f.拉昂Laon = BLaon拉昂
	f.拉昂Laon.SetParent(f)
	
	f.蒙泰居Montaigu = BMontaigu蒙泰居
	f.蒙泰居Montaigu.SetParent(f)
	
	f.锡尼Signy = BSigny锡尼
	f.锡尼Signy.SetParent(f)
	
	f.圣康坦Stquentin = BStquentin圣康坦
	f.圣康坦Stquentin.SetParent(f)
	
}
