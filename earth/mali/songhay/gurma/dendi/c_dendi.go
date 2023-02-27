package dendi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DendiCounty interface {
    feud.County
    BBura布拉() 	feud.Barony
    BIsuachara伊苏阿查拉() 	feud.Barony
    BJidu吉杜() 	feud.Barony
    BKimtir金蒂尔() 	feud.Barony
    BLangay朗盖() 	feud.Barony
    BNiamey尼亚美() 	feud.Barony
    BSay萨伊() 	feud.Barony
}

type 登迪DendiCounty struct {
	feud.BaseCounty
	布拉Bura 	feud.Barony
	伊苏阿查拉Isuachara 	feud.Barony
	吉杜Jidu 	feud.Barony
	金蒂尔Kimtir 	feud.Barony
	朗盖Langay 	feud.Barony
	尼亚美Niamey 	feud.Barony
	萨伊Say 	feud.Barony
}

func (c *登迪DendiCounty) BBura布拉() feud.Barony {
	return c.布拉Bura
}
    
func (c *登迪DendiCounty) BIsuachara伊苏阿查拉() feud.Barony {
	return c.伊苏阿查拉Isuachara
}
    
func (c *登迪DendiCounty) BJidu吉杜() feud.Barony {
	return c.吉杜Jidu
}
    
func (c *登迪DendiCounty) BKimtir金蒂尔() feud.Barony {
	return c.金蒂尔Kimtir
}
    
func (c *登迪DendiCounty) BLangay朗盖() feud.Barony {
	return c.朗盖Langay
}
    
func (c *登迪DendiCounty) BNiamey尼亚美() feud.Barony {
	return c.尼亚美Niamey
}
    
func (c *登迪DendiCounty) BSay萨伊() feud.Barony {
	return c.萨伊Say
}
    
var CDendi登迪 DendiCounty = &登迪DendiCounty{}

func init() {
	f := CDendi登迪.(*登迪DendiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1755",
		Title:     "dendi",
		TitleName: "登迪",
		TitleCode: "c_dendi",
		Baronies:  map[string]feud.Barony{},
	}

	f.布拉Bura = BBura布拉
	f.布拉Bura.SetParent(f)
	
	f.伊苏阿查拉Isuachara = BIsuachara伊苏阿查拉
	f.伊苏阿查拉Isuachara.SetParent(f)
	
	f.吉杜Jidu = BJidu吉杜
	f.吉杜Jidu.SetParent(f)
	
	f.金蒂尔Kimtir = BKimtir金蒂尔
	f.金蒂尔Kimtir.SetParent(f)
	
	f.朗盖Langay = BLangay朗盖
	f.朗盖Langay.SetParent(f)
	
	f.尼亚美Niamey = BNiamey尼亚美
	f.尼亚美Niamey.SetParent(f)
	
	f.萨伊Say = BSay萨伊
	f.萨伊Say.SetParent(f)
	
}
