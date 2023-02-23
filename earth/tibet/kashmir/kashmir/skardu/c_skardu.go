package skardu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SkarduCounty interface {
	feud.County
	BHilmat希尔迈特() feud.Barony
	BHuldi哈尔迪() feud.Barony
	BJanwai剑苇() feud.Barony
	BKhaplu贺菩劳() feud.Barony
	BKharkoo喀尔丘() feud.Barony
	BSkardu锡卡都() feud.Barony
}

type 锡卡都SkarduCounty struct {
	feud.BaseCounty
	希尔迈特Hilmat feud.Barony
	哈尔迪Huldi   feud.Barony
	剑苇Janwai   feud.Barony
	贺菩劳Khaplu  feud.Barony
	喀尔丘Kharkoo feud.Barony
	锡卡都Skardu  feud.Barony
}

func (c *锡卡都SkarduCounty) BHilmat希尔迈特() feud.Barony {
	return c.希尔迈特Hilmat
}

func (c *锡卡都SkarduCounty) BHuldi哈尔迪() feud.Barony {
	return c.哈尔迪Huldi
}

func (c *锡卡都SkarduCounty) BJanwai剑苇() feud.Barony {
	return c.剑苇Janwai
}

func (c *锡卡都SkarduCounty) BKhaplu贺菩劳() feud.Barony {
	return c.贺菩劳Khaplu
}

func (c *锡卡都SkarduCounty) BKharkoo喀尔丘() feud.Barony {
	return c.喀尔丘Kharkoo
}

func (c *锡卡都SkarduCounty) BSkardu锡卡都() feud.Barony {
	return c.锡卡都Skardu
}

var CSkardu锡卡都 SkarduCounty = &锡卡都SkarduCounty{}

func init() {
	f := CSkardu锡卡都.(*锡卡都SkarduCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1466",
		Title:     "skardu",
		TitleName: "锡卡都",
		TitleCode: "c_skardu",
		Baronies:  map[string]feud.Barony{},
	}

	f.希尔迈特Hilmat = BHilmat希尔迈特
	f.希尔迈特Hilmat.SetParent(f)

	f.哈尔迪Huldi = BHuldi哈尔迪
	f.哈尔迪Huldi.SetParent(f)

	f.剑苇Janwai = BJanwai剑苇
	f.剑苇Janwai.SetParent(f)

	f.贺菩劳Khaplu = BKhaplu贺菩劳
	f.贺菩劳Khaplu.SetParent(f)

	f.喀尔丘Kharkoo = BKharkoo喀尔丘
	f.喀尔丘Kharkoo.SetParent(f)

	f.锡卡都Skardu = BSkardu锡卡都
	f.锡卡都Skardu.SetParent(f)

}
