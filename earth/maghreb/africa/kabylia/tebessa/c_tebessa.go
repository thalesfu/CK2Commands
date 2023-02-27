package tebessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TebessaCounty interface {
    feud.County
    BBaghaya巴加耶() 	feud.Barony
    BKasserine卡塞林() 	feud.Barony
    BKhenchela罕西拉() 	feud.Barony
    BMaskiyana迈斯基亚奈() 	feud.Barony
    BMaydara迈德雷() 	feud.Barony
    BTala塔拉() 	feud.Barony
    BTebessa泰贝萨() 	feud.Barony
}

type 泰贝萨TebessaCounty struct {
	feud.BaseCounty
	巴加耶Baghaya 	feud.Barony
	卡塞林Kasserine 	feud.Barony
	罕西拉Khenchela 	feud.Barony
	迈斯基亚奈Maskiyana 	feud.Barony
	迈德雷Maydara 	feud.Barony
	塔拉Tala 	feud.Barony
	泰贝萨Tebessa 	feud.Barony
}

func (c *泰贝萨TebessaCounty) BBaghaya巴加耶() feud.Barony {
	return c.巴加耶Baghaya
}
    
func (c *泰贝萨TebessaCounty) BKasserine卡塞林() feud.Barony {
	return c.卡塞林Kasserine
}
    
func (c *泰贝萨TebessaCounty) BKhenchela罕西拉() feud.Barony {
	return c.罕西拉Khenchela
}
    
func (c *泰贝萨TebessaCounty) BMaskiyana迈斯基亚奈() feud.Barony {
	return c.迈斯基亚奈Maskiyana
}
    
func (c *泰贝萨TebessaCounty) BMaydara迈德雷() feud.Barony {
	return c.迈德雷Maydara
}
    
func (c *泰贝萨TebessaCounty) BTala塔拉() feud.Barony {
	return c.塔拉Tala
}
    
func (c *泰贝萨TebessaCounty) BTebessa泰贝萨() feud.Barony {
	return c.泰贝萨Tebessa
}
    
var CTebessa泰贝萨 TebessaCounty = &泰贝萨TebessaCounty{}

func init() {
	f := CTebessa泰贝萨.(*泰贝萨TebessaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1728",
		Title:     "tebessa",
		TitleName: "泰贝萨",
		TitleCode: "c_tebessa",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴加耶Baghaya = BBaghaya巴加耶
	f.巴加耶Baghaya.SetParent(f)
	
	f.卡塞林Kasserine = BKasserine卡塞林
	f.卡塞林Kasserine.SetParent(f)
	
	f.罕西拉Khenchela = BKhenchela罕西拉
	f.罕西拉Khenchela.SetParent(f)
	
	f.迈斯基亚奈Maskiyana = BMaskiyana迈斯基亚奈
	f.迈斯基亚奈Maskiyana.SetParent(f)
	
	f.迈德雷Maydara = BMaydara迈德雷
	f.迈德雷Maydara.SetParent(f)
	
	f.塔拉Tala = BTala塔拉
	f.塔拉Tala.SetParent(f)
	
	f.泰贝萨Tebessa = BTebessa泰贝萨
	f.泰贝萨Tebessa.SetParent(f)
	
}
