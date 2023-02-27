package bamiyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BamiyanCounty interface {
    feud.County
    BBamiyan梵衍那() 	feud.Barony
    BIstalif伊斯塔利夫() 	feud.Barony
    BShar_i_gholghola沙尔_格拉哈拉() 	feud.Barony
    BShibar希巴尔() 	feud.Barony
    BShihak希哈克() 	feud.Barony
    BShilif希利夫() 	feud.Barony
    BZakhak扎哈克() 	feud.Barony
}

type 梵衍那BamiyanCounty struct {
	feud.BaseCounty
	梵衍那Bamiyan 	feud.Barony
	伊斯塔利夫Istalif 	feud.Barony
	沙尔_格拉哈拉Shar_i_gholghola 	feud.Barony
	希巴尔Shibar 	feud.Barony
	希哈克Shihak 	feud.Barony
	希利夫Shilif 	feud.Barony
	扎哈克Zakhak 	feud.Barony
}

func (c *梵衍那BamiyanCounty) BBamiyan梵衍那() feud.Barony {
	return c.梵衍那Bamiyan
}
    
func (c *梵衍那BamiyanCounty) BIstalif伊斯塔利夫() feud.Barony {
	return c.伊斯塔利夫Istalif
}
    
func (c *梵衍那BamiyanCounty) BShar_i_gholghola沙尔_格拉哈拉() feud.Barony {
	return c.沙尔_格拉哈拉Shar_i_gholghola
}
    
func (c *梵衍那BamiyanCounty) BShibar希巴尔() feud.Barony {
	return c.希巴尔Shibar
}
    
func (c *梵衍那BamiyanCounty) BShihak希哈克() feud.Barony {
	return c.希哈克Shihak
}
    
func (c *梵衍那BamiyanCounty) BShilif希利夫() feud.Barony {
	return c.希利夫Shilif
}
    
func (c *梵衍那BamiyanCounty) BZakhak扎哈克() feud.Barony {
	return c.扎哈克Zakhak
}
    
var CBamiyan梵衍那 BamiyanCounty = &梵衍那BamiyanCounty{}

func init() {
	f := CBamiyan梵衍那.(*梵衍那BamiyanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1189",
		Title:     "bamiyan",
		TitleName: "梵衍那",
		TitleCode: "c_bamiyan",
		Baronies:  map[string]feud.Barony{},
	}

	f.梵衍那Bamiyan = BBamiyan梵衍那
	f.梵衍那Bamiyan.SetParent(f)
	
	f.伊斯塔利夫Istalif = BIstalif伊斯塔利夫
	f.伊斯塔利夫Istalif.SetParent(f)
	
	f.沙尔_格拉哈拉Shar_i_gholghola = BShar_i_gholghola沙尔_格拉哈拉
	f.沙尔_格拉哈拉Shar_i_gholghola.SetParent(f)
	
	f.希巴尔Shibar = BShibar希巴尔
	f.希巴尔Shibar.SetParent(f)
	
	f.希哈克Shihak = BShihak希哈克
	f.希哈克Shihak.SetParent(f)
	
	f.希利夫Shilif = BShilif希利夫
	f.希利夫Shilif.SetParent(f)
	
	f.扎哈克Zakhak = BZakhak扎哈克
	f.扎哈克Zakhak.SetParent(f)
	
}
