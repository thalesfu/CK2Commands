package mozhaysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MozhayskCounty interface {
    feud.County
    BKiknur基克努尔() 	feud.Barony
    BShakhunya沙胡尼亚() 	feud.Barony
    BTonshaevo通沙耶沃() 	feud.Barony
    BUren乌连() 	feud.Barony
    BVakhtan瓦赫坦() 	feud.Barony
    BVarnavino瓦尔纳维诺() 	feud.Barony
    BVolgamozhaysk莫扎伊斯克() 	feud.Barony
    BYaransk亚兰斯克() 	feud.Barony
}

type 莫扎伊斯克MozhayskCounty struct {
	feud.BaseCounty
	基克努尔Kiknur 	feud.Barony
	沙胡尼亚Shakhunya 	feud.Barony
	通沙耶沃Tonshaevo 	feud.Barony
	乌连Uren 	feud.Barony
	瓦赫坦Vakhtan 	feud.Barony
	瓦尔纳维诺Varnavino 	feud.Barony
	莫扎伊斯克Volgamozhaysk 	feud.Barony
	亚兰斯克Yaransk 	feud.Barony
}

func (c *莫扎伊斯克MozhayskCounty) BKiknur基克努尔() feud.Barony {
	return c.基克努尔Kiknur
}
    
func (c *莫扎伊斯克MozhayskCounty) BShakhunya沙胡尼亚() feud.Barony {
	return c.沙胡尼亚Shakhunya
}
    
func (c *莫扎伊斯克MozhayskCounty) BTonshaevo通沙耶沃() feud.Barony {
	return c.通沙耶沃Tonshaevo
}
    
func (c *莫扎伊斯克MozhayskCounty) BUren乌连() feud.Barony {
	return c.乌连Uren
}
    
func (c *莫扎伊斯克MozhayskCounty) BVakhtan瓦赫坦() feud.Barony {
	return c.瓦赫坦Vakhtan
}
    
func (c *莫扎伊斯克MozhayskCounty) BVarnavino瓦尔纳维诺() feud.Barony {
	return c.瓦尔纳维诺Varnavino
}
    
func (c *莫扎伊斯克MozhayskCounty) BVolgamozhaysk莫扎伊斯克() feud.Barony {
	return c.莫扎伊斯克Volgamozhaysk
}
    
func (c *莫扎伊斯克MozhayskCounty) BYaransk亚兰斯克() feud.Barony {
	return c.亚兰斯克Yaransk
}
    
var CMozhaysk莫扎伊斯克 MozhayskCounty = &莫扎伊斯克MozhayskCounty{}

func init() {
	f := CMozhaysk莫扎伊斯克.(*莫扎伊斯克MozhayskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "587",
		Title:     "mozhaysk",
		TitleName: "莫扎伊斯克",
		TitleCode: "c_mozhaysk",
		Baronies:  map[string]feud.Barony{},
	}

	f.基克努尔Kiknur = BKiknur基克努尔
	f.基克努尔Kiknur.SetParent(f)
	
	f.沙胡尼亚Shakhunya = BShakhunya沙胡尼亚
	f.沙胡尼亚Shakhunya.SetParent(f)
	
	f.通沙耶沃Tonshaevo = BTonshaevo通沙耶沃
	f.通沙耶沃Tonshaevo.SetParent(f)
	
	f.乌连Uren = BUren乌连
	f.乌连Uren.SetParent(f)
	
	f.瓦赫坦Vakhtan = BVakhtan瓦赫坦
	f.瓦赫坦Vakhtan.SetParent(f)
	
	f.瓦尔纳维诺Varnavino = BVarnavino瓦尔纳维诺
	f.瓦尔纳维诺Varnavino.SetParent(f)
	
	f.莫扎伊斯克Volgamozhaysk = BVolgamozhaysk莫扎伊斯克
	f.莫扎伊斯克Volgamozhaysk.SetParent(f)
	
	f.亚兰斯克Yaransk = BYaransk亚兰斯克
	f.亚兰斯克Yaransk.SetParent(f)
	
}
