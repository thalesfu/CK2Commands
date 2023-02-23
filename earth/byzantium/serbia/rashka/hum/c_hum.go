package hum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HumCounty interface {
	feud.County
	BBelacrkva贝拉茨尔克瓦() feud.Barony
	BBradarevo布罗达雷沃() feud.Barony
	BMedun梅敦() feud.Barony
	BMoraca莫拉查() feud.Barony
	BNovipazar特尔戈维什泰() feud.Barony
	BPec佩奇() feud.Barony
	BPlav普拉夫() feud.Barony
	BStupovi斯图波维() feud.Barony
}

type 拉斯HumCounty struct {
	feud.BaseCounty
	贝拉茨尔克瓦Belacrkva feud.Barony
	布罗达雷沃Bradarevo  feud.Barony
	梅敦Medun         feud.Barony
	莫拉查Moraca       feud.Barony
	特尔戈维什泰Novipazar feud.Barony
	佩奇Pec           feud.Barony
	普拉夫Plav         feud.Barony
	斯图波维Stupovi     feud.Barony
}

func (c *拉斯HumCounty) BBelacrkva贝拉茨尔克瓦() feud.Barony {
	return c.贝拉茨尔克瓦Belacrkva
}

func (c *拉斯HumCounty) BBradarevo布罗达雷沃() feud.Barony {
	return c.布罗达雷沃Bradarevo
}

func (c *拉斯HumCounty) BMedun梅敦() feud.Barony {
	return c.梅敦Medun
}

func (c *拉斯HumCounty) BMoraca莫拉查() feud.Barony {
	return c.莫拉查Moraca
}

func (c *拉斯HumCounty) BNovipazar特尔戈维什泰() feud.Barony {
	return c.特尔戈维什泰Novipazar
}

func (c *拉斯HumCounty) BPec佩奇() feud.Barony {
	return c.佩奇Pec
}

func (c *拉斯HumCounty) BPlav普拉夫() feud.Barony {
	return c.普拉夫Plav
}

func (c *拉斯HumCounty) BStupovi斯图波维() feud.Barony {
	return c.斯图波维Stupovi
}

var CHum拉斯 HumCounty = &拉斯HumCounty{}

func init() {
	f := CHum拉斯.(*拉斯HumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "503",
		Title:     "hum",
		TitleName: "拉斯",
		TitleCode: "c_hum",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝拉茨尔克瓦Belacrkva = BBelacrkva贝拉茨尔克瓦
	f.贝拉茨尔克瓦Belacrkva.SetParent(f)

	f.布罗达雷沃Bradarevo = BBradarevo布罗达雷沃
	f.布罗达雷沃Bradarevo.SetParent(f)

	f.梅敦Medun = BMedun梅敦
	f.梅敦Medun.SetParent(f)

	f.莫拉查Moraca = BMoraca莫拉查
	f.莫拉查Moraca.SetParent(f)

	f.特尔戈维什泰Novipazar = BNovipazar特尔戈维什泰
	f.特尔戈维什泰Novipazar.SetParent(f)

	f.佩奇Pec = BPec佩奇
	f.佩奇Pec.SetParent(f)

	f.普拉夫Plav = BPlav普拉夫
	f.普拉夫Plav.SetParent(f)

	f.斯图波维Stupovi = BStupovi斯图波维
	f.斯图波维Stupovi.SetParent(f)

}
