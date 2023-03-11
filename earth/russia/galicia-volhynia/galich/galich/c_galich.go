package galich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GalichCounty interface {
    feud.County
    BBykovna布科维纳() 	feud.Barony
    BGalich加利奇() 	feud.Barony
    BKolomyia科洛梅亚() 	feud.Barony
    BRohatin罗加京() 	feud.Barony
    BTovmach托夫马奇() 	feud.Barony
    BTysmenytsia季斯梅尼齐亚() 	feud.Barony
    BVasyliv瓦西列夫() 	feud.Barony
}

type 加利奇GalichCounty struct {
	feud.BaseCounty
	布科维纳Bykovna 	feud.Barony
	加利奇Galich 	feud.Barony
	科洛梅亚Kolomyia 	feud.Barony
	罗加京Rohatin 	feud.Barony
	托夫马奇Tovmach 	feud.Barony
	季斯梅尼齐亚Tysmenytsia 	feud.Barony
	瓦西列夫Vasyliv 	feud.Barony
}

func (c *加利奇GalichCounty) BBykovna布科维纳() feud.Barony {
	return c.布科维纳Bykovna
}
    
func (c *加利奇GalichCounty) BGalich加利奇() feud.Barony {
	return c.加利奇Galich
}
    
func (c *加利奇GalichCounty) BKolomyia科洛梅亚() feud.Barony {
	return c.科洛梅亚Kolomyia
}
    
func (c *加利奇GalichCounty) BRohatin罗加京() feud.Barony {
	return c.罗加京Rohatin
}
    
func (c *加利奇GalichCounty) BTovmach托夫马奇() feud.Barony {
	return c.托夫马奇Tovmach
}
    
func (c *加利奇GalichCounty) BTysmenytsia季斯梅尼齐亚() feud.Barony {
	return c.季斯梅尼齐亚Tysmenytsia
}
    
func (c *加利奇GalichCounty) BVasyliv瓦西列夫() feud.Barony {
	return c.瓦西列夫Vasyliv
}
    
var CGalich加利奇 GalichCounty = &加利奇GalichCounty{}

func init() {
	f := CGalich加利奇.(*加利奇GalichCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "536",
		Title:     "galich",
		TitleName: "加利奇",
		TitleCode: "c_galich",
		Baronies:  map[string]feud.Barony{},
	}

	f.布科维纳Bykovna = BBykovna布科维纳
	f.布科维纳Bykovna.SetParent(f)
	
	f.加利奇Galich = BGalich加利奇
	f.加利奇Galich.SetParent(f)
	
	f.科洛梅亚Kolomyia = BKolomyia科洛梅亚
	f.科洛梅亚Kolomyia.SetParent(f)
	
	f.罗加京Rohatin = BRohatin罗加京
	f.罗加京Rohatin.SetParent(f)
	
	f.托夫马奇Tovmach = BTovmach托夫马奇
	f.托夫马奇Tovmach.SetParent(f)
	
	f.季斯梅尼齐亚Tysmenytsia = BTysmenytsia季斯梅尼齐亚
	f.季斯梅尼齐亚Tysmenytsia.SetParent(f)
	
	f.瓦西列夫Vasyliv = BVasyliv瓦西列夫
	f.瓦西列夫Vasyliv.SetParent(f)
	
}
