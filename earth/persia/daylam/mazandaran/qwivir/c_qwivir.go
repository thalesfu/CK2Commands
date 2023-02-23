package qwivir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QwivirCounty interface {
	feud.County
	BDamqan达姆甘() feud.Barony
	BDarvar达夫瓦() feud.Barony
	BDehnamak德纳马克() feud.Barony
	BGerdkuh吉尔德库赫() feud.Barony
	BKohandej科汉德() feud.Barony
	BSangsar桑格萨尔() feud.Barony
	BSemnan塞姆南() feud.Barony
	BSharequmis沙赫尔古米斯() feud.Barony
}

type 奎维尔QwivirCounty struct {
	feud.BaseCounty
	达姆甘Damqan        feud.Barony
	达夫瓦Darvar        feud.Barony
	德纳马克Dehnamak     feud.Barony
	吉尔德库赫Gerdkuh     feud.Barony
	科汉德Kohandej      feud.Barony
	桑格萨尔Sangsar      feud.Barony
	塞姆南Semnan        feud.Barony
	沙赫尔古米斯Sharequmis feud.Barony
}

func (c *奎维尔QwivirCounty) BDamqan达姆甘() feud.Barony {
	return c.达姆甘Damqan
}

func (c *奎维尔QwivirCounty) BDarvar达夫瓦() feud.Barony {
	return c.达夫瓦Darvar
}

func (c *奎维尔QwivirCounty) BDehnamak德纳马克() feud.Barony {
	return c.德纳马克Dehnamak
}

func (c *奎维尔QwivirCounty) BGerdkuh吉尔德库赫() feud.Barony {
	return c.吉尔德库赫Gerdkuh
}

func (c *奎维尔QwivirCounty) BKohandej科汉德() feud.Barony {
	return c.科汉德Kohandej
}

func (c *奎维尔QwivirCounty) BSangsar桑格萨尔() feud.Barony {
	return c.桑格萨尔Sangsar
}

func (c *奎维尔QwivirCounty) BSemnan塞姆南() feud.Barony {
	return c.塞姆南Semnan
}

func (c *奎维尔QwivirCounty) BSharequmis沙赫尔古米斯() feud.Barony {
	return c.沙赫尔古米斯Sharequmis
}

var CQwivir奎维尔 QwivirCounty = &奎维尔QwivirCounty{}

func init() {
	f := CQwivir奎维尔.(*奎维尔QwivirCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "660",
		Title:     "qwivir",
		TitleName: "奎维尔",
		TitleCode: "c_qwivir",
		Baronies:  map[string]feud.Barony{},
	}

	f.达姆甘Damqan = BDamqan达姆甘
	f.达姆甘Damqan.SetParent(f)

	f.达夫瓦Darvar = BDarvar达夫瓦
	f.达夫瓦Darvar.SetParent(f)

	f.德纳马克Dehnamak = BDehnamak德纳马克
	f.德纳马克Dehnamak.SetParent(f)

	f.吉尔德库赫Gerdkuh = BGerdkuh吉尔德库赫
	f.吉尔德库赫Gerdkuh.SetParent(f)

	f.科汉德Kohandej = BKohandej科汉德
	f.科汉德Kohandej.SetParent(f)

	f.桑格萨尔Sangsar = BSangsar桑格萨尔
	f.桑格萨尔Sangsar.SetParent(f)

	f.塞姆南Semnan = BSemnan塞姆南
	f.塞姆南Semnan.SetParent(f)

	f.沙赫尔古米斯Sharequmis = BSharequmis沙赫尔古米斯
	f.沙赫尔古米斯Sharequmis.SetParent(f)

}
