package bahrein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BahreinCounty interface {
	feud.County
	BMurwab穆尔瓦比() feud.Barony
	BTuraynah图赖纳() feud.Barony
	BZubarah祖巴拉() feud.Barony
}

type 巴林BahreinCounty struct {
	feud.BaseCounty
	穆尔瓦比Murwab  feud.Barony
	图赖纳Turaynah feud.Barony
	祖巴拉Zubarah  feud.Barony
}

func (c *巴林BahreinCounty) BMurwab穆尔瓦比() feud.Barony {
	return c.穆尔瓦比Murwab
}

func (c *巴林BahreinCounty) BTuraynah图赖纳() feud.Barony {
	return c.图赖纳Turaynah
}

func (c *巴林BahreinCounty) BZubarah祖巴拉() feud.Barony {
	return c.祖巴拉Zubarah
}

var CBahrein巴林 BahreinCounty = &巴林BahreinCounty{}

func init() {
	f := CBahrein巴林.(*巴林BahreinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "653",
		Title:     "bahrein",
		TitleName: "巴林",
		TitleCode: "c_bahrein",
		Baronies:  map[string]feud.Barony{},
	}

	f.穆尔瓦比Murwab = BMurwab穆尔瓦比
	f.穆尔瓦比Murwab.SetParent(f)

	f.图赖纳Turaynah = BTuraynah图赖纳
	f.图赖纳Turaynah.SetParent(f)

	f.祖巴拉Zubarah = BZubarah祖巴拉
	f.祖巴拉Zubarah.SetParent(f)

}
