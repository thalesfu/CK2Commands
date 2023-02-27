package xixia

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/jiuquan"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/nagormo"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/nangqen"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/qinghai"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type XixiaKingdom interface {
    feud.Kingdom
    DJiuquan酒泉() 	jiuquan.JiuquanDuke
    DNagormo格尔木() 	nagormo.NagormoDuke
    DNangqen囊谦() 	nangqen.NangqenDuke
    DQinghai青海() 	qinghai.QinghaiDuke
}

type 西夏XixiaKingdom struct {
	feud.BaseKingdom
	酒泉Jiuquan 	jiuquan.JiuquanDuke
	格尔木Nagormo 	nagormo.NagormoDuke
	囊谦Nangqen 	nangqen.NangqenDuke
	青海Qinghai 	qinghai.QinghaiDuke
}

func (k *西夏XixiaKingdom) DJiuquan酒泉() jiuquan.JiuquanDuke {
	return k.酒泉Jiuquan
}
    
func (k *西夏XixiaKingdom) DNagormo格尔木() nagormo.NagormoDuke {
	return k.格尔木Nagormo
}
    
func (k *西夏XixiaKingdom) DNangqen囊谦() nangqen.NangqenDuke {
	return k.囊谦Nangqen
}
    
func (k *西夏XixiaKingdom) DQinghai青海() qinghai.QinghaiDuke {
	return k.青海Qinghai
}
    
var KXixia西夏 XixiaKingdom = &西夏XixiaKingdom{}

func init() {
	f := KXixia西夏.(*西夏XixiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "xixia",
		TitleName: "西夏",
		TitleCode: "k_xixia",
		Dukes:  map[string]feud.Duke{},
	}

	f.酒泉Jiuquan = jiuquan.DJiuquan酒泉
	f.酒泉Jiuquan.SetParent(f)
	
	f.格尔木Nagormo = nagormo.DNagormo格尔木
	f.格尔木Nagormo.SetParent(f)
	
	f.囊谦Nangqen = nangqen.DNangqen囊谦
	f.囊谦Nangqen.SetParent(f)
	
	f.青海Qinghai = qinghai.DQinghai青海
	f.青海Qinghai.SetParent(f)
	
}
