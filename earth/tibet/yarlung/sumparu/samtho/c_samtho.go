package samtho

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SamthoCounty interface {
    feud.County
    BDrangtsen长颈() 	feud.Barony
    BRolagang若拉岗() 	feud.Barony
    BTraoyang朝阳() 	feud.Barony
    BTsoring错仁() 	feud.Barony
    BYulpun玉盘() 	feud.Barony
    BYungpo永波() 	feud.Barony
    BZholhun雪环() 	feud.Barony
}

type 萨托SamthoCounty struct {
	feud.BaseCounty
	长颈Drangtsen 	feud.Barony
	若拉岗Rolagang 	feud.Barony
	朝阳Traoyang 	feud.Barony
	错仁Tsoring 	feud.Barony
	玉盘Yulpun 	feud.Barony
	永波Yungpo 	feud.Barony
	雪环Zholhun 	feud.Barony
}

func (c *萨托SamthoCounty) BDrangtsen长颈() feud.Barony {
	return c.长颈Drangtsen
}
    
func (c *萨托SamthoCounty) BRolagang若拉岗() feud.Barony {
	return c.若拉岗Rolagang
}
    
func (c *萨托SamthoCounty) BTraoyang朝阳() feud.Barony {
	return c.朝阳Traoyang
}
    
func (c *萨托SamthoCounty) BTsoring错仁() feud.Barony {
	return c.错仁Tsoring
}
    
func (c *萨托SamthoCounty) BYulpun玉盘() feud.Barony {
	return c.玉盘Yulpun
}
    
func (c *萨托SamthoCounty) BYungpo永波() feud.Barony {
	return c.永波Yungpo
}
    
func (c *萨托SamthoCounty) BZholhun雪环() feud.Barony {
	return c.雪环Zholhun
}
    
var CSamtho萨托 SamthoCounty = &萨托SamthoCounty{}

func init() {
	f := CSamtho萨托.(*萨托SamthoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1495",
		Title:     "samtho",
		TitleName: "萨托",
		TitleCode: "c_samtho",
		Baronies:  map[string]feud.Barony{},
	}

	f.长颈Drangtsen = BDrangtsen长颈
	f.长颈Drangtsen.SetParent(f)
	
	f.若拉岗Rolagang = BRolagang若拉岗
	f.若拉岗Rolagang.SetParent(f)
	
	f.朝阳Traoyang = BTraoyang朝阳
	f.朝阳Traoyang.SetParent(f)
	
	f.错仁Tsoring = BTsoring错仁
	f.错仁Tsoring.SetParent(f)
	
	f.玉盘Yulpun = BYulpun玉盘
	f.玉盘Yulpun.SetParent(f)
	
	f.永波Yungpo = BYungpo永波
	f.永波Yungpo.SetParent(f)
	
	f.雪环Zholhun = BZholhun雪环
	f.雪环Zholhun.SetParent(f)
	
}
