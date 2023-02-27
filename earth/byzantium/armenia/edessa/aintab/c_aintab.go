package aintab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AintabCounty interface {
    feud.County
    BAintab艾因塔卜() 	feud.Barony
    BAmanus阿曼努斯() 	feud.Barony
    BDuluk杜鲁克() 	feud.Barony
    BHromgla赫罗姆格拉() 	feud.Barony
    BKyrrhos库洛斯() 	feud.Barony
    BMarash马拉什() 	feud.Barony
    BRaban雷班() 	feud.Barony
    BRavendan雷文丹() 	feud.Barony
}

type 艾因塔卜AintabCounty struct {
	feud.BaseCounty
	艾因塔卜Aintab 	feud.Barony
	阿曼努斯Amanus 	feud.Barony
	杜鲁克Duluk 	feud.Barony
	赫罗姆格拉Hromgla 	feud.Barony
	库洛斯Kyrrhos 	feud.Barony
	马拉什Marash 	feud.Barony
	雷班Raban 	feud.Barony
	雷文丹Ravendan 	feud.Barony
}

func (c *艾因塔卜AintabCounty) BAintab艾因塔卜() feud.Barony {
	return c.艾因塔卜Aintab
}
    
func (c *艾因塔卜AintabCounty) BAmanus阿曼努斯() feud.Barony {
	return c.阿曼努斯Amanus
}
    
func (c *艾因塔卜AintabCounty) BDuluk杜鲁克() feud.Barony {
	return c.杜鲁克Duluk
}
    
func (c *艾因塔卜AintabCounty) BHromgla赫罗姆格拉() feud.Barony {
	return c.赫罗姆格拉Hromgla
}
    
func (c *艾因塔卜AintabCounty) BKyrrhos库洛斯() feud.Barony {
	return c.库洛斯Kyrrhos
}
    
func (c *艾因塔卜AintabCounty) BMarash马拉什() feud.Barony {
	return c.马拉什Marash
}
    
func (c *艾因塔卜AintabCounty) BRaban雷班() feud.Barony {
	return c.雷班Raban
}
    
func (c *艾因塔卜AintabCounty) BRavendan雷文丹() feud.Barony {
	return c.雷文丹Ravendan
}
    
var CAintab艾因塔卜 AintabCounty = &艾因塔卜AintabCounty{}

func init() {
	f := CAintab艾因塔卜.(*艾因塔卜AintabCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "734",
		Title:     "aintab",
		TitleName: "艾因塔卜",
		TitleCode: "c_aintab",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾因塔卜Aintab = BAintab艾因塔卜
	f.艾因塔卜Aintab.SetParent(f)
	
	f.阿曼努斯Amanus = BAmanus阿曼努斯
	f.阿曼努斯Amanus.SetParent(f)
	
	f.杜鲁克Duluk = BDuluk杜鲁克
	f.杜鲁克Duluk.SetParent(f)
	
	f.赫罗姆格拉Hromgla = BHromgla赫罗姆格拉
	f.赫罗姆格拉Hromgla.SetParent(f)
	
	f.库洛斯Kyrrhos = BKyrrhos库洛斯
	f.库洛斯Kyrrhos.SetParent(f)
	
	f.马拉什Marash = BMarash马拉什
	f.马拉什Marash.SetParent(f)
	
	f.雷班Raban = BRaban雷班
	f.雷班Raban.SetParent(f)
	
	f.雷文丹Ravendan = BRavendan雷文丹
	f.雷文丹Ravendan.SetParent(f)
	
}
