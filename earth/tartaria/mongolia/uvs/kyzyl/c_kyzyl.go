package kyzyl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KyzylCounty interface {
    feud.County
    BAzarga阿扎尔加() 	feud.Barony
    BChadaana恰丹() 	feud.Barony
    BKyzyl克孜勒() 	feud.Barony
    BSaikhan_zam赛罕扎木() 	feud.Barony
    BTuran图兰() 	feud.Barony
    BUlagan乌拉甘() 	feud.Barony
    BUlala乌拉拉() 	feud.Barony
}

type 哈尔嘎斯KyzylCounty struct {
	feud.BaseCounty
	阿扎尔加Azarga 	feud.Barony
	恰丹Chadaana 	feud.Barony
	克孜勒Kyzyl 	feud.Barony
	赛罕扎木Saikhan_zam 	feud.Barony
	图兰Turan 	feud.Barony
	乌拉甘Ulagan 	feud.Barony
	乌拉拉Ulala 	feud.Barony
}

func (c *哈尔嘎斯KyzylCounty) BAzarga阿扎尔加() feud.Barony {
	return c.阿扎尔加Azarga
}
    
func (c *哈尔嘎斯KyzylCounty) BChadaana恰丹() feud.Barony {
	return c.恰丹Chadaana
}
    
func (c *哈尔嘎斯KyzylCounty) BKyzyl克孜勒() feud.Barony {
	return c.克孜勒Kyzyl
}
    
func (c *哈尔嘎斯KyzylCounty) BSaikhan_zam赛罕扎木() feud.Barony {
	return c.赛罕扎木Saikhan_zam
}
    
func (c *哈尔嘎斯KyzylCounty) BTuran图兰() feud.Barony {
	return c.图兰Turan
}
    
func (c *哈尔嘎斯KyzylCounty) BUlagan乌拉甘() feud.Barony {
	return c.乌拉甘Ulagan
}
    
func (c *哈尔嘎斯KyzylCounty) BUlala乌拉拉() feud.Barony {
	return c.乌拉拉Ulala
}
    
var CKyzyl哈尔嘎斯 KyzylCounty = &哈尔嘎斯KyzylCounty{}

func init() {
	f := CKyzyl哈尔嘎斯.(*哈尔嘎斯KyzylCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1459",
		Title:     "kyzyl",
		TitleName: "哈尔嘎斯",
		TitleCode: "c_kyzyl",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿扎尔加Azarga = BAzarga阿扎尔加
	f.阿扎尔加Azarga.SetParent(f)
	
	f.恰丹Chadaana = BChadaana恰丹
	f.恰丹Chadaana.SetParent(f)
	
	f.克孜勒Kyzyl = BKyzyl克孜勒
	f.克孜勒Kyzyl.SetParent(f)
	
	f.赛罕扎木Saikhan_zam = BSaikhan_zam赛罕扎木
	f.赛罕扎木Saikhan_zam.SetParent(f)
	
	f.图兰Turan = BTuran图兰
	f.图兰Turan.SetParent(f)
	
	f.乌拉甘Ulagan = BUlagan乌拉甘
	f.乌拉甘Ulagan.SetParent(f)
	
	f.乌拉拉Ulala = BUlala乌拉拉
	f.乌拉拉Ulala.SetParent(f)
	
}
