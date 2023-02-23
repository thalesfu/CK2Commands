package fife

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柯科迪KirkcaldyBarony struct {
	feud.BaseBarony
}

var BKirkcaldy柯科迪 feud.Barony = &柯科迪KirkcaldyBarony{}

func init() {
	f := BKirkcaldy柯科迪.(*柯科迪KirkcaldyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirkcaldy",
		TitleName: "柯科迪",
		TitleCode: "b_kirkcaldy",
	}
}
