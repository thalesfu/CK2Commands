package herakleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫拉克利亚HerakleiaBarony struct {
	feud.BaseBarony
}

var BHerakleia赫拉克利亚 feud.Barony = &赫拉克利亚HerakleiaBarony{}

func init() {
	f := BHerakleia赫拉克利亚.(*赫拉克利亚HerakleiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "herakleia",
		TitleName: "赫拉克利亚",
		TitleCode: "b_herakleia",
	}
}
