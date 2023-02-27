package bulgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BulgarCounty interface {
    feud.County
    BArbuga阿尔布加() 	feud.Barony
    BBalimer巴勒梅尔() 	feud.Barony
    BBulgar保加尔() 	feud.Barony
    BIq厄克() 	feud.Barony
    BKoshki科什基() 	feud.Barony
    BSuar苏瓦尔() 	feud.Barony
    BTawille泰维勒() 	feud.Barony
    BTetyushi捷秋希() 	feud.Barony
}

type 保加尔BulgarCounty struct {
	feud.BaseCounty
	阿尔布加Arbuga 	feud.Barony
	巴勒梅尔Balimer 	feud.Barony
	保加尔Bulgar 	feud.Barony
	厄克Iq 	feud.Barony
	科什基Koshki 	feud.Barony
	苏瓦尔Suar 	feud.Barony
	泰维勒Tawille 	feud.Barony
	捷秋希Tetyushi 	feud.Barony
}

func (c *保加尔BulgarCounty) BArbuga阿尔布加() feud.Barony {
	return c.阿尔布加Arbuga
}
    
func (c *保加尔BulgarCounty) BBalimer巴勒梅尔() feud.Barony {
	return c.巴勒梅尔Balimer
}
    
func (c *保加尔BulgarCounty) BBulgar保加尔() feud.Barony {
	return c.保加尔Bulgar
}
    
func (c *保加尔BulgarCounty) BIq厄克() feud.Barony {
	return c.厄克Iq
}
    
func (c *保加尔BulgarCounty) BKoshki科什基() feud.Barony {
	return c.科什基Koshki
}
    
func (c *保加尔BulgarCounty) BSuar苏瓦尔() feud.Barony {
	return c.苏瓦尔Suar
}
    
func (c *保加尔BulgarCounty) BTawille泰维勒() feud.Barony {
	return c.泰维勒Tawille
}
    
func (c *保加尔BulgarCounty) BTetyushi捷秋希() feud.Barony {
	return c.捷秋希Tetyushi
}
    
var CBulgar保加尔 BulgarCounty = &保加尔BulgarCounty{}

func init() {
	f := CBulgar保加尔.(*保加尔BulgarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "610",
		Title:     "bulgar",
		TitleName: "保加尔",
		TitleCode: "c_bulgar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔布加Arbuga = BArbuga阿尔布加
	f.阿尔布加Arbuga.SetParent(f)
	
	f.巴勒梅尔Balimer = BBalimer巴勒梅尔
	f.巴勒梅尔Balimer.SetParent(f)
	
	f.保加尔Bulgar = BBulgar保加尔
	f.保加尔Bulgar.SetParent(f)
	
	f.厄克Iq = BIq厄克
	f.厄克Iq.SetParent(f)
	
	f.科什基Koshki = BKoshki科什基
	f.科什基Koshki.SetParent(f)
	
	f.苏瓦尔Suar = BSuar苏瓦尔
	f.苏瓦尔Suar.SetParent(f)
	
	f.泰维勒Tawille = BTawille泰维勒
	f.泰维勒Tawille.SetParent(f)
	
	f.捷秋希Tetyushi = BTetyushi捷秋希
	f.捷秋希Tetyushi.SetParent(f)
	
}
