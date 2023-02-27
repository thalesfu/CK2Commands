package muztau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MuztauCounty interface {
    feud.County
    BKhokh_khot呼和浩特() 	feud.Barony
    BKurchum库尔丘姆() 	feud.Barony
    BMuztau木斯岛() 	feud.Barony
    BOskemen厄斯克门() 	feud.Barony
    BUrunkhayka乌伦海卡() 	feud.Barony
    BUryl乌雷尔() 	feud.Barony
    BZyryan济良() 	feud.Barony
}

type 西萨彦MuztauCounty struct {
	feud.BaseCounty
	呼和浩特Khokh_khot 	feud.Barony
	库尔丘姆Kurchum 	feud.Barony
	木斯岛Muztau 	feud.Barony
	厄斯克门Oskemen 	feud.Barony
	乌伦海卡Urunkhayka 	feud.Barony
	乌雷尔Uryl 	feud.Barony
	济良Zyryan 	feud.Barony
}

func (c *西萨彦MuztauCounty) BKhokh_khot呼和浩特() feud.Barony {
	return c.呼和浩特Khokh_khot
}
    
func (c *西萨彦MuztauCounty) BKurchum库尔丘姆() feud.Barony {
	return c.库尔丘姆Kurchum
}
    
func (c *西萨彦MuztauCounty) BMuztau木斯岛() feud.Barony {
	return c.木斯岛Muztau
}
    
func (c *西萨彦MuztauCounty) BOskemen厄斯克门() feud.Barony {
	return c.厄斯克门Oskemen
}
    
func (c *西萨彦MuztauCounty) BUrunkhayka乌伦海卡() feud.Barony {
	return c.乌伦海卡Urunkhayka
}
    
func (c *西萨彦MuztauCounty) BUryl乌雷尔() feud.Barony {
	return c.乌雷尔Uryl
}
    
func (c *西萨彦MuztauCounty) BZyryan济良() feud.Barony {
	return c.济良Zyryan
}
    
var CMuztau西萨彦 MuztauCounty = &西萨彦MuztauCounty{}

func init() {
	f := CMuztau西萨彦.(*西萨彦MuztauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1454",
		Title:     "muztau",
		TitleName: "西萨彦",
		TitleCode: "c_muztau",
		Baronies:  map[string]feud.Barony{},
	}

	f.呼和浩特Khokh_khot = BKhokh_khot呼和浩特
	f.呼和浩特Khokh_khot.SetParent(f)
	
	f.库尔丘姆Kurchum = BKurchum库尔丘姆
	f.库尔丘姆Kurchum.SetParent(f)
	
	f.木斯岛Muztau = BMuztau木斯岛
	f.木斯岛Muztau.SetParent(f)
	
	f.厄斯克门Oskemen = BOskemen厄斯克门
	f.厄斯克门Oskemen.SetParent(f)
	
	f.乌伦海卡Urunkhayka = BUrunkhayka乌伦海卡
	f.乌伦海卡Urunkhayka.SetParent(f)
	
	f.乌雷尔Uryl = BUryl乌雷尔
	f.乌雷尔Uryl.SetParent(f)
	
	f.济良Zyryan = BZyryan济良
	f.济良Zyryan.SetParent(f)
	
}
