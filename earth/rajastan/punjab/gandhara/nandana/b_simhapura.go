package nandana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 狮子城SimhapuraBarony struct {
	feud.BaseBarony
}

var BSimhapura狮子城 feud.Barony = &狮子城SimhapuraBarony{}

func init() {
	f := BSimhapura狮子城.(*狮子城SimhapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "simhapura",
		TitleName: "狮子城",
		TitleCode: "b_simhapura",
	}
}
