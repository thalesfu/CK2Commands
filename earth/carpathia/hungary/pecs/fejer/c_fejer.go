package fejer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FejerCounty interface {
    feud.County
    BAdony奥多尼() 	feud.Barony
    BBonyhad博尼哈德() 	feud.Barony
    BDombovar栋博堡() 	feud.Barony
    BMor莫尔() 	feud.Barony
    BSarbogard沙尔博加德() 	feud.Barony
    BSzekszard塞克萨德() 	feud.Barony
    BTamasi陶马希() 	feud.Barony
    BVal瓦尔() 	feud.Barony
}

type 费耶尔FejerCounty struct {
	feud.BaseCounty
	奥多尼Adony 	feud.Barony
	博尼哈德Bonyhad 	feud.Barony
	栋博堡Dombovar 	feud.Barony
	莫尔Mor 	feud.Barony
	沙尔博加德Sarbogard 	feud.Barony
	塞克萨德Szekszard 	feud.Barony
	陶马希Tamasi 	feud.Barony
	瓦尔Val 	feud.Barony
}

func (c *费耶尔FejerCounty) BAdony奥多尼() feud.Barony {
	return c.奥多尼Adony
}
    
func (c *费耶尔FejerCounty) BBonyhad博尼哈德() feud.Barony {
	return c.博尼哈德Bonyhad
}
    
func (c *费耶尔FejerCounty) BDombovar栋博堡() feud.Barony {
	return c.栋博堡Dombovar
}
    
func (c *费耶尔FejerCounty) BMor莫尔() feud.Barony {
	return c.莫尔Mor
}
    
func (c *费耶尔FejerCounty) BSarbogard沙尔博加德() feud.Barony {
	return c.沙尔博加德Sarbogard
}
    
func (c *费耶尔FejerCounty) BSzekszard塞克萨德() feud.Barony {
	return c.塞克萨德Szekszard
}
    
func (c *费耶尔FejerCounty) BTamasi陶马希() feud.Barony {
	return c.陶马希Tamasi
}
    
func (c *费耶尔FejerCounty) BVal瓦尔() feud.Barony {
	return c.瓦尔Val
}
    
var CFejer费耶尔 FejerCounty = &费耶尔FejerCounty{}

func init() {
	f := CFejer费耶尔.(*费耶尔FejerCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "451",
		Title:     "fejer",
		TitleName: "费耶尔",
		TitleCode: "c_fejer",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥多尼Adony = BAdony奥多尼
	f.奥多尼Adony.SetParent(f)
	
	f.博尼哈德Bonyhad = BBonyhad博尼哈德
	f.博尼哈德Bonyhad.SetParent(f)
	
	f.栋博堡Dombovar = BDombovar栋博堡
	f.栋博堡Dombovar.SetParent(f)
	
	f.莫尔Mor = BMor莫尔
	f.莫尔Mor.SetParent(f)
	
	f.沙尔博加德Sarbogard = BSarbogard沙尔博加德
	f.沙尔博加德Sarbogard.SetParent(f)
	
	f.塞克萨德Szekszard = BSzekszard塞克萨德
	f.塞克萨德Szekszard.SetParent(f)
	
	f.陶马希Tamasi = BTamasi陶马希
	f.陶马希Tamasi.SetParent(f)
	
	f.瓦尔Val = BVal瓦尔
	f.瓦尔Val.SetParent(f)
	
}
