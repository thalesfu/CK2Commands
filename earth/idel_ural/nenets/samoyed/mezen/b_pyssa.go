package mezen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩萨PyssaBarony struct {
	feud.BaseBarony
}

var BPyssa佩萨 feud.Barony = &佩萨PyssaBarony{}

func init() {
    f := BPyssa佩萨.(*佩萨PyssaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pyssa",
		TitleName: "佩萨",
		TitleCode: "b_pyssa",
	}
}
