package akroinon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AkroinonCounty interface {
	feud.County
	BAkroinon阿克罗伊农() feud.Barony
	BCelenaea刻莱奈() feud.Barony
	BEumenia尤门尼亚() feud.Barony
}

type 阿克罗伊农AkroinonCounty struct {
	feud.BaseCounty
	阿克罗伊农Akroinon feud.Barony
	刻莱奈Celenaea   feud.Barony
	尤门尼亚Eumenia   feud.Barony
}

func (c *阿克罗伊农AkroinonCounty) BAkroinon阿克罗伊农() feud.Barony {
	return c.阿克罗伊农Akroinon
}

func (c *阿克罗伊农AkroinonCounty) BCelenaea刻莱奈() feud.Barony {
	return c.刻莱奈Celenaea
}

func (c *阿克罗伊农AkroinonCounty) BEumenia尤门尼亚() feud.Barony {
	return c.尤门尼亚Eumenia
}

var CAkroinon阿克罗伊农 AkroinonCounty = &阿克罗伊农AkroinonCounty{}

func init() {
	f := CAkroinon阿克罗伊农.(*阿克罗伊农AkroinonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1932",
		Title:     "akroinon",
		TitleName: "阿克罗伊农",
		TitleCode: "c_akroinon",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克罗伊农Akroinon = BAkroinon阿克罗伊农
	f.阿克罗伊农Akroinon.SetParent(f)

	f.刻莱奈Celenaea = BCelenaea刻莱奈
	f.刻莱奈Celenaea.SetParent(f)

	f.尤门尼亚Eumenia = BEumenia尤门尼亚
	f.尤门尼亚Eumenia.SetParent(f)

}
