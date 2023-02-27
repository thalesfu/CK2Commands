package sirjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SirjanCounty interface {
    feud.County
    BAbarkuh阿巴尔古() 	feud.Barony
    BBorougheeyeh博鲁盖耶() 	feud.Barony
    BDehbid代比德() 	feud.Barony
    BFaragheh法拉盖() 	feud.Barony
    BMehrabad梅赫拉巴德() 	feud.Barony
    BShahrbabak沙赫尔巴巴克() 	feud.Barony
    BSirjan锡尔詹() 	feud.Barony
}

type 锡尔詹SirjanCounty struct {
	feud.BaseCounty
	阿巴尔古Abarkuh 	feud.Barony
	博鲁盖耶Borougheeyeh 	feud.Barony
	代比德Dehbid 	feud.Barony
	法拉盖Faragheh 	feud.Barony
	梅赫拉巴德Mehrabad 	feud.Barony
	沙赫尔巴巴克Shahrbabak 	feud.Barony
	锡尔詹Sirjan 	feud.Barony
}

func (c *锡尔詹SirjanCounty) BAbarkuh阿巴尔古() feud.Barony {
	return c.阿巴尔古Abarkuh
}
    
func (c *锡尔詹SirjanCounty) BBorougheeyeh博鲁盖耶() feud.Barony {
	return c.博鲁盖耶Borougheeyeh
}
    
func (c *锡尔詹SirjanCounty) BDehbid代比德() feud.Barony {
	return c.代比德Dehbid
}
    
func (c *锡尔詹SirjanCounty) BFaragheh法拉盖() feud.Barony {
	return c.法拉盖Faragheh
}
    
func (c *锡尔詹SirjanCounty) BMehrabad梅赫拉巴德() feud.Barony {
	return c.梅赫拉巴德Mehrabad
}
    
func (c *锡尔詹SirjanCounty) BShahrbabak沙赫尔巴巴克() feud.Barony {
	return c.沙赫尔巴巴克Shahrbabak
}
    
func (c *锡尔詹SirjanCounty) BSirjan锡尔詹() feud.Barony {
	return c.锡尔詹Sirjan
}
    
var CSirjan锡尔詹 SirjanCounty = &锡尔詹SirjanCounty{}

func init() {
	f := CSirjan锡尔詹.(*锡尔詹SirjanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "640",
		Title:     "sirjan",
		TitleName: "锡尔詹",
		TitleCode: "c_sirjan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴尔古Abarkuh = BAbarkuh阿巴尔古
	f.阿巴尔古Abarkuh.SetParent(f)
	
	f.博鲁盖耶Borougheeyeh = BBorougheeyeh博鲁盖耶
	f.博鲁盖耶Borougheeyeh.SetParent(f)
	
	f.代比德Dehbid = BDehbid代比德
	f.代比德Dehbid.SetParent(f)
	
	f.法拉盖Faragheh = BFaragheh法拉盖
	f.法拉盖Faragheh.SetParent(f)
	
	f.梅赫拉巴德Mehrabad = BMehrabad梅赫拉巴德
	f.梅赫拉巴德Mehrabad.SetParent(f)
	
	f.沙赫尔巴巴克Shahrbabak = BShahrbabak沙赫尔巴巴克
	f.沙赫尔巴巴克Shahrbabak.SetParent(f)
	
	f.锡尔詹Sirjan = BSirjan锡尔詹
	f.锡尔詹Sirjan.SetParent(f)
	
}
