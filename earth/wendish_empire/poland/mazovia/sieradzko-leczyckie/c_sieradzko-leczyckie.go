package sieradzko-leczyckie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Sieradzko-leczyckieCounty interface {
    feud.County
    BCzersk切尔斯克() 	feud.Barony
    BGostynin戈斯蒂宁() 	feud.Barony
    BGrojec格罗杰茨() 	feud.Barony
    BLowicz沃维奇() 	feud.Barony
    BRawska拉瓦() 	feud.Barony
    BSochaczew索哈切夫() 	feud.Barony
    BWarszawa华沙() 	feud.Barony
}

type 戈斯蒂宁Sieradzko-leczyckieCounty struct {
	feud.BaseCounty
	切尔斯克Czersk 	feud.Barony
	戈斯蒂宁Gostynin 	feud.Barony
	格罗杰茨Grojec 	feud.Barony
	沃维奇Lowicz 	feud.Barony
	拉瓦Rawska 	feud.Barony
	索哈切夫Sochaczew 	feud.Barony
	华沙Warszawa 	feud.Barony
}

func (c *戈斯蒂宁Sieradzko-leczyckieCounty) BCzersk切尔斯克() feud.Barony {
	return c.切尔斯克Czersk
}
    
func (c *戈斯蒂宁Sieradzko-leczyckieCounty) BGostynin戈斯蒂宁() feud.Barony {
	return c.戈斯蒂宁Gostynin
}
    
func (c *戈斯蒂宁Sieradzko-leczyckieCounty) BGrojec格罗杰茨() feud.Barony {
	return c.格罗杰茨Grojec
}
    
func (c *戈斯蒂宁Sieradzko-leczyckieCounty) BLowicz沃维奇() feud.Barony {
	return c.沃维奇Lowicz
}
    
func (c *戈斯蒂宁Sieradzko-leczyckieCounty) BRawska拉瓦() feud.Barony {
	return c.拉瓦Rawska
}
    
func (c *戈斯蒂宁Sieradzko-leczyckieCounty) BSochaczew索哈切夫() feud.Barony {
	return c.索哈切夫Sochaczew
}
    
func (c *戈斯蒂宁Sieradzko-leczyckieCounty) BWarszawa华沙() feud.Barony {
	return c.华沙Warszawa
}
    
var CSieradzko-leczyckie戈斯蒂宁 Sieradzko-leczyckieCounty = &戈斯蒂宁Sieradzko-leczyckieCounty{}

func init() {
	f := CSieradzko-leczyckie戈斯蒂宁.(*戈斯蒂宁Sieradzko-leczyckieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "528",
		Title:     "sieradzko-leczyckie",
		TitleName: "戈斯蒂宁",
		TitleCode: "c_sieradzko-leczyckie",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔斯克Czersk = BCzersk切尔斯克
	f.切尔斯克Czersk.SetParent(f)
	
	f.戈斯蒂宁Gostynin = BGostynin戈斯蒂宁
	f.戈斯蒂宁Gostynin.SetParent(f)
	
	f.格罗杰茨Grojec = BGrojec格罗杰茨
	f.格罗杰茨Grojec.SetParent(f)
	
	f.沃维奇Lowicz = BLowicz沃维奇
	f.沃维奇Lowicz.SetParent(f)
	
	f.拉瓦Rawska = BRawska拉瓦
	f.拉瓦Rawska.SetParent(f)
	
	f.索哈切夫Sochaczew = BSochaczew索哈切夫
	f.索哈切夫Sochaczew.SetParent(f)
	
	f.华沙Warszawa = BWarszawa华沙
	f.华沙Warszawa.SetParent(f)
	
}
