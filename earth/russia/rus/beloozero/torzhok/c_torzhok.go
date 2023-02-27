package torzhok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TorzhokCounty interface {
    feud.County
    BFirovo菲罗沃() 	feud.Barony
    BKamenka卡缅卡() 	feud.Barony
    BOstashkov奥斯塔什科夫() 	feud.Barony
    BRashkino拉什基诺() 	feud.Barony
    BTorzhok托尔若克() 	feud.Barony
    BVyshny_volochyok上沃洛乔克() 	feud.Barony
    BZizino济济诺() 	feud.Barony
}

type 托尔若克TorzhokCounty struct {
	feud.BaseCounty
	菲罗沃Firovo 	feud.Barony
	卡缅卡Kamenka 	feud.Barony
	奥斯塔什科夫Ostashkov 	feud.Barony
	拉什基诺Rashkino 	feud.Barony
	托尔若克Torzhok 	feud.Barony
	上沃洛乔克Vyshny_volochyok 	feud.Barony
	济济诺Zizino 	feud.Barony
}

func (c *托尔若克TorzhokCounty) BFirovo菲罗沃() feud.Barony {
	return c.菲罗沃Firovo
}
    
func (c *托尔若克TorzhokCounty) BKamenka卡缅卡() feud.Barony {
	return c.卡缅卡Kamenka
}
    
func (c *托尔若克TorzhokCounty) BOstashkov奥斯塔什科夫() feud.Barony {
	return c.奥斯塔什科夫Ostashkov
}
    
func (c *托尔若克TorzhokCounty) BRashkino拉什基诺() feud.Barony {
	return c.拉什基诺Rashkino
}
    
func (c *托尔若克TorzhokCounty) BTorzhok托尔若克() feud.Barony {
	return c.托尔若克Torzhok
}
    
func (c *托尔若克TorzhokCounty) BVyshny_volochyok上沃洛乔克() feud.Barony {
	return c.上沃洛乔克Vyshny_volochyok
}
    
func (c *托尔若克TorzhokCounty) BZizino济济诺() feud.Barony {
	return c.济济诺Zizino
}
    
var CTorzhok托尔若克 TorzhokCounty = &托尔若克TorzhokCounty{}

func init() {
	f := CTorzhok托尔若克.(*托尔若克TorzhokCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "412",
		Title:     "torzhok",
		TitleName: "托尔若克",
		TitleCode: "c_torzhok",
		Baronies:  map[string]feud.Barony{},
	}

	f.菲罗沃Firovo = BFirovo菲罗沃
	f.菲罗沃Firovo.SetParent(f)
	
	f.卡缅卡Kamenka = BKamenka卡缅卡
	f.卡缅卡Kamenka.SetParent(f)
	
	f.奥斯塔什科夫Ostashkov = BOstashkov奥斯塔什科夫
	f.奥斯塔什科夫Ostashkov.SetParent(f)
	
	f.拉什基诺Rashkino = BRashkino拉什基诺
	f.拉什基诺Rashkino.SetParent(f)
	
	f.托尔若克Torzhok = BTorzhok托尔若克
	f.托尔若克Torzhok.SetParent(f)
	
	f.上沃洛乔克Vyshny_volochyok = BVyshny_volochyok上沃洛乔克
	f.上沃洛乔克Vyshny_volochyok.SetParent(f)
	
	f.济济诺Zizino = BZizino济济诺
	f.济济诺Zizino.SetParent(f)
	
}
