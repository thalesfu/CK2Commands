package eu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EuCounty interface {
    feud.County
    BAumale欧马勒() 	feud.Barony
    BDrincourt德里安库尔() 	feud.Barony
    BEu厄镇() 	feud.Barony
    BForges福尔日() 	feud.Barony
    BGournay古尔奈() 	feud.Barony
    BMortemer莫特梅() 	feud.Barony
    BSerqueux塞尔克() 	feud.Barony
}

type 厄镇EuCounty struct {
	feud.BaseCounty
	欧马勒Aumale 	feud.Barony
	德里安库尔Drincourt 	feud.Barony
	厄镇Eu 	feud.Barony
	福尔日Forges 	feud.Barony
	古尔奈Gournay 	feud.Barony
	莫特梅Mortemer 	feud.Barony
	塞尔克Serqueux 	feud.Barony
}

func (c *厄镇EuCounty) BAumale欧马勒() feud.Barony {
	return c.欧马勒Aumale
}
    
func (c *厄镇EuCounty) BDrincourt德里安库尔() feud.Barony {
	return c.德里安库尔Drincourt
}
    
func (c *厄镇EuCounty) BEu厄镇() feud.Barony {
	return c.厄镇Eu
}
    
func (c *厄镇EuCounty) BForges福尔日() feud.Barony {
	return c.福尔日Forges
}
    
func (c *厄镇EuCounty) BGournay古尔奈() feud.Barony {
	return c.古尔奈Gournay
}
    
func (c *厄镇EuCounty) BMortemer莫特梅() feud.Barony {
	return c.莫特梅Mortemer
}
    
func (c *厄镇EuCounty) BSerqueux塞尔克() feud.Barony {
	return c.塞尔克Serqueux
}
    
var CEu厄镇 EuCounty = &厄镇EuCounty{}

func init() {
	f := CEu厄镇.(*厄镇EuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "96",
		Title:     "eu",
		TitleName: "厄镇",
		TitleCode: "c_eu",
		Baronies:  map[string]feud.Barony{},
	}

	f.欧马勒Aumale = BAumale欧马勒
	f.欧马勒Aumale.SetParent(f)
	
	f.德里安库尔Drincourt = BDrincourt德里安库尔
	f.德里安库尔Drincourt.SetParent(f)
	
	f.厄镇Eu = BEu厄镇
	f.厄镇Eu.SetParent(f)
	
	f.福尔日Forges = BForges福尔日
	f.福尔日Forges.SetParent(f)
	
	f.古尔奈Gournay = BGournay古尔奈
	f.古尔奈Gournay.SetParent(f)
	
	f.莫特梅Mortemer = BMortemer莫特梅
	f.莫特梅Mortemer.SetParent(f)
	
	f.塞尔克Serqueux = BSerqueux塞尔克
	f.塞尔克Serqueux.SetParent(f)
	
}
