package tlemcen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TlemcenCounty interface {
	feud.County
	BBedrabine贝德拉宾() feud.Barony
	BDhaya扎耶() feud.Barony
	BElhaicaiba哈凯巴() feud.Barony
	BLetelagh泰拉格() feud.Barony
	BMuaskar穆阿斯凯尔() feud.Barony
	BSidibelabbes西迪贝勒阿贝斯() feud.Barony
	BTessala特斯萨拉() feud.Barony
	BTlemcen特莱姆森() feud.Barony
}

type 特莱姆森TlemcenCounty struct {
	feud.BaseCounty
	贝德拉宾Bedrabine       feud.Barony
	扎耶Dhaya             feud.Barony
	哈凯巴Elhaicaiba       feud.Barony
	泰拉格Letelagh         feud.Barony
	穆阿斯凯尔Muaskar        feud.Barony
	西迪贝勒阿贝斯Sidibelabbes feud.Barony
	特斯萨拉Tessala         feud.Barony
	特莱姆森Tlemcen         feud.Barony
}

func (c *特莱姆森TlemcenCounty) BBedrabine贝德拉宾() feud.Barony {
	return c.贝德拉宾Bedrabine
}

func (c *特莱姆森TlemcenCounty) BDhaya扎耶() feud.Barony {
	return c.扎耶Dhaya
}

func (c *特莱姆森TlemcenCounty) BElhaicaiba哈凯巴() feud.Barony {
	return c.哈凯巴Elhaicaiba
}

func (c *特莱姆森TlemcenCounty) BLetelagh泰拉格() feud.Barony {
	return c.泰拉格Letelagh
}

func (c *特莱姆森TlemcenCounty) BMuaskar穆阿斯凯尔() feud.Barony {
	return c.穆阿斯凯尔Muaskar
}

func (c *特莱姆森TlemcenCounty) BSidibelabbes西迪贝勒阿贝斯() feud.Barony {
	return c.西迪贝勒阿贝斯Sidibelabbes
}

func (c *特莱姆森TlemcenCounty) BTessala特斯萨拉() feud.Barony {
	return c.特斯萨拉Tessala
}

func (c *特莱姆森TlemcenCounty) BTlemcen特莱姆森() feud.Barony {
	return c.特莱姆森Tlemcen
}

var CTlemcen特莱姆森 TlemcenCounty = &特莱姆森TlemcenCounty{}

func init() {
	f := CTlemcen特莱姆森.(*特莱姆森TlemcenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "834",
		Title:     "tlemcen",
		TitleName: "特莱姆森",
		TitleCode: "c_tlemcen",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝德拉宾Bedrabine = BBedrabine贝德拉宾
	f.贝德拉宾Bedrabine.SetParent(f)

	f.扎耶Dhaya = BDhaya扎耶
	f.扎耶Dhaya.SetParent(f)

	f.哈凯巴Elhaicaiba = BElhaicaiba哈凯巴
	f.哈凯巴Elhaicaiba.SetParent(f)

	f.泰拉格Letelagh = BLetelagh泰拉格
	f.泰拉格Letelagh.SetParent(f)

	f.穆阿斯凯尔Muaskar = BMuaskar穆阿斯凯尔
	f.穆阿斯凯尔Muaskar.SetParent(f)

	f.西迪贝勒阿贝斯Sidibelabbes = BSidibelabbes西迪贝勒阿贝斯
	f.西迪贝勒阿贝斯Sidibelabbes.SetParent(f)

	f.特斯萨拉Tessala = BTessala特斯萨拉
	f.特斯萨拉Tessala.SetParent(f)

	f.特莱姆森Tlemcen = BTlemcen特莱姆森
	f.特莱姆森Tlemcen.SetParent(f)

}
