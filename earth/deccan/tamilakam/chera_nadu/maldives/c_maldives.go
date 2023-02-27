package maldives

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaldivesCounty interface {
    feud.County
    BDhanbidhoo达恩比杜() 	feud.Barony
    BFuvahmulah孚瓦姆拉() 	feud.Barony
    BGan噶恩() 	feud.Barony
    BIsdhoo伊斯杜() 	feud.Barony
    BMahal摩诃罗() 	feud.Barony
    BMundoo曼都() 	feud.Barony
    BSalliballu萨利巴鲁() 	feud.Barony
}

type 花环洲MaldivesCounty struct {
	feud.BaseCounty
	达恩比杜Dhanbidhoo 	feud.Barony
	孚瓦姆拉Fuvahmulah 	feud.Barony
	噶恩Gan 	feud.Barony
	伊斯杜Isdhoo 	feud.Barony
	摩诃罗Mahal 	feud.Barony
	曼都Mundoo 	feud.Barony
	萨利巴鲁Salliballu 	feud.Barony
}

func (c *花环洲MaldivesCounty) BDhanbidhoo达恩比杜() feud.Barony {
	return c.达恩比杜Dhanbidhoo
}
    
func (c *花环洲MaldivesCounty) BFuvahmulah孚瓦姆拉() feud.Barony {
	return c.孚瓦姆拉Fuvahmulah
}
    
func (c *花环洲MaldivesCounty) BGan噶恩() feud.Barony {
	return c.噶恩Gan
}
    
func (c *花环洲MaldivesCounty) BIsdhoo伊斯杜() feud.Barony {
	return c.伊斯杜Isdhoo
}
    
func (c *花环洲MaldivesCounty) BMahal摩诃罗() feud.Barony {
	return c.摩诃罗Mahal
}
    
func (c *花环洲MaldivesCounty) BMundoo曼都() feud.Barony {
	return c.曼都Mundoo
}
    
func (c *花环洲MaldivesCounty) BSalliballu萨利巴鲁() feud.Barony {
	return c.萨利巴鲁Salliballu
}
    
var CMaldives花环洲 MaldivesCounty = &花环洲MaldivesCounty{}

func init() {
	f := CMaldives花环洲.(*花环洲MaldivesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1360",
		Title:     "maldives",
		TitleName: "花环洲",
		TitleCode: "c_maldives",
		Baronies:  map[string]feud.Barony{},
	}

	f.达恩比杜Dhanbidhoo = BDhanbidhoo达恩比杜
	f.达恩比杜Dhanbidhoo.SetParent(f)
	
	f.孚瓦姆拉Fuvahmulah = BFuvahmulah孚瓦姆拉
	f.孚瓦姆拉Fuvahmulah.SetParent(f)
	
	f.噶恩Gan = BGan噶恩
	f.噶恩Gan.SetParent(f)
	
	f.伊斯杜Isdhoo = BIsdhoo伊斯杜
	f.伊斯杜Isdhoo.SetParent(f)
	
	f.摩诃罗Mahal = BMahal摩诃罗
	f.摩诃罗Mahal.SetParent(f)
	
	f.曼都Mundoo = BMundoo曼都
	f.曼都Mundoo.SetParent(f)
	
	f.萨利巴鲁Salliballu = BSalliballu萨利巴鲁
	f.萨利巴鲁Salliballu.SetParent(f)
	
}
