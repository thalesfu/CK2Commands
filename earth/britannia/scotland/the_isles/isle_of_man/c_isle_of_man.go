package isle_of_man

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Isle_of_manCounty interface {
    feud.County
    BDouglas道格拉斯() 	feud.Barony
    BInis_patraic帕特里克岛() 	feud.Barony
    BKirk_michael柯克迈克尔() 	feud.Barony
    BLaxey拉克西() 	feud.Barony
    BMaughold马科尔德() 	feud.Barony
    BPeel皮尔() 	feud.Barony
    BRushen鲁森() 	feud.Barony
    BSulby萨尔比() 	feud.Barony
}

type 马恩岛Isle_of_manCounty struct {
	feud.BaseCounty
	道格拉斯Douglas 	feud.Barony
	帕特里克岛Inis_patraic 	feud.Barony
	柯克迈克尔Kirk_michael 	feud.Barony
	拉克西Laxey 	feud.Barony
	马科尔德Maughold 	feud.Barony
	皮尔Peel 	feud.Barony
	鲁森Rushen 	feud.Barony
	萨尔比Sulby 	feud.Barony
}

func (c *马恩岛Isle_of_manCounty) BDouglas道格拉斯() feud.Barony {
	return c.道格拉斯Douglas
}
    
func (c *马恩岛Isle_of_manCounty) BInis_patraic帕特里克岛() feud.Barony {
	return c.帕特里克岛Inis_patraic
}
    
func (c *马恩岛Isle_of_manCounty) BKirk_michael柯克迈克尔() feud.Barony {
	return c.柯克迈克尔Kirk_michael
}
    
func (c *马恩岛Isle_of_manCounty) BLaxey拉克西() feud.Barony {
	return c.拉克西Laxey
}
    
func (c *马恩岛Isle_of_manCounty) BMaughold马科尔德() feud.Barony {
	return c.马科尔德Maughold
}
    
func (c *马恩岛Isle_of_manCounty) BPeel皮尔() feud.Barony {
	return c.皮尔Peel
}
    
func (c *马恩岛Isle_of_manCounty) BRushen鲁森() feud.Barony {
	return c.鲁森Rushen
}
    
func (c *马恩岛Isle_of_manCounty) BSulby萨尔比() feud.Barony {
	return c.萨尔比Sulby
}
    
var CIsle_of_man马恩岛 Isle_of_manCounty = &马恩岛Isle_of_manCounty{}

func init() {
	f := CIsle_of_man马恩岛.(*马恩岛Isle_of_manCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "54",
		Title:     "isle_of_man",
		TitleName: "马恩岛",
		TitleCode: "c_isle_of_man",
		Baronies:  map[string]feud.Barony{},
	}

	f.道格拉斯Douglas = BDouglas道格拉斯
	f.道格拉斯Douglas.SetParent(f)
	
	f.帕特里克岛Inis_patraic = BInis_patraic帕特里克岛
	f.帕特里克岛Inis_patraic.SetParent(f)
	
	f.柯克迈克尔Kirk_michael = BKirk_michael柯克迈克尔
	f.柯克迈克尔Kirk_michael.SetParent(f)
	
	f.拉克西Laxey = BLaxey拉克西
	f.拉克西Laxey.SetParent(f)
	
	f.马科尔德Maughold = BMaughold马科尔德
	f.马科尔德Maughold.SetParent(f)
	
	f.皮尔Peel = BPeel皮尔
	f.皮尔Peel.SetParent(f)
	
	f.鲁森Rushen = BRushen鲁森
	f.鲁森Rushen.SetParent(f)
	
	f.萨尔比Sulby = BSulby萨尔比
	f.萨尔比Sulby.SetParent(f)
	
}
