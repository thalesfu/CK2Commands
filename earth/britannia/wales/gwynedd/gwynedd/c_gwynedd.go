package gwynedd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GwyneddCounty interface {
	feud.County
	BAbergwyngregyn阿伯古因格雷金() feud.Barony
	BCaernarfon卡那封() feud.Barony
	BConwy康威() feud.Barony
	BDegannwy迪加努伊() feud.Barony
	BHarlech哈勒赫() feud.Barony
	BLlanrhos兰罗斯() feud.Barony
}

type 圭内思GwyneddCounty struct {
	feud.BaseCounty
	阿伯古因格雷金Abergwyngregyn feud.Barony
	卡那封Caernarfon         feud.Barony
	康威Conwy               feud.Barony
	迪加努伊Degannwy          feud.Barony
	哈勒赫Harlech            feud.Barony
	兰罗斯Llanrhos           feud.Barony
}

func (c *圭内思GwyneddCounty) BAbergwyngregyn阿伯古因格雷金() feud.Barony {
	return c.阿伯古因格雷金Abergwyngregyn
}

func (c *圭内思GwyneddCounty) BCaernarfon卡那封() feud.Barony {
	return c.卡那封Caernarfon
}

func (c *圭内思GwyneddCounty) BConwy康威() feud.Barony {
	return c.康威Conwy
}

func (c *圭内思GwyneddCounty) BDegannwy迪加努伊() feud.Barony {
	return c.迪加努伊Degannwy
}

func (c *圭内思GwyneddCounty) BHarlech哈勒赫() feud.Barony {
	return c.哈勒赫Harlech
}

func (c *圭内思GwyneddCounty) BLlanrhos兰罗斯() feud.Barony {
	return c.兰罗斯Llanrhos
}

var CGwynedd圭内思 GwyneddCounty = &圭内思GwyneddCounty{}

func init() {
	f := CGwynedd圭内思.(*圭内思GwyneddCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "64",
		Title:     "gwynedd",
		TitleName: "圭内思",
		TitleCode: "c_gwynedd",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伯古因格雷金Abergwyngregyn = BAbergwyngregyn阿伯古因格雷金
	f.阿伯古因格雷金Abergwyngregyn.SetParent(f)

	f.卡那封Caernarfon = BCaernarfon卡那封
	f.卡那封Caernarfon.SetParent(f)

	f.康威Conwy = BConwy康威
	f.康威Conwy.SetParent(f)

	f.迪加努伊Degannwy = BDegannwy迪加努伊
	f.迪加努伊Degannwy.SetParent(f)

	f.哈勒赫Harlech = BHarlech哈勒赫
	f.哈勒赫Harlech.SetParent(f)

	f.兰罗斯Llanrhos = BLlanrhos兰罗斯
	f.兰罗斯Llanrhos.SetParent(f)

}
