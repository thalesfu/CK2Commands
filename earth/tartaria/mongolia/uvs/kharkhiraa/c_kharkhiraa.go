package kharkhiraa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KharkhiraaCounty interface {
	feud.County
	BBaatar巴特尔() feud.Barony
	BBumbag奔巴图() feud.Barony
	BKharkhiraa哈尔齐拉() feud.Barony
	BSutay苏泰() feud.Barony
	BTsetseg车车格() feud.Barony
}

type 哈尔齐拉KharkhiraaCounty struct {
	feud.BaseCounty
	巴特尔Baatar      feud.Barony
	奔巴图Bumbag      feud.Barony
	哈尔齐拉Kharkhiraa feud.Barony
	苏泰Sutay        feud.Barony
	车车格Tsetseg     feud.Barony
}

func (c *哈尔齐拉KharkhiraaCounty) BBaatar巴特尔() feud.Barony {
	return c.巴特尔Baatar
}

func (c *哈尔齐拉KharkhiraaCounty) BBumbag奔巴图() feud.Barony {
	return c.奔巴图Bumbag
}

func (c *哈尔齐拉KharkhiraaCounty) BKharkhiraa哈尔齐拉() feud.Barony {
	return c.哈尔齐拉Kharkhiraa
}

func (c *哈尔齐拉KharkhiraaCounty) BSutay苏泰() feud.Barony {
	return c.苏泰Sutay
}

func (c *哈尔齐拉KharkhiraaCounty) BTsetseg车车格() feud.Barony {
	return c.车车格Tsetseg
}

var CKharkhiraa哈尔齐拉 KharkhiraaCounty = &哈尔齐拉KharkhiraaCounty{}

func init() {
	f := CKharkhiraa哈尔齐拉.(*哈尔齐拉KharkhiraaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1907",
		Title:     "kharkhiraa",
		TitleName: "哈尔齐拉",
		TitleCode: "c_kharkhiraa",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴特尔Baatar = BBaatar巴特尔
	f.巴特尔Baatar.SetParent(f)

	f.奔巴图Bumbag = BBumbag奔巴图
	f.奔巴图Bumbag.SetParent(f)

	f.哈尔齐拉Kharkhiraa = BKharkhiraa哈尔齐拉
	f.哈尔齐拉Kharkhiraa.SetParent(f)

	f.苏泰Sutay = BSutay苏泰
	f.苏泰Sutay.SetParent(f)

	f.车车格Tsetseg = BTsetseg车车格
	f.车车格Tsetseg.SetParent(f)

}
