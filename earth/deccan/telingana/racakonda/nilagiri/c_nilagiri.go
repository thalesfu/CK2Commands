package nilagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NilagiriCounty interface {
    feud.County
    BAvulagadda阿努拉贾达() 	feud.Barony
    BDevarakonda提婆罗军荼() 	feud.Barony
    BMaheshwaram摩醯湿伐蓝() 	feud.Barony
    BNilagiri尼罗耆厘() 	feud.Barony
    BPanagallu潘加洛() 	feud.Barony
    BVadapally瓦达帕里() 	feud.Barony
    BZuchhip叙侈波() 	feud.Barony
}

type 尼罗耆厘NilagiriCounty struct {
	feud.BaseCounty
	阿努拉贾达Avulagadda 	feud.Barony
	提婆罗军荼Devarakonda 	feud.Barony
	摩醯湿伐蓝Maheshwaram 	feud.Barony
	尼罗耆厘Nilagiri 	feud.Barony
	潘加洛Panagallu 	feud.Barony
	瓦达帕里Vadapally 	feud.Barony
	叙侈波Zuchhip 	feud.Barony
}

func (c *尼罗耆厘NilagiriCounty) BAvulagadda阿努拉贾达() feud.Barony {
	return c.阿努拉贾达Avulagadda
}
    
func (c *尼罗耆厘NilagiriCounty) BDevarakonda提婆罗军荼() feud.Barony {
	return c.提婆罗军荼Devarakonda
}
    
func (c *尼罗耆厘NilagiriCounty) BMaheshwaram摩醯湿伐蓝() feud.Barony {
	return c.摩醯湿伐蓝Maheshwaram
}
    
func (c *尼罗耆厘NilagiriCounty) BNilagiri尼罗耆厘() feud.Barony {
	return c.尼罗耆厘Nilagiri
}
    
func (c *尼罗耆厘NilagiriCounty) BPanagallu潘加洛() feud.Barony {
	return c.潘加洛Panagallu
}
    
func (c *尼罗耆厘NilagiriCounty) BVadapally瓦达帕里() feud.Barony {
	return c.瓦达帕里Vadapally
}
    
func (c *尼罗耆厘NilagiriCounty) BZuchhip叙侈波() feud.Barony {
	return c.叙侈波Zuchhip
}
    
var CNilagiri尼罗耆厘 NilagiriCounty = &尼罗耆厘NilagiriCounty{}

func init() {
	f := CNilagiri尼罗耆厘.(*尼罗耆厘NilagiriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1253",
		Title:     "nilagiri",
		TitleName: "尼罗耆厘",
		TitleCode: "c_nilagiri",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿努拉贾达Avulagadda = BAvulagadda阿努拉贾达
	f.阿努拉贾达Avulagadda.SetParent(f)
	
	f.提婆罗军荼Devarakonda = BDevarakonda提婆罗军荼
	f.提婆罗军荼Devarakonda.SetParent(f)
	
	f.摩醯湿伐蓝Maheshwaram = BMaheshwaram摩醯湿伐蓝
	f.摩醯湿伐蓝Maheshwaram.SetParent(f)
	
	f.尼罗耆厘Nilagiri = BNilagiri尼罗耆厘
	f.尼罗耆厘Nilagiri.SetParent(f)
	
	f.潘加洛Panagallu = BPanagallu潘加洛
	f.潘加洛Panagallu.SetParent(f)
	
	f.瓦达帕里Vadapally = BVadapally瓦达帕里
	f.瓦达帕里Vadapally.SetParent(f)
	
	f.叙侈波Zuchhip = BZuchhip叙侈波
	f.叙侈波Zuchhip.SetParent(f)
	
}
