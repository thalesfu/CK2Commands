package kebbi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KebbiCounty interface {
	feud.County
	BArgungu阿尔贡古() feud.Barony
	BDigira迪吉拉() feud.Barony
	BKebbi凯比() feud.Barony
	BSokoto索科托() feud.Barony
}

type 凯比KebbiCounty struct {
	feud.BaseCounty
	阿尔贡古Argungu feud.Barony
	迪吉拉Digira   feud.Barony
	凯比Kebbi     feud.Barony
	索科托Sokoto   feud.Barony
}

func (c *凯比KebbiCounty) BArgungu阿尔贡古() feud.Barony {
	return c.阿尔贡古Argungu
}

func (c *凯比KebbiCounty) BDigira迪吉拉() feud.Barony {
	return c.迪吉拉Digira
}

func (c *凯比KebbiCounty) BKebbi凯比() feud.Barony {
	return c.凯比Kebbi
}

func (c *凯比KebbiCounty) BSokoto索科托() feud.Barony {
	return c.索科托Sokoto
}

var CKebbi凯比 KebbiCounty = &凯比KebbiCounty{}

func init() {
	f := CKebbi凯比.(*凯比KebbiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1752",
		Title:     "kebbi",
		TitleName: "凯比",
		TitleCode: "c_kebbi",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔贡古Argungu = BArgungu阿尔贡古
	f.阿尔贡古Argungu.SetParent(f)

	f.迪吉拉Digira = BDigira迪吉拉
	f.迪吉拉Digira.SetParent(f)

	f.凯比Kebbi = BKebbi凯比
	f.凯比Kebbi.SetParent(f)

	f.索科托Sokoto = BSokoto索科托
	f.索科托Sokoto.SetParent(f)

}
