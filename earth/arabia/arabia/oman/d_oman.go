package oman

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/oman/dhu_zabi"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/oman/duqm"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/oman/hajar"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/oman/muscat"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OmanDuke interface {
    feud.Duke
    CDhu_zabi佐尔法() 	dhu_zabi.Dhu_zabiCounty
    CDuqm尼兹瓦() 	duqm.DuqmCounty
    CHajar苏哈尔() 	hajar.HajarCounty
    CMuscat马斯喀特() 	muscat.MuscatCounty
}

type 阿曼OmanDuke struct {
	feud.BaseDuke
	佐尔法Dhu_zabi 	dhu_zabi.Dhu_zabiCounty
	尼兹瓦Duqm 	duqm.DuqmCounty
	苏哈尔Hajar 	hajar.HajarCounty
	马斯喀特Muscat 	muscat.MuscatCounty
}

func (d *阿曼OmanDuke) CDhu_zabi佐尔法() dhu_zabi.Dhu_zabiCounty {
	return d.佐尔法Dhu_zabi
}
    
func (d *阿曼OmanDuke) CDuqm尼兹瓦() duqm.DuqmCounty {
	return d.尼兹瓦Duqm
}
    
func (d *阿曼OmanDuke) CHajar苏哈尔() hajar.HajarCounty {
	return d.苏哈尔Hajar
}
    
func (d *阿曼OmanDuke) CMuscat马斯喀特() muscat.MuscatCounty {
	return d.马斯喀特Muscat
}
    
var DOman阿曼 OmanDuke = &阿曼OmanDuke{}

func init() {
	f := DOman阿曼.(*阿曼OmanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "oman",
		TitleName: "阿曼",
		TitleCode: "d_oman",
		Counties:  map[string]feud.County{},
	}

	f.佐尔法Dhu_zabi = dhu_zabi.CDhu_zabi佐尔法
	f.佐尔法Dhu_zabi.SetParent(f)
	
	f.尼兹瓦Duqm = duqm.CDuqm尼兹瓦
	f.尼兹瓦Duqm.SetParent(f)
	
	f.苏哈尔Hajar = hajar.CHajar苏哈尔
	f.苏哈尔Hajar.SetParent(f)
	
	f.马斯喀特Muscat = muscat.CMuscat马斯喀特
	f.马斯喀特Muscat.SetParent(f)
	
}
