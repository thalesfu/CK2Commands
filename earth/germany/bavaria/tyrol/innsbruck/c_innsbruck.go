package innsbruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type InnsbruckCounty interface {
    feud.County
    BFugen菲根() 	feud.Barony
    BInnsbruck因斯布鲁克() 	feud.Barony
    BJenbach延巴赫() 	feud.Barony
    BKitzbuhel基茨比厄尔() 	feud.Barony
    BKufstein库夫施泰因() 	feud.Barony
    BSchwaz施瓦茨() 	feud.Barony
    BStams施塔姆斯() 	feud.Barony
}

type 因斯布鲁克InnsbruckCounty struct {
	feud.BaseCounty
	菲根Fugen 	feud.Barony
	因斯布鲁克Innsbruck 	feud.Barony
	延巴赫Jenbach 	feud.Barony
	基茨比厄尔Kitzbuhel 	feud.Barony
	库夫施泰因Kufstein 	feud.Barony
	施瓦茨Schwaz 	feud.Barony
	施塔姆斯Stams 	feud.Barony
}

func (c *因斯布鲁克InnsbruckCounty) BFugen菲根() feud.Barony {
	return c.菲根Fugen
}
    
func (c *因斯布鲁克InnsbruckCounty) BInnsbruck因斯布鲁克() feud.Barony {
	return c.因斯布鲁克Innsbruck
}
    
func (c *因斯布鲁克InnsbruckCounty) BJenbach延巴赫() feud.Barony {
	return c.延巴赫Jenbach
}
    
func (c *因斯布鲁克InnsbruckCounty) BKitzbuhel基茨比厄尔() feud.Barony {
	return c.基茨比厄尔Kitzbuhel
}
    
func (c *因斯布鲁克InnsbruckCounty) BKufstein库夫施泰因() feud.Barony {
	return c.库夫施泰因Kufstein
}
    
func (c *因斯布鲁克InnsbruckCounty) BSchwaz施瓦茨() feud.Barony {
	return c.施瓦茨Schwaz
}
    
func (c *因斯布鲁克InnsbruckCounty) BStams施塔姆斯() feud.Barony {
	return c.施塔姆斯Stams
}
    
var CInnsbruck因斯布鲁克 InnsbruckCounty = &因斯布鲁克InnsbruckCounty{}

func init() {
	f := CInnsbruck因斯布鲁克.(*因斯布鲁克InnsbruckCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "359",
		Title:     "innsbruck",
		TitleName: "因斯布鲁克",
		TitleCode: "c_innsbruck",
		Baronies:  map[string]feud.Barony{},
	}

	f.菲根Fugen = BFugen菲根
	f.菲根Fugen.SetParent(f)
	
	f.因斯布鲁克Innsbruck = BInnsbruck因斯布鲁克
	f.因斯布鲁克Innsbruck.SetParent(f)
	
	f.延巴赫Jenbach = BJenbach延巴赫
	f.延巴赫Jenbach.SetParent(f)
	
	f.基茨比厄尔Kitzbuhel = BKitzbuhel基茨比厄尔
	f.基茨比厄尔Kitzbuhel.SetParent(f)
	
	f.库夫施泰因Kufstein = BKufstein库夫施泰因
	f.库夫施泰因Kufstein.SetParent(f)
	
	f.施瓦茨Schwaz = BSchwaz施瓦茨
	f.施瓦茨Schwaz.SetParent(f)
	
	f.施塔姆斯Stams = BStams施塔姆斯
	f.施塔姆斯Stams.SetParent(f)
	
}
