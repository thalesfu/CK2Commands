package godwad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GodwadCounty interface {
    feud.County
    BAchalgarh阿遮罗姞利呬() 	feud.Barony
    BBarodiya婆卢蒂耶() 	feud.Barony
    BBhillamala毗罗摩罗() 	feud.Barony
    BCandravati旃陀罗伐底() 	feud.Barony
    BJalor迦卢尔() 	feud.Barony
    BKhudala丘陀罗() 	feud.Barony
    BNibnutvadi尼奴怛婆提() 	feud.Barony
}

type 乔荼婆荼GodwadCounty struct {
	feud.BaseCounty
	阿遮罗姞利呬Achalgarh 	feud.Barony
	婆卢蒂耶Barodiya 	feud.Barony
	毗罗摩罗Bhillamala 	feud.Barony
	旃陀罗伐底Candravati 	feud.Barony
	迦卢尔Jalor 	feud.Barony
	丘陀罗Khudala 	feud.Barony
	尼奴怛婆提Nibnutvadi 	feud.Barony
}

func (c *乔荼婆荼GodwadCounty) BAchalgarh阿遮罗姞利呬() feud.Barony {
	return c.阿遮罗姞利呬Achalgarh
}
    
func (c *乔荼婆荼GodwadCounty) BBarodiya婆卢蒂耶() feud.Barony {
	return c.婆卢蒂耶Barodiya
}
    
func (c *乔荼婆荼GodwadCounty) BBhillamala毗罗摩罗() feud.Barony {
	return c.毗罗摩罗Bhillamala
}
    
func (c *乔荼婆荼GodwadCounty) BCandravati旃陀罗伐底() feud.Barony {
	return c.旃陀罗伐底Candravati
}
    
func (c *乔荼婆荼GodwadCounty) BJalor迦卢尔() feud.Barony {
	return c.迦卢尔Jalor
}
    
func (c *乔荼婆荼GodwadCounty) BKhudala丘陀罗() feud.Barony {
	return c.丘陀罗Khudala
}
    
func (c *乔荼婆荼GodwadCounty) BNibnutvadi尼奴怛婆提() feud.Barony {
	return c.尼奴怛婆提Nibnutvadi
}
    
var CGodwad乔荼婆荼 GodwadCounty = &乔荼婆荼GodwadCounty{}

func init() {
	f := CGodwad乔荼婆荼.(*乔荼婆荼GodwadCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1174",
		Title:     "godwad",
		TitleName: "乔荼婆荼",
		TitleCode: "c_godwad",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿遮罗姞利呬Achalgarh = BAchalgarh阿遮罗姞利呬
	f.阿遮罗姞利呬Achalgarh.SetParent(f)
	
	f.婆卢蒂耶Barodiya = BBarodiya婆卢蒂耶
	f.婆卢蒂耶Barodiya.SetParent(f)
	
	f.毗罗摩罗Bhillamala = BBhillamala毗罗摩罗
	f.毗罗摩罗Bhillamala.SetParent(f)
	
	f.旃陀罗伐底Candravati = BCandravati旃陀罗伐底
	f.旃陀罗伐底Candravati.SetParent(f)
	
	f.迦卢尔Jalor = BJalor迦卢尔
	f.迦卢尔Jalor.SetParent(f)
	
	f.丘陀罗Khudala = BKhudala丘陀罗
	f.丘陀罗Khudala.SetParent(f)
	
	f.尼奴怛婆提Nibnutvadi = BNibnutvadi尼奴怛婆提
	f.尼奴怛婆提Nibnutvadi.SetParent(f)
	
}
