package sistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SistanCounty interface {
	feud.County
	BAdar阿达尔() feud.Barony
	BDehak代哈克() feud.Barony
	BEsfandak埃斯凡达克() feud.Barony
	BHaozdar豪兹达() feud.Barony
	BShahresukhteh沙赫尔苏赫特() feud.Barony
}

type 古夫斯山SistanCounty struct {
	feud.BaseCounty
	阿达尔Adar             feud.Barony
	代哈克Dehak            feud.Barony
	埃斯凡达克Esfandak       feud.Barony
	豪兹达Haozdar          feud.Barony
	沙赫尔苏赫特Shahresukhteh feud.Barony
}

func (c *古夫斯山SistanCounty) BAdar阿达尔() feud.Barony {
	return c.阿达尔Adar
}

func (c *古夫斯山SistanCounty) BDehak代哈克() feud.Barony {
	return c.代哈克Dehak
}

func (c *古夫斯山SistanCounty) BEsfandak埃斯凡达克() feud.Barony {
	return c.埃斯凡达克Esfandak
}

func (c *古夫斯山SistanCounty) BHaozdar豪兹达() feud.Barony {
	return c.豪兹达Haozdar
}

func (c *古夫斯山SistanCounty) BShahresukhteh沙赫尔苏赫特() feud.Barony {
	return c.沙赫尔苏赫特Shahresukhteh
}

var CSistan古夫斯山 SistanCounty = &古夫斯山SistanCounty{}

func init() {
	f := CSistan古夫斯山.(*古夫斯山SistanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "637",
		Title:     "sistan",
		TitleName: "古夫斯山",
		TitleCode: "c_sistan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿达尔Adar = BAdar阿达尔
	f.阿达尔Adar.SetParent(f)

	f.代哈克Dehak = BDehak代哈克
	f.代哈克Dehak.SetParent(f)

	f.埃斯凡达克Esfandak = BEsfandak埃斯凡达克
	f.埃斯凡达克Esfandak.SetParent(f)

	f.豪兹达Haozdar = BHaozdar豪兹达
	f.豪兹达Haozdar.SetParent(f)

	f.沙赫尔苏赫特Shahresukhteh = BShahresukhteh沙赫尔苏赫特
	f.沙赫尔苏赫特Shahresukhteh.SetParent(f)

}
