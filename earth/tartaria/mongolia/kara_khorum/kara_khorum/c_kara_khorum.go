package kara_khorum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Kara_khorumCounty interface {
    feud.County
    BBayan巴彦() 	feud.Barony
    BKara_khorum哈剌和林() 	feud.Barony
    BKhamag哈木黑() 	feud.Barony
    BLun隆() 	feud.Barony
    BLuut龙城() 	feud.Barony
    BNoin_ula诺颜乌拉() 	feud.Barony
    BZuunmod昭莫多() 	feud.Barony
}

type 哈剌和林Kara_khorumCounty struct {
	feud.BaseCounty
	巴彦Bayan 	feud.Barony
	哈剌和林Kara_khorum 	feud.Barony
	哈木黑Khamag 	feud.Barony
	隆Lun 	feud.Barony
	龙城Luut 	feud.Barony
	诺颜乌拉Noin_ula 	feud.Barony
	昭莫多Zuunmod 	feud.Barony
}

func (c *哈剌和林Kara_khorumCounty) BBayan巴彦() feud.Barony {
	return c.巴彦Bayan
}
    
func (c *哈剌和林Kara_khorumCounty) BKara_khorum哈剌和林() feud.Barony {
	return c.哈剌和林Kara_khorum
}
    
func (c *哈剌和林Kara_khorumCounty) BKhamag哈木黑() feud.Barony {
	return c.哈木黑Khamag
}
    
func (c *哈剌和林Kara_khorumCounty) BLun隆() feud.Barony {
	return c.隆Lun
}
    
func (c *哈剌和林Kara_khorumCounty) BLuut龙城() feud.Barony {
	return c.龙城Luut
}
    
func (c *哈剌和林Kara_khorumCounty) BNoin_ula诺颜乌拉() feud.Barony {
	return c.诺颜乌拉Noin_ula
}
    
func (c *哈剌和林Kara_khorumCounty) BZuunmod昭莫多() feud.Barony {
	return c.昭莫多Zuunmod
}
    
var CKara_khorum哈剌和林 Kara_khorumCounty = &哈剌和林Kara_khorumCounty{}

func init() {
	f := CKara_khorum哈剌和林.(*哈剌和林Kara_khorumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1457",
		Title:     "kara_khorum",
		TitleName: "哈剌和林",
		TitleCode: "c_kara_khorum",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴彦Bayan = BBayan巴彦
	f.巴彦Bayan.SetParent(f)
	
	f.哈剌和林Kara_khorum = BKara_khorum哈剌和林
	f.哈剌和林Kara_khorum.SetParent(f)
	
	f.哈木黑Khamag = BKhamag哈木黑
	f.哈木黑Khamag.SetParent(f)
	
	f.隆Lun = BLun隆
	f.隆Lun.SetParent(f)
	
	f.龙城Luut = BLuut龙城
	f.龙城Luut.SetParent(f)
	
	f.诺颜乌拉Noin_ula = BNoin_ula诺颜乌拉
	f.诺颜乌拉Noin_ula.SetParent(f)
	
	f.昭莫多Zuunmod = BZuunmod昭莫多
	f.昭莫多Zuunmod.SetParent(f)
	
}
