package manyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ManyapuraCounty interface {
    feud.County
    BAralaguppe阿剌罗加笈卑() 	feud.Barony
    BBasaralu巴萨罗洛() 	feud.Barony
    BGovindanahalli瞿频陀纳诃梨() 	feud.Barony
    BKikkeri基科利() 	feud.Barony
    BManyapura曼耶补罗() 	feud.Barony
    BMelukote莫鲁克特() 	feud.Barony
    BTuruvekere图鲁韦凯雷() 	feud.Barony
}

type 摩尼耶补罗ManyapuraCounty struct {
	feud.BaseCounty
	阿剌罗加笈卑Aralaguppe 	feud.Barony
	巴萨罗洛Basaralu 	feud.Barony
	瞿频陀纳诃梨Govindanahalli 	feud.Barony
	基科利Kikkeri 	feud.Barony
	曼耶补罗Manyapura 	feud.Barony
	莫鲁克特Melukote 	feud.Barony
	图鲁韦凯雷Turuvekere 	feud.Barony
}

func (c *摩尼耶补罗ManyapuraCounty) BAralaguppe阿剌罗加笈卑() feud.Barony {
	return c.阿剌罗加笈卑Aralaguppe
}
    
func (c *摩尼耶补罗ManyapuraCounty) BBasaralu巴萨罗洛() feud.Barony {
	return c.巴萨罗洛Basaralu
}
    
func (c *摩尼耶补罗ManyapuraCounty) BGovindanahalli瞿频陀纳诃梨() feud.Barony {
	return c.瞿频陀纳诃梨Govindanahalli
}
    
func (c *摩尼耶补罗ManyapuraCounty) BKikkeri基科利() feud.Barony {
	return c.基科利Kikkeri
}
    
func (c *摩尼耶补罗ManyapuraCounty) BManyapura曼耶补罗() feud.Barony {
	return c.曼耶补罗Manyapura
}
    
func (c *摩尼耶补罗ManyapuraCounty) BMelukote莫鲁克特() feud.Barony {
	return c.莫鲁克特Melukote
}
    
func (c *摩尼耶补罗ManyapuraCounty) BTuruvekere图鲁韦凯雷() feud.Barony {
	return c.图鲁韦凯雷Turuvekere
}
    
var CManyapura摩尼耶补罗 ManyapuraCounty = &摩尼耶补罗ManyapuraCounty{}

func init() {
	f := CManyapura摩尼耶补罗.(*摩尼耶补罗ManyapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1118",
		Title:     "manyapura",
		TitleName: "摩尼耶补罗",
		TitleCode: "c_manyapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿剌罗加笈卑Aralaguppe = BAralaguppe阿剌罗加笈卑
	f.阿剌罗加笈卑Aralaguppe.SetParent(f)
	
	f.巴萨罗洛Basaralu = BBasaralu巴萨罗洛
	f.巴萨罗洛Basaralu.SetParent(f)
	
	f.瞿频陀纳诃梨Govindanahalli = BGovindanahalli瞿频陀纳诃梨
	f.瞿频陀纳诃梨Govindanahalli.SetParent(f)
	
	f.基科利Kikkeri = BKikkeri基科利
	f.基科利Kikkeri.SetParent(f)
	
	f.曼耶补罗Manyapura = BManyapura曼耶补罗
	f.曼耶补罗Manyapura.SetParent(f)
	
	f.莫鲁克特Melukote = BMelukote莫鲁克特
	f.莫鲁克特Melukote.SetParent(f)
	
	f.图鲁韦凯雷Turuvekere = BTuruvekere图鲁韦凯雷
	f.图鲁韦凯雷Turuvekere.SetParent(f)
	
}
