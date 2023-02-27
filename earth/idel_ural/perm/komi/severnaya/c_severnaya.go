package severnaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SevernayaCounty interface {
    feud.County
    BAbar阿巴尔() 	feud.Barony
    BIlych伊雷奇() 	feud.Barony
    BPokcha波克恰() 	feud.Barony
    BSevernaya谢韦尔纳亚() 	feud.Barony
    BSevkcha谢夫克恰() 	feud.Barony
    BSherlyaga舍尔利亚加() 	feud.Barony
    BUst_ilych乌斯季伊雷奇() 	feud.Barony
}

type 谢韦尔纳亚SevernayaCounty struct {
	feud.BaseCounty
	阿巴尔Abar 	feud.Barony
	伊雷奇Ilych 	feud.Barony
	波克恰Pokcha 	feud.Barony
	谢韦尔纳亚Severnaya 	feud.Barony
	谢夫克恰Sevkcha 	feud.Barony
	舍尔利亚加Sherlyaga 	feud.Barony
	乌斯季伊雷奇Ust_ilych 	feud.Barony
}

func (c *谢韦尔纳亚SevernayaCounty) BAbar阿巴尔() feud.Barony {
	return c.阿巴尔Abar
}
    
func (c *谢韦尔纳亚SevernayaCounty) BIlych伊雷奇() feud.Barony {
	return c.伊雷奇Ilych
}
    
func (c *谢韦尔纳亚SevernayaCounty) BPokcha波克恰() feud.Barony {
	return c.波克恰Pokcha
}
    
func (c *谢韦尔纳亚SevernayaCounty) BSevernaya谢韦尔纳亚() feud.Barony {
	return c.谢韦尔纳亚Severnaya
}
    
func (c *谢韦尔纳亚SevernayaCounty) BSevkcha谢夫克恰() feud.Barony {
	return c.谢夫克恰Sevkcha
}
    
func (c *谢韦尔纳亚SevernayaCounty) BSherlyaga舍尔利亚加() feud.Barony {
	return c.舍尔利亚加Sherlyaga
}
    
func (c *谢韦尔纳亚SevernayaCounty) BUst_ilych乌斯季伊雷奇() feud.Barony {
	return c.乌斯季伊雷奇Ust_ilych
}
    
var CSevernaya谢韦尔纳亚 SevernayaCounty = &谢韦尔纳亚SevernayaCounty{}

func init() {
	f := CSevernaya谢韦尔纳亚.(*谢韦尔纳亚SevernayaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1837",
		Title:     "severnaya",
		TitleName: "谢韦尔纳亚",
		TitleCode: "c_severnaya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴尔Abar = BAbar阿巴尔
	f.阿巴尔Abar.SetParent(f)
	
	f.伊雷奇Ilych = BIlych伊雷奇
	f.伊雷奇Ilych.SetParent(f)
	
	f.波克恰Pokcha = BPokcha波克恰
	f.波克恰Pokcha.SetParent(f)
	
	f.谢韦尔纳亚Severnaya = BSevernaya谢韦尔纳亚
	f.谢韦尔纳亚Severnaya.SetParent(f)
	
	f.谢夫克恰Sevkcha = BSevkcha谢夫克恰
	f.谢夫克恰Sevkcha.SetParent(f)
	
	f.舍尔利亚加Sherlyaga = BSherlyaga舍尔利亚加
	f.舍尔利亚加Sherlyaga.SetParent(f)
	
	f.乌斯季伊雷奇Ust_ilych = BUst_ilych乌斯季伊雷奇
	f.乌斯季伊雷奇Ust_ilych.SetParent(f)
	
}
