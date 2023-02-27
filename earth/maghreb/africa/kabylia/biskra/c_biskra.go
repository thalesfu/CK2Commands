package biskra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BiskraCounty interface {
    feud.County
    BBatna巴特纳() 	feud.Barony
    BBigou比勾() 	feud.Barony
    BBiskra比斯卡拉() 	feud.Barony
    BBranis卜拉尼斯() 	feud.Barony
    BMagra马格拉() 	feud.Barony
    BMekhadma梅克哈德马() 	feud.Barony
    BNebka奈卜卡() 	feud.Barony
    BSidiokba西迪奥克巴() 	feud.Barony
}

type 比斯卡拉BiskraCounty struct {
	feud.BaseCounty
	巴特纳Batna 	feud.Barony
	比勾Bigou 	feud.Barony
	比斯卡拉Biskra 	feud.Barony
	卜拉尼斯Branis 	feud.Barony
	马格拉Magra 	feud.Barony
	梅克哈德马Mekhadma 	feud.Barony
	奈卜卡Nebka 	feud.Barony
	西迪奥克巴Sidiokba 	feud.Barony
}

func (c *比斯卡拉BiskraCounty) BBatna巴特纳() feud.Barony {
	return c.巴特纳Batna
}
    
func (c *比斯卡拉BiskraCounty) BBigou比勾() feud.Barony {
	return c.比勾Bigou
}
    
func (c *比斯卡拉BiskraCounty) BBiskra比斯卡拉() feud.Barony {
	return c.比斯卡拉Biskra
}
    
func (c *比斯卡拉BiskraCounty) BBranis卜拉尼斯() feud.Barony {
	return c.卜拉尼斯Branis
}
    
func (c *比斯卡拉BiskraCounty) BMagra马格拉() feud.Barony {
	return c.马格拉Magra
}
    
func (c *比斯卡拉BiskraCounty) BMekhadma梅克哈德马() feud.Barony {
	return c.梅克哈德马Mekhadma
}
    
func (c *比斯卡拉BiskraCounty) BNebka奈卜卡() feud.Barony {
	return c.奈卜卡Nebka
}
    
func (c *比斯卡拉BiskraCounty) BSidiokba西迪奥克巴() feud.Barony {
	return c.西迪奥克巴Sidiokba
}
    
var CBiskra比斯卡拉 BiskraCounty = &比斯卡拉BiskraCounty{}

func init() {
	f := CBiskra比斯卡拉.(*比斯卡拉BiskraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "823",
		Title:     "biskra",
		TitleName: "比斯卡拉",
		TitleCode: "c_biskra",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴特纳Batna = BBatna巴特纳
	f.巴特纳Batna.SetParent(f)
	
	f.比勾Bigou = BBigou比勾
	f.比勾Bigou.SetParent(f)
	
	f.比斯卡拉Biskra = BBiskra比斯卡拉
	f.比斯卡拉Biskra.SetParent(f)
	
	f.卜拉尼斯Branis = BBranis卜拉尼斯
	f.卜拉尼斯Branis.SetParent(f)
	
	f.马格拉Magra = BMagra马格拉
	f.马格拉Magra.SetParent(f)
	
	f.梅克哈德马Mekhadma = BMekhadma梅克哈德马
	f.梅克哈德马Mekhadma.SetParent(f)
	
	f.奈卜卡Nebka = BNebka奈卜卡
	f.奈卜卡Nebka.SetParent(f)
	
	f.西迪奥克巴Sidiokba = BSidiokba西迪奥克巴
	f.西迪奥克巴Sidiokba.SetParent(f)
	
}
