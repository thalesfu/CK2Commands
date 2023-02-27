package temes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TemesCounty interface {
    feud.County
    BBuzias布齐亚什() 	feud.Barony
    BCsak恰克() 	feud.Barony
    BDetta代塔() 	feud.Barony
    BKaransebes考兰舍贝什() 	feud.Barony
    BKevevara凯韦瓦洛() 	feud.Barony
    BLugos卢戈斯() 	feud.Barony
    BTemesvar泰梅什堡() 	feud.Barony
    BVersecz韦尔谢茨() 	feud.Barony
}

type 泰梅什TemesCounty struct {
	feud.BaseCounty
	布齐亚什Buzias 	feud.Barony
	恰克Csak 	feud.Barony
	代塔Detta 	feud.Barony
	考兰舍贝什Karansebes 	feud.Barony
	凯韦瓦洛Kevevara 	feud.Barony
	卢戈斯Lugos 	feud.Barony
	泰梅什堡Temesvar 	feud.Barony
	韦尔谢茨Versecz 	feud.Barony
}

func (c *泰梅什TemesCounty) BBuzias布齐亚什() feud.Barony {
	return c.布齐亚什Buzias
}
    
func (c *泰梅什TemesCounty) BCsak恰克() feud.Barony {
	return c.恰克Csak
}
    
func (c *泰梅什TemesCounty) BDetta代塔() feud.Barony {
	return c.代塔Detta
}
    
func (c *泰梅什TemesCounty) BKaransebes考兰舍贝什() feud.Barony {
	return c.考兰舍贝什Karansebes
}
    
func (c *泰梅什TemesCounty) BKevevara凯韦瓦洛() feud.Barony {
	return c.凯韦瓦洛Kevevara
}
    
func (c *泰梅什TemesCounty) BLugos卢戈斯() feud.Barony {
	return c.卢戈斯Lugos
}
    
func (c *泰梅什TemesCounty) BTemesvar泰梅什堡() feud.Barony {
	return c.泰梅什堡Temesvar
}
    
func (c *泰梅什TemesCounty) BVersecz韦尔谢茨() feud.Barony {
	return c.韦尔谢茨Versecz
}
    
var CTemes泰梅什 TemesCounty = &泰梅什TemesCounty{}

func init() {
	f := CTemes泰梅什.(*泰梅什TemesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "517",
		Title:     "temes",
		TitleName: "泰梅什",
		TitleCode: "c_temes",
		Baronies:  map[string]feud.Barony{},
	}

	f.布齐亚什Buzias = BBuzias布齐亚什
	f.布齐亚什Buzias.SetParent(f)
	
	f.恰克Csak = BCsak恰克
	f.恰克Csak.SetParent(f)
	
	f.代塔Detta = BDetta代塔
	f.代塔Detta.SetParent(f)
	
	f.考兰舍贝什Karansebes = BKaransebes考兰舍贝什
	f.考兰舍贝什Karansebes.SetParent(f)
	
	f.凯韦瓦洛Kevevara = BKevevara凯韦瓦洛
	f.凯韦瓦洛Kevevara.SetParent(f)
	
	f.卢戈斯Lugos = BLugos卢戈斯
	f.卢戈斯Lugos.SetParent(f)
	
	f.泰梅什堡Temesvar = BTemesvar泰梅什堡
	f.泰梅什堡Temesvar.SetParent(f)
	
	f.韦尔谢茨Versecz = BVersecz韦尔谢茨
	f.韦尔谢茨Versecz.SetParent(f)
	
}
