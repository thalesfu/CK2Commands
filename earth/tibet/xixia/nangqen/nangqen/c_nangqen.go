package nangqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NangqenCounty interface {
    feud.County
    BDongba东坝() 	feud.Barony
    BGayong尕涌() 	feud.Barony
    BJuela觉拉() 	feud.Barony
    BMaozhuang毛庄() 	feud.Barony
    BNiangla娘拉() 	feud.Barony
    BTana_b达那() 	feud.Barony
    BXangda香达() 	feud.Barony
}

type 囊谦NangqenCounty struct {
	feud.BaseCounty
	东坝Dongba 	feud.Barony
	尕涌Gayong 	feud.Barony
	觉拉Juela 	feud.Barony
	毛庄Maozhuang 	feud.Barony
	娘拉Niangla 	feud.Barony
	达那Tana_b 	feud.Barony
	香达Xangda 	feud.Barony
}

func (c *囊谦NangqenCounty) BDongba东坝() feud.Barony {
	return c.东坝Dongba
}
    
func (c *囊谦NangqenCounty) BGayong尕涌() feud.Barony {
	return c.尕涌Gayong
}
    
func (c *囊谦NangqenCounty) BJuela觉拉() feud.Barony {
	return c.觉拉Juela
}
    
func (c *囊谦NangqenCounty) BMaozhuang毛庄() feud.Barony {
	return c.毛庄Maozhuang
}
    
func (c *囊谦NangqenCounty) BNiangla娘拉() feud.Barony {
	return c.娘拉Niangla
}
    
func (c *囊谦NangqenCounty) BTana_b达那() feud.Barony {
	return c.达那Tana_b
}
    
func (c *囊谦NangqenCounty) BXangda香达() feud.Barony {
	return c.香达Xangda
}
    
var CNangqen囊谦 NangqenCounty = &囊谦NangqenCounty{}

func init() {
	f := CNangqen囊谦.(*囊谦NangqenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1503",
		Title:     "nangqen",
		TitleName: "囊谦",
		TitleCode: "c_nangqen",
		Baronies:  map[string]feud.Barony{},
	}

	f.东坝Dongba = BDongba东坝
	f.东坝Dongba.SetParent(f)
	
	f.尕涌Gayong = BGayong尕涌
	f.尕涌Gayong.SetParent(f)
	
	f.觉拉Juela = BJuela觉拉
	f.觉拉Juela.SetParent(f)
	
	f.毛庄Maozhuang = BMaozhuang毛庄
	f.毛庄Maozhuang.SetParent(f)
	
	f.娘拉Niangla = BNiangla娘拉
	f.娘拉Niangla.SetParent(f)
	
	f.达那Tana_b = BTana_b达那
	f.达那Tana_b.SetParent(f)
	
	f.香达Xangda = BXangda香达
	f.香达Xangda.SetParent(f)
	
}
