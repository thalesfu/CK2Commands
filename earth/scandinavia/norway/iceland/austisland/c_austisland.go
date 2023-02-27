package austisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AustislandCounty interface {
    feud.County
    BEgilsstadir埃伊尔斯塔济() 	feud.Barony
    BEskilfjordur埃斯基尔菲厄泽() 	feud.Barony
    BHoffel霍费尔() 	feud.Barony
    BKirkjubaer基尔丘拜尔() 	feud.Barony
    BReydarfjall雷扎尔菲亚尔() 	feud.Barony
    BReydarfjordur雷扎尔菲厄泽() 	feud.Barony
    BValpjotstadur瓦尔肖夫斯塔泽() 	feud.Barony
}

type 东冰岛AustislandCounty struct {
	feud.BaseCounty
	埃伊尔斯塔济Egilsstadir 	feud.Barony
	埃斯基尔菲厄泽Eskilfjordur 	feud.Barony
	霍费尔Hoffel 	feud.Barony
	基尔丘拜尔Kirkjubaer 	feud.Barony
	雷扎尔菲亚尔Reydarfjall 	feud.Barony
	雷扎尔菲厄泽Reydarfjordur 	feud.Barony
	瓦尔肖夫斯塔泽Valpjotstadur 	feud.Barony
}

func (c *东冰岛AustislandCounty) BEgilsstadir埃伊尔斯塔济() feud.Barony {
	return c.埃伊尔斯塔济Egilsstadir
}
    
func (c *东冰岛AustislandCounty) BEskilfjordur埃斯基尔菲厄泽() feud.Barony {
	return c.埃斯基尔菲厄泽Eskilfjordur
}
    
func (c *东冰岛AustislandCounty) BHoffel霍费尔() feud.Barony {
	return c.霍费尔Hoffel
}
    
func (c *东冰岛AustislandCounty) BKirkjubaer基尔丘拜尔() feud.Barony {
	return c.基尔丘拜尔Kirkjubaer
}
    
func (c *东冰岛AustislandCounty) BReydarfjall雷扎尔菲亚尔() feud.Barony {
	return c.雷扎尔菲亚尔Reydarfjall
}
    
func (c *东冰岛AustislandCounty) BReydarfjordur雷扎尔菲厄泽() feud.Barony {
	return c.雷扎尔菲厄泽Reydarfjordur
}
    
func (c *东冰岛AustislandCounty) BValpjotstadur瓦尔肖夫斯塔泽() feud.Barony {
	return c.瓦尔肖夫斯塔泽Valpjotstadur
}
    
var CAustisland东冰岛 AustislandCounty = &东冰岛AustislandCounty{}

func init() {
	f := CAustisland东冰岛.(*东冰岛AustislandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2",
		Title:     "austisland",
		TitleName: "东冰岛",
		TitleCode: "c_austisland",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃伊尔斯塔济Egilsstadir = BEgilsstadir埃伊尔斯塔济
	f.埃伊尔斯塔济Egilsstadir.SetParent(f)
	
	f.埃斯基尔菲厄泽Eskilfjordur = BEskilfjordur埃斯基尔菲厄泽
	f.埃斯基尔菲厄泽Eskilfjordur.SetParent(f)
	
	f.霍费尔Hoffel = BHoffel霍费尔
	f.霍费尔Hoffel.SetParent(f)
	
	f.基尔丘拜尔Kirkjubaer = BKirkjubaer基尔丘拜尔
	f.基尔丘拜尔Kirkjubaer.SetParent(f)
	
	f.雷扎尔菲亚尔Reydarfjall = BReydarfjall雷扎尔菲亚尔
	f.雷扎尔菲亚尔Reydarfjall.SetParent(f)
	
	f.雷扎尔菲厄泽Reydarfjordur = BReydarfjordur雷扎尔菲厄泽
	f.雷扎尔菲厄泽Reydarfjordur.SetParent(f)
	
	f.瓦尔肖夫斯塔泽Valpjotstadur = BValpjotstadur瓦尔肖夫斯塔泽
	f.瓦尔肖夫斯塔泽Valpjotstadur.SetParent(f)
	
}
