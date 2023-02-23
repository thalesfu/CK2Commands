package khantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhantiaCounty interface {
	feud.County
	BBeloyarskiy别洛亚尔() feud.Barony
	BBerezovo别廖佐沃() feud.Barony
	BChanty汉特() feud.Barony
	BDjinesh吉涅什() feud.Barony
	BIgrim伊格里姆() feud.Barony
	BNyagyn尼亚甘() feud.Barony
	BPnobe普诺比耶() feud.Barony
	BSherkala舍尔卡拉() feud.Barony
}

type 汉特KhantiaCounty struct {
	feud.BaseCounty
	别洛亚尔Beloyarskiy feud.Barony
	别廖佐沃Berezovo    feud.Barony
	汉特Chanty        feud.Barony
	吉涅什Djinesh      feud.Barony
	伊格里姆Igrim       feud.Barony
	尼亚甘Nyagyn       feud.Barony
	普诺比耶Pnobe       feud.Barony
	舍尔卡拉Sherkala    feud.Barony
}

func (c *汉特KhantiaCounty) BBeloyarskiy别洛亚尔() feud.Barony {
	return c.别洛亚尔Beloyarskiy
}

func (c *汉特KhantiaCounty) BBerezovo别廖佐沃() feud.Barony {
	return c.别廖佐沃Berezovo
}

func (c *汉特KhantiaCounty) BChanty汉特() feud.Barony {
	return c.汉特Chanty
}

func (c *汉特KhantiaCounty) BDjinesh吉涅什() feud.Barony {
	return c.吉涅什Djinesh
}

func (c *汉特KhantiaCounty) BIgrim伊格里姆() feud.Barony {
	return c.伊格里姆Igrim
}

func (c *汉特KhantiaCounty) BNyagyn尼亚甘() feud.Barony {
	return c.尼亚甘Nyagyn
}

func (c *汉特KhantiaCounty) BPnobe普诺比耶() feud.Barony {
	return c.普诺比耶Pnobe
}

func (c *汉特KhantiaCounty) BSherkala舍尔卡拉() feud.Barony {
	return c.舍尔卡拉Sherkala
}

var CKhantia汉特 KhantiaCounty = &汉特KhantiaCounty{}

func init() {
	f := CKhantia汉特.(*汉特KhantiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "890",
		Title:     "khantia",
		TitleName: "汉特",
		TitleCode: "c_khantia",
		Baronies:  map[string]feud.Barony{},
	}

	f.别洛亚尔Beloyarskiy = BBeloyarskiy别洛亚尔
	f.别洛亚尔Beloyarskiy.SetParent(f)

	f.别廖佐沃Berezovo = BBerezovo别廖佐沃
	f.别廖佐沃Berezovo.SetParent(f)

	f.汉特Chanty = BChanty汉特
	f.汉特Chanty.SetParent(f)

	f.吉涅什Djinesh = BDjinesh吉涅什
	f.吉涅什Djinesh.SetParent(f)

	f.伊格里姆Igrim = BIgrim伊格里姆
	f.伊格里姆Igrim.SetParent(f)

	f.尼亚甘Nyagyn = BNyagyn尼亚甘
	f.尼亚甘Nyagyn.SetParent(f)

	f.普诺比耶Pnobe = BPnobe普诺比耶
	f.普诺比耶Pnobe.SetParent(f)

	f.舍尔卡拉Sherkala = BSherkala舍尔卡拉
	f.舍尔卡拉Sherkala.SetParent(f)

}
