package kuwait

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KuwaitCounty interface {
	feud.County
	BAlfalalheel富海希勒() feud.Barony
	BAlwafra沃夫拉() feud.Barony
	BAnthemusias安塞穆夏斯() feud.Barony
	BArrawdatayn劳扎塔因() feud.Barony
	BAshshuaybah舒艾拜() feud.Barony
	BIkaros伊卡洛斯() feud.Barony
	BKhinan希兰() feud.Barony
	BKuwait科威特() feud.Barony
}

type 科威特KuwaitCounty struct {
	feud.BaseCounty
	富海希勒Alfalalheel  feud.Barony
	沃夫拉Alwafra       feud.Barony
	安塞穆夏斯Anthemusias feud.Barony
	劳扎塔因Arrawdatayn  feud.Barony
	舒艾拜Ashshuaybah   feud.Barony
	伊卡洛斯Ikaros       feud.Barony
	希兰Khinan         feud.Barony
	科威特Kuwait        feud.Barony
}

func (c *科威特KuwaitCounty) BAlfalalheel富海希勒() feud.Barony {
	return c.富海希勒Alfalalheel
}

func (c *科威特KuwaitCounty) BAlwafra沃夫拉() feud.Barony {
	return c.沃夫拉Alwafra
}

func (c *科威特KuwaitCounty) BAnthemusias安塞穆夏斯() feud.Barony {
	return c.安塞穆夏斯Anthemusias
}

func (c *科威特KuwaitCounty) BArrawdatayn劳扎塔因() feud.Barony {
	return c.劳扎塔因Arrawdatayn
}

func (c *科威特KuwaitCounty) BAshshuaybah舒艾拜() feud.Barony {
	return c.舒艾拜Ashshuaybah
}

func (c *科威特KuwaitCounty) BIkaros伊卡洛斯() feud.Barony {
	return c.伊卡洛斯Ikaros
}

func (c *科威特KuwaitCounty) BKhinan希兰() feud.Barony {
	return c.希兰Khinan
}

func (c *科威特KuwaitCounty) BKuwait科威特() feud.Barony {
	return c.科威特Kuwait
}

var CKuwait科威特 KuwaitCounty = &科威特KuwaitCounty{}

func init() {
	f := CKuwait科威特.(*科威特KuwaitCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "650",
		Title:     "kuwait",
		TitleName: "科威特",
		TitleCode: "c_kuwait",
		Baronies:  map[string]feud.Barony{},
	}

	f.富海希勒Alfalalheel = BAlfalalheel富海希勒
	f.富海希勒Alfalalheel.SetParent(f)

	f.沃夫拉Alwafra = BAlwafra沃夫拉
	f.沃夫拉Alwafra.SetParent(f)

	f.安塞穆夏斯Anthemusias = BAnthemusias安塞穆夏斯
	f.安塞穆夏斯Anthemusias.SetParent(f)

	f.劳扎塔因Arrawdatayn = BArrawdatayn劳扎塔因
	f.劳扎塔因Arrawdatayn.SetParent(f)

	f.舒艾拜Ashshuaybah = BAshshuaybah舒艾拜
	f.舒艾拜Ashshuaybah.SetParent(f)

	f.伊卡洛斯Ikaros = BIkaros伊卡洛斯
	f.伊卡洛斯Ikaros.SetParent(f)

	f.希兰Khinan = BKhinan希兰
	f.希兰Khinan.SetParent(f)

	f.科威特Kuwait = BKuwait科威特
	f.科威特Kuwait.SetParent(f)

}
