package viscaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ViscayaCounty interface {
    feud.County
    BBilbao毕尔巴鄂() 	feud.Barony
    BEibar埃瓦尔() 	feud.Barony
    BGuernica格尔尼卡() 	feud.Barony
    BIrun伊伦() 	feud.Barony
    BOnate奥尼亚蒂() 	feud.Barony
    BSansebastian圣塞瓦斯蒂安() 	feud.Barony
    BTolosa托洛萨() 	feud.Barony
    BVitoria维多利亚() 	feud.Barony
}

type 比斯开ViscayaCounty struct {
	feud.BaseCounty
	毕尔巴鄂Bilbao 	feud.Barony
	埃瓦尔Eibar 	feud.Barony
	格尔尼卡Guernica 	feud.Barony
	伊伦Irun 	feud.Barony
	奥尼亚蒂Onate 	feud.Barony
	圣塞瓦斯蒂安Sansebastian 	feud.Barony
	托洛萨Tolosa 	feud.Barony
	维多利亚Vitoria 	feud.Barony
}

func (c *比斯开ViscayaCounty) BBilbao毕尔巴鄂() feud.Barony {
	return c.毕尔巴鄂Bilbao
}
    
func (c *比斯开ViscayaCounty) BEibar埃瓦尔() feud.Barony {
	return c.埃瓦尔Eibar
}
    
func (c *比斯开ViscayaCounty) BGuernica格尔尼卡() feud.Barony {
	return c.格尔尼卡Guernica
}
    
func (c *比斯开ViscayaCounty) BIrun伊伦() feud.Barony {
	return c.伊伦Irun
}
    
func (c *比斯开ViscayaCounty) BOnate奥尼亚蒂() feud.Barony {
	return c.奥尼亚蒂Onate
}
    
func (c *比斯开ViscayaCounty) BSansebastian圣塞瓦斯蒂安() feud.Barony {
	return c.圣塞瓦斯蒂安Sansebastian
}
    
func (c *比斯开ViscayaCounty) BTolosa托洛萨() feud.Barony {
	return c.托洛萨Tolosa
}
    
func (c *比斯开ViscayaCounty) BVitoria维多利亚() feud.Barony {
	return c.维多利亚Vitoria
}
    
var CViscaya比斯开 ViscayaCounty = &比斯开ViscayaCounty{}

func init() {
	f := CViscaya比斯开.(*比斯开ViscayaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "153",
		Title:     "viscaya",
		TitleName: "比斯开",
		TitleCode: "c_viscaya",
		Baronies:  map[string]feud.Barony{},
	}

	f.毕尔巴鄂Bilbao = BBilbao毕尔巴鄂
	f.毕尔巴鄂Bilbao.SetParent(f)
	
	f.埃瓦尔Eibar = BEibar埃瓦尔
	f.埃瓦尔Eibar.SetParent(f)
	
	f.格尔尼卡Guernica = BGuernica格尔尼卡
	f.格尔尼卡Guernica.SetParent(f)
	
	f.伊伦Irun = BIrun伊伦
	f.伊伦Irun.SetParent(f)
	
	f.奥尼亚蒂Onate = BOnate奥尼亚蒂
	f.奥尼亚蒂Onate.SetParent(f)
	
	f.圣塞瓦斯蒂安Sansebastian = BSansebastian圣塞瓦斯蒂安
	f.圣塞瓦斯蒂安Sansebastian.SetParent(f)
	
	f.托洛萨Tolosa = BTolosa托洛萨
	f.托洛萨Tolosa.SetParent(f)
	
	f.维多利亚Vitoria = BVitoria维多利亚
	f.维多利亚Vitoria.SetParent(f)
	
}
