package pettau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比斯特里察BistricaBarony struct {
	feud.BaseBarony
}

var BBistrica比斯特里察 feud.Barony = &比斯特里察BistricaBarony{}

func init() {
	f := BBistrica比斯特里察.(*比斯特里察BistricaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bistrica",
		TitleName: "比斯特里察",
		TitleCode: "b_bistrica",
	}
}
