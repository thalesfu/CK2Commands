package vishera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VisheraCounty interface {
    feud.County
    BAkchim阿克奇姆() 	feud.Barony
    BLypiya雷皮亚() 	feud.Barony
    BMoyva莫伊瓦() 	feud.Barony
    BNiols尼奥尔斯() 	feud.Barony
    BVels韦尔斯() 	feud.Barony
    BVishera维舍拉() 	feud.Barony
    BYazva亚济瓦() 	feud.Barony
}

type 维舍拉VisheraCounty struct {
	feud.BaseCounty
	阿克奇姆Akchim 	feud.Barony
	雷皮亚Lypiya 	feud.Barony
	莫伊瓦Moyva 	feud.Barony
	尼奥尔斯Niols 	feud.Barony
	韦尔斯Vels 	feud.Barony
	维舍拉Vishera 	feud.Barony
	亚济瓦Yazva 	feud.Barony
}

func (c *维舍拉VisheraCounty) BAkchim阿克奇姆() feud.Barony {
	return c.阿克奇姆Akchim
}
    
func (c *维舍拉VisheraCounty) BLypiya雷皮亚() feud.Barony {
	return c.雷皮亚Lypiya
}
    
func (c *维舍拉VisheraCounty) BMoyva莫伊瓦() feud.Barony {
	return c.莫伊瓦Moyva
}
    
func (c *维舍拉VisheraCounty) BNiols尼奥尔斯() feud.Barony {
	return c.尼奥尔斯Niols
}
    
func (c *维舍拉VisheraCounty) BVels韦尔斯() feud.Barony {
	return c.韦尔斯Vels
}
    
func (c *维舍拉VisheraCounty) BVishera维舍拉() feud.Barony {
	return c.维舍拉Vishera
}
    
func (c *维舍拉VisheraCounty) BYazva亚济瓦() feud.Barony {
	return c.亚济瓦Yazva
}
    
var CVishera维舍拉 VisheraCounty = &维舍拉VisheraCounty{}

func init() {
	f := CVishera维舍拉.(*维舍拉VisheraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1844",
		Title:     "vishera",
		TitleName: "维舍拉",
		TitleCode: "c_vishera",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克奇姆Akchim = BAkchim阿克奇姆
	f.阿克奇姆Akchim.SetParent(f)
	
	f.雷皮亚Lypiya = BLypiya雷皮亚
	f.雷皮亚Lypiya.SetParent(f)
	
	f.莫伊瓦Moyva = BMoyva莫伊瓦
	f.莫伊瓦Moyva.SetParent(f)
	
	f.尼奥尔斯Niols = BNiols尼奥尔斯
	f.尼奥尔斯Niols.SetParent(f)
	
	f.韦尔斯Vels = BVels韦尔斯
	f.韦尔斯Vels.SetParent(f)
	
	f.维舍拉Vishera = BVishera维舍拉
	f.维舍拉Vishera.SetParent(f)
	
	f.亚济瓦Yazva = BYazva亚济瓦
	f.亚济瓦Yazva.SetParent(f)
	
}
