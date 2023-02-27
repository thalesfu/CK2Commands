package acalapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AcalapuraCounty interface {
    feud.County
    BAcalapura不动城() 	feud.Barony
    BAmravati阿摩罗伐底() 	feud.Barony
    BBangomunda槃瞿闷荼() 	feud.Barony
    BBhojapur菩阇补罗() 	feud.Barony
    BKalam加拉姆() 	feud.Barony
    BNarnala纳尔那拉() 	feud.Barony
    BYavatmal亚瓦特马尔() 	feud.Barony
}

type 不动城AcalapuraCounty struct {
	feud.BaseCounty
	不动城Acalapura 	feud.Barony
	阿摩罗伐底Amravati 	feud.Barony
	槃瞿闷荼Bangomunda 	feud.Barony
	菩阇补罗Bhojapur 	feud.Barony
	加拉姆Kalam 	feud.Barony
	纳尔那拉Narnala 	feud.Barony
	亚瓦特马尔Yavatmal 	feud.Barony
}

func (c *不动城AcalapuraCounty) BAcalapura不动城() feud.Barony {
	return c.不动城Acalapura
}
    
func (c *不动城AcalapuraCounty) BAmravati阿摩罗伐底() feud.Barony {
	return c.阿摩罗伐底Amravati
}
    
func (c *不动城AcalapuraCounty) BBangomunda槃瞿闷荼() feud.Barony {
	return c.槃瞿闷荼Bangomunda
}
    
func (c *不动城AcalapuraCounty) BBhojapur菩阇补罗() feud.Barony {
	return c.菩阇补罗Bhojapur
}
    
func (c *不动城AcalapuraCounty) BKalam加拉姆() feud.Barony {
	return c.加拉姆Kalam
}
    
func (c *不动城AcalapuraCounty) BNarnala纳尔那拉() feud.Barony {
	return c.纳尔那拉Narnala
}
    
func (c *不动城AcalapuraCounty) BYavatmal亚瓦特马尔() feud.Barony {
	return c.亚瓦特马尔Yavatmal
}
    
var CAcalapura不动城 AcalapuraCounty = &不动城AcalapuraCounty{}

func init() {
	f := CAcalapura不动城.(*不动城AcalapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1285",
		Title:     "acalapura",
		TitleName: "不动城",
		TitleCode: "c_acalapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.不动城Acalapura = BAcalapura不动城
	f.不动城Acalapura.SetParent(f)
	
	f.阿摩罗伐底Amravati = BAmravati阿摩罗伐底
	f.阿摩罗伐底Amravati.SetParent(f)
	
	f.槃瞿闷荼Bangomunda = BBangomunda槃瞿闷荼
	f.槃瞿闷荼Bangomunda.SetParent(f)
	
	f.菩阇补罗Bhojapur = BBhojapur菩阇补罗
	f.菩阇补罗Bhojapur.SetParent(f)
	
	f.加拉姆Kalam = BKalam加拉姆
	f.加拉姆Kalam.SetParent(f)
	
	f.纳尔那拉Narnala = BNarnala纳尔那拉
	f.纳尔那拉Narnala.SetParent(f)
	
	f.亚瓦特马尔Yavatmal = BYavatmal亚瓦特马尔
	f.亚瓦特马尔Yavatmal.SetParent(f)
	
}
