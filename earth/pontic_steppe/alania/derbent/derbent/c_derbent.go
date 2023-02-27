package derbent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DerbentCounty interface {
    feud.County
    BChikkulkan奇库勒坎() 	feud.Barony
    BDatuna达图那() 	feud.Barony
    BDerbent杰尔宾特() 	feud.Barony
    BHumraj胡拉贾() 	feud.Barony
    BJuma拒马() 	feud.Barony
    BKuli库利() 	feud.Barony
    BNarinkala纳伦卡拉() 	feud.Barony
    BTayus塔乌斯() 	feud.Barony
}

type 杰尔宾特DerbentCounty struct {
	feud.BaseCounty
	奇库勒坎Chikkulkan 	feud.Barony
	达图那Datuna 	feud.Barony
	杰尔宾特Derbent 	feud.Barony
	胡拉贾Humraj 	feud.Barony
	拒马Juma 	feud.Barony
	库利Kuli 	feud.Barony
	纳伦卡拉Narinkala 	feud.Barony
	塔乌斯Tayus 	feud.Barony
}

func (c *杰尔宾特DerbentCounty) BChikkulkan奇库勒坎() feud.Barony {
	return c.奇库勒坎Chikkulkan
}
    
func (c *杰尔宾特DerbentCounty) BDatuna达图那() feud.Barony {
	return c.达图那Datuna
}
    
func (c *杰尔宾特DerbentCounty) BDerbent杰尔宾特() feud.Barony {
	return c.杰尔宾特Derbent
}
    
func (c *杰尔宾特DerbentCounty) BHumraj胡拉贾() feud.Barony {
	return c.胡拉贾Humraj
}
    
func (c *杰尔宾特DerbentCounty) BJuma拒马() feud.Barony {
	return c.拒马Juma
}
    
func (c *杰尔宾特DerbentCounty) BKuli库利() feud.Barony {
	return c.库利Kuli
}
    
func (c *杰尔宾特DerbentCounty) BNarinkala纳伦卡拉() feud.Barony {
	return c.纳伦卡拉Narinkala
}
    
func (c *杰尔宾特DerbentCounty) BTayus塔乌斯() feud.Barony {
	return c.塔乌斯Tayus
}
    
var CDerbent杰尔宾特 DerbentCounty = &杰尔宾特DerbentCounty{}

func init() {
	f := CDerbent杰尔宾特.(*杰尔宾特DerbentCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "674",
		Title:     "derbent",
		TitleName: "杰尔宾特",
		TitleCode: "c_derbent",
		Baronies:  map[string]feud.Barony{},
	}

	f.奇库勒坎Chikkulkan = BChikkulkan奇库勒坎
	f.奇库勒坎Chikkulkan.SetParent(f)
	
	f.达图那Datuna = BDatuna达图那
	f.达图那Datuna.SetParent(f)
	
	f.杰尔宾特Derbent = BDerbent杰尔宾特
	f.杰尔宾特Derbent.SetParent(f)
	
	f.胡拉贾Humraj = BHumraj胡拉贾
	f.胡拉贾Humraj.SetParent(f)
	
	f.拒马Juma = BJuma拒马
	f.拒马Juma.SetParent(f)
	
	f.库利Kuli = BKuli库利
	f.库利Kuli.SetParent(f)
	
	f.纳伦卡拉Narinkala = BNarinkala纳伦卡拉
	f.纳伦卡拉Narinkala.SetParent(f)
	
	f.塔乌斯Tayus = BTayus塔乌斯
	f.塔乌斯Tayus.SetParent(f)
	
}
