package suf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SufCounty interface {
	feud.County
	BBidane比当() feud.Barony
	BFatnassa法特纳萨() feud.Barony
	BSuf苏夫() feud.Barony
}

type 苏夫SufCounty struct {
	feud.BaseCounty
	比当Bidane     feud.Barony
	法特纳萨Fatnassa feud.Barony
	苏夫Suf        feud.Barony
}

func (c *苏夫SufCounty) BBidane比当() feud.Barony {
	return c.比当Bidane
}

func (c *苏夫SufCounty) BFatnassa法特纳萨() feud.Barony {
	return c.法特纳萨Fatnassa
}

func (c *苏夫SufCounty) BSuf苏夫() feud.Barony {
	return c.苏夫Suf
}

var CSuf苏夫 SufCounty = &苏夫SufCounty{}

func init() {
	f := CSuf苏夫.(*苏夫SufCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1725",
		Title:     "suf",
		TitleName: "苏夫",
		TitleCode: "c_suf",
		Baronies:  map[string]feud.Barony{},
	}

	f.比当Bidane = BBidane比当
	f.比当Bidane.SetParent(f)

	f.法特纳萨Fatnassa = BFatnassa法特纳萨
	f.法特纳萨Fatnassa.SetParent(f)

	f.苏夫Suf = BSuf苏夫
	f.苏夫Suf.SetParent(f)

}
