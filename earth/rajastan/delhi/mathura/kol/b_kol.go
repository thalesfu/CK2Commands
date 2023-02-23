package kol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘罗KolBarony struct {
	feud.BaseBarony
}

var BKol拘罗 feud.Barony = &拘罗KolBarony{}

func init() {
	f := BKol拘罗.(*拘罗KolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kol",
		TitleName: "拘罗",
		TitleCode: "b_kol",
	}
}
