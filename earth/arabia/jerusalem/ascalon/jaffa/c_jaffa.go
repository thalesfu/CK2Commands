package jaffa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JaffaCounty interface {
	feud.County
	BArsuf阿尔苏夫() feud.Barony
	BBeitdejan贝特达甘() feud.Barony
	BIbelin伊贝林() feud.Barony
	BJaffa雅法() feud.Barony
	BLydda吕大() feud.Barony
	BQula库尔拉() feud.Barony
	BRamleh拉姆拉() feud.Barony
	BYazur亚祖尔() feud.Barony
}

type 雅法JaffaCounty struct {
	feud.BaseCounty
	阿尔苏夫Arsuf     feud.Barony
	贝特达甘Beitdejan feud.Barony
	伊贝林Ibelin     feud.Barony
	雅法Jaffa       feud.Barony
	吕大Lydda       feud.Barony
	库尔拉Qula       feud.Barony
	拉姆拉Ramleh     feud.Barony
	亚祖尔Yazur      feud.Barony
}

func (c *雅法JaffaCounty) BArsuf阿尔苏夫() feud.Barony {
	return c.阿尔苏夫Arsuf
}

func (c *雅法JaffaCounty) BBeitdejan贝特达甘() feud.Barony {
	return c.贝特达甘Beitdejan
}

func (c *雅法JaffaCounty) BIbelin伊贝林() feud.Barony {
	return c.伊贝林Ibelin
}

func (c *雅法JaffaCounty) BJaffa雅法() feud.Barony {
	return c.雅法Jaffa
}

func (c *雅法JaffaCounty) BLydda吕大() feud.Barony {
	return c.吕大Lydda
}

func (c *雅法JaffaCounty) BQula库尔拉() feud.Barony {
	return c.库尔拉Qula
}

func (c *雅法JaffaCounty) BRamleh拉姆拉() feud.Barony {
	return c.拉姆拉Ramleh
}

func (c *雅法JaffaCounty) BYazur亚祖尔() feud.Barony {
	return c.亚祖尔Yazur
}

var CJaffa雅法 JaffaCounty = &雅法JaffaCounty{}

func init() {
	f := CJaffa雅法.(*雅法JaffaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "775",
		Title:     "jaffa",
		TitleName: "雅法",
		TitleCode: "c_jaffa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔苏夫Arsuf = BArsuf阿尔苏夫
	f.阿尔苏夫Arsuf.SetParent(f)

	f.贝特达甘Beitdejan = BBeitdejan贝特达甘
	f.贝特达甘Beitdejan.SetParent(f)

	f.伊贝林Ibelin = BIbelin伊贝林
	f.伊贝林Ibelin.SetParent(f)

	f.雅法Jaffa = BJaffa雅法
	f.雅法Jaffa.SetParent(f)

	f.吕大Lydda = BLydda吕大
	f.吕大Lydda.SetParent(f)

	f.库尔拉Qula = BQula库尔拉
	f.库尔拉Qula.SetParent(f)

	f.拉姆拉Ramleh = BRamleh拉姆拉
	f.拉姆拉Ramleh.SetParent(f)

	f.亚祖尔Yazur = BYazur亚祖尔
	f.亚祖尔Yazur.SetParent(f)

}
