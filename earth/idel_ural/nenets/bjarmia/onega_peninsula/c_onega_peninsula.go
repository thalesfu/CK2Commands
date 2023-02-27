package onega_peninsula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Onega_peninsulaCounty interface {
    feud.County
    BKyanda基扬达() 	feud.Barony
    BLetniy列特尼() 	feud.Barony
    BLetnyaya列特尼亚亚() 	feud.Barony
    BOnega_peninsula奥涅加() 	feud.Barony
    BUna乌纳() 	feud.Barony
    BYaren_ga亚连加() 	feud.Barony
    BZhizgin日日金() 	feud.Barony
}

type 奥涅加半岛Onega_peninsulaCounty struct {
	feud.BaseCounty
	基扬达Kyanda 	feud.Barony
	列特尼Letniy 	feud.Barony
	列特尼亚亚Letnyaya 	feud.Barony
	奥涅加Onega_peninsula 	feud.Barony
	乌纳Una 	feud.Barony
	亚连加Yaren_ga 	feud.Barony
	日日金Zhizgin 	feud.Barony
}

func (c *奥涅加半岛Onega_peninsulaCounty) BKyanda基扬达() feud.Barony {
	return c.基扬达Kyanda
}
    
func (c *奥涅加半岛Onega_peninsulaCounty) BLetniy列特尼() feud.Barony {
	return c.列特尼Letniy
}
    
func (c *奥涅加半岛Onega_peninsulaCounty) BLetnyaya列特尼亚亚() feud.Barony {
	return c.列特尼亚亚Letnyaya
}
    
func (c *奥涅加半岛Onega_peninsulaCounty) BOnega_peninsula奥涅加() feud.Barony {
	return c.奥涅加Onega_peninsula
}
    
func (c *奥涅加半岛Onega_peninsulaCounty) BUna乌纳() feud.Barony {
	return c.乌纳Una
}
    
func (c *奥涅加半岛Onega_peninsulaCounty) BYaren_ga亚连加() feud.Barony {
	return c.亚连加Yaren_ga
}
    
func (c *奥涅加半岛Onega_peninsulaCounty) BZhizgin日日金() feud.Barony {
	return c.日日金Zhizgin
}
    
var COnega_peninsula奥涅加半岛 Onega_peninsulaCounty = &奥涅加半岛Onega_peninsulaCounty{}

func init() {
	f := COnega_peninsula奥涅加半岛.(*奥涅加半岛Onega_peninsulaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1821",
		Title:     "onega_peninsula",
		TitleName: "奥涅加半岛",
		TitleCode: "c_onega_peninsula",
		Baronies:  map[string]feud.Barony{},
	}

	f.基扬达Kyanda = BKyanda基扬达
	f.基扬达Kyanda.SetParent(f)
	
	f.列特尼Letniy = BLetniy列特尼
	f.列特尼Letniy.SetParent(f)
	
	f.列特尼亚亚Letnyaya = BLetnyaya列特尼亚亚
	f.列特尼亚亚Letnyaya.SetParent(f)
	
	f.奥涅加Onega_peninsula = BOnega_peninsula奥涅加
	f.奥涅加Onega_peninsula.SetParent(f)
	
	f.乌纳Una = BUna乌纳
	f.乌纳Una.SetParent(f)
	
	f.亚连加Yaren_ga = BYaren_ga亚连加
	f.亚连加Yaren_ga.SetParent(f)
	
	f.日日金Zhizgin = BZhizgin日日金
	f.日日金Zhizgin.SetParent(f)
	
}
