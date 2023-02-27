package kabul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KabulCounty interface {
    feud.County
    BAdinapur那揭罗喝国() 	feud.Barony
    BKabul迦布罗() 	feud.Barony
    BKapisa迦毕试() 	feud.Barony
    BKharabat哈拉巴特() 	feud.Barony
    BKunar库纳尔() 	feud.Barony
    BLampara兰帕那() 	feud.Barony
    BNagarahara那揭罗曷() 	feud.Barony
}

type 迦布罗KabulCounty struct {
	feud.BaseCounty
	那揭罗喝国Adinapur 	feud.Barony
	迦布罗Kabul 	feud.Barony
	迦毕试Kapisa 	feud.Barony
	哈拉巴特Kharabat 	feud.Barony
	库纳尔Kunar 	feud.Barony
	兰帕那Lampara 	feud.Barony
	那揭罗曷Nagarahara 	feud.Barony
}

func (c *迦布罗KabulCounty) BAdinapur那揭罗喝国() feud.Barony {
	return c.那揭罗喝国Adinapur
}
    
func (c *迦布罗KabulCounty) BKabul迦布罗() feud.Barony {
	return c.迦布罗Kabul
}
    
func (c *迦布罗KabulCounty) BKapisa迦毕试() feud.Barony {
	return c.迦毕试Kapisa
}
    
func (c *迦布罗KabulCounty) BKharabat哈拉巴特() feud.Barony {
	return c.哈拉巴特Kharabat
}
    
func (c *迦布罗KabulCounty) BKunar库纳尔() feud.Barony {
	return c.库纳尔Kunar
}
    
func (c *迦布罗KabulCounty) BLampara兰帕那() feud.Barony {
	return c.兰帕那Lampara
}
    
func (c *迦布罗KabulCounty) BNagarahara那揭罗曷() feud.Barony {
	return c.那揭罗曷Nagarahara
}
    
var CKabul迦布罗 KabulCounty = &迦布罗KabulCounty{}

func init() {
	f := CKabul迦布罗.(*迦布罗KabulCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1185",
		Title:     "kabul",
		TitleName: "迦布罗",
		TitleCode: "c_kabul",
		Baronies:  map[string]feud.Barony{},
	}

	f.那揭罗喝国Adinapur = BAdinapur那揭罗喝国
	f.那揭罗喝国Adinapur.SetParent(f)
	
	f.迦布罗Kabul = BKabul迦布罗
	f.迦布罗Kabul.SetParent(f)
	
	f.迦毕试Kapisa = BKapisa迦毕试
	f.迦毕试Kapisa.SetParent(f)
	
	f.哈拉巴特Kharabat = BKharabat哈拉巴特
	f.哈拉巴特Kharabat.SetParent(f)
	
	f.库纳尔Kunar = BKunar库纳尔
	f.库纳尔Kunar.SetParent(f)
	
	f.兰帕那Lampara = BLampara兰帕那
	f.兰帕那Lampara.SetParent(f)
	
	f.那揭罗曷Nagarahara = BNagarahara那揭罗曷
	f.那揭罗曷Nagarahara.SetParent(f)
	
}
