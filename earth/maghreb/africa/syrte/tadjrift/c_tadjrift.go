package tadjrift

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TadjriftCounty interface {
    feud.County
    BAssultan苏丹() 	feud.Barony
    BGicherra吉舍雷() 	feud.Barony
    BOuifat乌伊费特() 	feud.Barony
    BQaryat_it_zaydan加里耶特伊斋丹() 	feud.Barony
    BTadjrift塔杰里夫特() 	feud.Barony
    BWaddan韦丹() 	feud.Barony
    BZaafran扎夫兰() 	feud.Barony
}

type 塔杰里夫特TadjriftCounty struct {
	feud.BaseCounty
	苏丹Assultan 	feud.Barony
	吉舍雷Gicherra 	feud.Barony
	乌伊费特Ouifat 	feud.Barony
	加里耶特伊斋丹Qaryat_it_zaydan 	feud.Barony
	塔杰里夫特Tadjrift 	feud.Barony
	韦丹Waddan 	feud.Barony
	扎夫兰Zaafran 	feud.Barony
}

func (c *塔杰里夫特TadjriftCounty) BAssultan苏丹() feud.Barony {
	return c.苏丹Assultan
}
    
func (c *塔杰里夫特TadjriftCounty) BGicherra吉舍雷() feud.Barony {
	return c.吉舍雷Gicherra
}
    
func (c *塔杰里夫特TadjriftCounty) BOuifat乌伊费特() feud.Barony {
	return c.乌伊费特Ouifat
}
    
func (c *塔杰里夫特TadjriftCounty) BQaryat_it_zaydan加里耶特伊斋丹() feud.Barony {
	return c.加里耶特伊斋丹Qaryat_it_zaydan
}
    
func (c *塔杰里夫特TadjriftCounty) BTadjrift塔杰里夫特() feud.Barony {
	return c.塔杰里夫特Tadjrift
}
    
func (c *塔杰里夫特TadjriftCounty) BWaddan韦丹() feud.Barony {
	return c.韦丹Waddan
}
    
func (c *塔杰里夫特TadjriftCounty) BZaafran扎夫兰() feud.Barony {
	return c.扎夫兰Zaafran
}
    
var CTadjrift塔杰里夫特 TadjriftCounty = &塔杰里夫特TadjriftCounty{}

func init() {
	f := CTadjrift塔杰里夫特.(*塔杰里夫特TadjriftCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1722",
		Title:     "tadjrift",
		TitleName: "塔杰里夫特",
		TitleCode: "c_tadjrift",
		Baronies:  map[string]feud.Barony{},
	}

	f.苏丹Assultan = BAssultan苏丹
	f.苏丹Assultan.SetParent(f)
	
	f.吉舍雷Gicherra = BGicherra吉舍雷
	f.吉舍雷Gicherra.SetParent(f)
	
	f.乌伊费特Ouifat = BOuifat乌伊费特
	f.乌伊费特Ouifat.SetParent(f)
	
	f.加里耶特伊斋丹Qaryat_it_zaydan = BQaryat_it_zaydan加里耶特伊斋丹
	f.加里耶特伊斋丹Qaryat_it_zaydan.SetParent(f)
	
	f.塔杰里夫特Tadjrift = BTadjrift塔杰里夫特
	f.塔杰里夫特Tadjrift.SetParent(f)
	
	f.韦丹Waddan = BWaddan韦丹
	f.韦丹Waddan.SetParent(f)
	
	f.扎夫兰Zaafran = BZaafran扎夫兰
	f.扎夫兰Zaafran.SetParent(f)
	
}
