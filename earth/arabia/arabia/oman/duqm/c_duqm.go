package duqm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DuqmCounty interface {
    feud.County
    BAljazir杰济拉() 	feud.Barony
    BBank班克() 	feud.Barony
    BDuqm杜库姆() 	feud.Barony
    BHarima哈日玛() 	feud.Barony
    BHubara胡巴拉() 	feud.Barony
    BMahut马胡尔() 	feud.Barony
    BMasirah马西拉() 	feud.Barony
    BNimr尼姆() 	feud.Barony
    BNizwa奈兹瓦() 	feud.Barony
}

type 尼兹瓦DuqmCounty struct {
	feud.BaseCounty
	杰济拉Aljazir 	feud.Barony
	班克Bank 	feud.Barony
	杜库姆Duqm 	feud.Barony
	哈日玛Harima 	feud.Barony
	胡巴拉Hubara 	feud.Barony
	马胡尔Mahut 	feud.Barony
	马西拉Masirah 	feud.Barony
	尼姆Nimr 	feud.Barony
	奈兹瓦Nizwa 	feud.Barony
}

func (c *尼兹瓦DuqmCounty) BAljazir杰济拉() feud.Barony {
	return c.杰济拉Aljazir
}
    
func (c *尼兹瓦DuqmCounty) BBank班克() feud.Barony {
	return c.班克Bank
}
    
func (c *尼兹瓦DuqmCounty) BDuqm杜库姆() feud.Barony {
	return c.杜库姆Duqm
}
    
func (c *尼兹瓦DuqmCounty) BHarima哈日玛() feud.Barony {
	return c.哈日玛Harima
}
    
func (c *尼兹瓦DuqmCounty) BHubara胡巴拉() feud.Barony {
	return c.胡巴拉Hubara
}
    
func (c *尼兹瓦DuqmCounty) BMahut马胡尔() feud.Barony {
	return c.马胡尔Mahut
}
    
func (c *尼兹瓦DuqmCounty) BMasirah马西拉() feud.Barony {
	return c.马西拉Masirah
}
    
func (c *尼兹瓦DuqmCounty) BNimr尼姆() feud.Barony {
	return c.尼姆Nimr
}
    
func (c *尼兹瓦DuqmCounty) BNizwa奈兹瓦() feud.Barony {
	return c.奈兹瓦Nizwa
}
    
var CDuqm尼兹瓦 DuqmCounty = &尼兹瓦DuqmCounty{}

func init() {
	f := CDuqm尼兹瓦.(*尼兹瓦DuqmCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "867",
		Title:     "duqm",
		TitleName: "尼兹瓦",
		TitleCode: "c_duqm",
		Baronies:  map[string]feud.Barony{},
	}

	f.杰济拉Aljazir = BAljazir杰济拉
	f.杰济拉Aljazir.SetParent(f)
	
	f.班克Bank = BBank班克
	f.班克Bank.SetParent(f)
	
	f.杜库姆Duqm = BDuqm杜库姆
	f.杜库姆Duqm.SetParent(f)
	
	f.哈日玛Harima = BHarima哈日玛
	f.哈日玛Harima.SetParent(f)
	
	f.胡巴拉Hubara = BHubara胡巴拉
	f.胡巴拉Hubara.SetParent(f)
	
	f.马胡尔Mahut = BMahut马胡尔
	f.马胡尔Mahut.SetParent(f)
	
	f.马西拉Masirah = BMasirah马西拉
	f.马西拉Masirah.SetParent(f)
	
	f.尼姆Nimr = BNimr尼姆
	f.尼姆Nimr.SetParent(f)
	
	f.奈兹瓦Nizwa = BNizwa奈兹瓦
	f.奈兹瓦Nizwa.SetParent(f)
	
}
