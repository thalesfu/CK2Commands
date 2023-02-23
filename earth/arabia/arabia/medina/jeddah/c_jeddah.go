package jeddah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JeddahCounty interface {
	feud.County
	BBahrah拜赫拉() feud.Barony
	BJeddah吉达() feud.Barony
	BKhulays胡莱斯() feud.Barony
	BQunfudhah贡费宰() feud.Barony
	BUsfan欧斯凡() feud.Barony
}

type 吉达JeddahCounty struct {
	feud.BaseCounty
	拜赫拉Bahrah    feud.Barony
	吉达Jeddah     feud.Barony
	胡莱斯Khulays   feud.Barony
	贡费宰Qunfudhah feud.Barony
	欧斯凡Usfan     feud.Barony
}

func (c *吉达JeddahCounty) BBahrah拜赫拉() feud.Barony {
	return c.拜赫拉Bahrah
}

func (c *吉达JeddahCounty) BJeddah吉达() feud.Barony {
	return c.吉达Jeddah
}

func (c *吉达JeddahCounty) BKhulays胡莱斯() feud.Barony {
	return c.胡莱斯Khulays
}

func (c *吉达JeddahCounty) BQunfudhah贡费宰() feud.Barony {
	return c.贡费宰Qunfudhah
}

func (c *吉达JeddahCounty) BUsfan欧斯凡() feud.Barony {
	return c.欧斯凡Usfan
}

var CJeddah吉达 JeddahCounty = &吉达JeddahCounty{}

func init() {
	f := CJeddah吉达.(*吉达JeddahCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "716",
		Title:     "jeddah",
		TitleName: "吉达",
		TitleCode: "c_jeddah",
		Baronies:  map[string]feud.Barony{},
	}

	f.拜赫拉Bahrah = BBahrah拜赫拉
	f.拜赫拉Bahrah.SetParent(f)

	f.吉达Jeddah = BJeddah吉达
	f.吉达Jeddah.SetParent(f)

	f.胡莱斯Khulays = BKhulays胡莱斯
	f.胡莱斯Khulays.SetParent(f)

	f.贡费宰Qunfudhah = BQunfudhah贡费宰
	f.贡费宰Qunfudhah.SetParent(f)

	f.欧斯凡Usfan = BUsfan欧斯凡
	f.欧斯凡Usfan.SetParent(f)

}
