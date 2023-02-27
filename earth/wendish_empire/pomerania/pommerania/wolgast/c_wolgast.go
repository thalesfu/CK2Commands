package wolgast

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WolgastCounty interface {
    feud.County
    BAnklam安克拉姆() 	feud.Barony
    BEggesin埃格辛() 	feud.Barony
    BKemnitz凯姆尼茨() 	feud.Barony
    BUeckermunde于克明德() 	feud.Barony
    BUsedom乌瑟多姆() 	feud.Barony
    BWolgast沃尔加斯特() 	feud.Barony
    BZinnowitz钦诺维茨() 	feud.Barony
    BZussow曲索() 	feud.Barony
}

type 沃尔加斯特WolgastCounty struct {
	feud.BaseCounty
	安克拉姆Anklam 	feud.Barony
	埃格辛Eggesin 	feud.Barony
	凯姆尼茨Kemnitz 	feud.Barony
	于克明德Ueckermunde 	feud.Barony
	乌瑟多姆Usedom 	feud.Barony
	沃尔加斯特Wolgast 	feud.Barony
	钦诺维茨Zinnowitz 	feud.Barony
	曲索Zussow 	feud.Barony
}

func (c *沃尔加斯特WolgastCounty) BAnklam安克拉姆() feud.Barony {
	return c.安克拉姆Anklam
}
    
func (c *沃尔加斯特WolgastCounty) BEggesin埃格辛() feud.Barony {
	return c.埃格辛Eggesin
}
    
func (c *沃尔加斯特WolgastCounty) BKemnitz凯姆尼茨() feud.Barony {
	return c.凯姆尼茨Kemnitz
}
    
func (c *沃尔加斯特WolgastCounty) BUeckermunde于克明德() feud.Barony {
	return c.于克明德Ueckermunde
}
    
func (c *沃尔加斯特WolgastCounty) BUsedom乌瑟多姆() feud.Barony {
	return c.乌瑟多姆Usedom
}
    
func (c *沃尔加斯特WolgastCounty) BWolgast沃尔加斯特() feud.Barony {
	return c.沃尔加斯特Wolgast
}
    
func (c *沃尔加斯特WolgastCounty) BZinnowitz钦诺维茨() feud.Barony {
	return c.钦诺维茨Zinnowitz
}
    
func (c *沃尔加斯特WolgastCounty) BZussow曲索() feud.Barony {
	return c.曲索Zussow
}
    
var CWolgast沃尔加斯特 WolgastCounty = &沃尔加斯特WolgastCounty{}

func init() {
	f := CWolgast沃尔加斯特.(*沃尔加斯特WolgastCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "308",
		Title:     "wolgast",
		TitleName: "沃尔加斯特",
		TitleCode: "c_wolgast",
		Baronies:  map[string]feud.Barony{},
	}

	f.安克拉姆Anklam = BAnklam安克拉姆
	f.安克拉姆Anklam.SetParent(f)
	
	f.埃格辛Eggesin = BEggesin埃格辛
	f.埃格辛Eggesin.SetParent(f)
	
	f.凯姆尼茨Kemnitz = BKemnitz凯姆尼茨
	f.凯姆尼茨Kemnitz.SetParent(f)
	
	f.于克明德Ueckermunde = BUeckermunde于克明德
	f.于克明德Ueckermunde.SetParent(f)
	
	f.乌瑟多姆Usedom = BUsedom乌瑟多姆
	f.乌瑟多姆Usedom.SetParent(f)
	
	f.沃尔加斯特Wolgast = BWolgast沃尔加斯特
	f.沃尔加斯特Wolgast.SetParent(f)
	
	f.钦诺维茨Zinnowitz = BZinnowitz钦诺维茨
	f.钦诺维茨Zinnowitz.SetParent(f)
	
	f.曲索Zussow = BZussow曲索
	f.曲索Zussow.SetParent(f)
	
}
