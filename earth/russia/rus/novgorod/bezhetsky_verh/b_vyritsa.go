package bezhetsky_verh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维里察VyritsaBarony struct {
	feud.BaseBarony
}

var BVyritsa维里察 feud.Barony = &维里察VyritsaBarony{}

func init() {
    f := BVyritsa维里察.(*维里察VyritsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyritsa",
		TitleName: "维里察",
		TitleCode: "b_vyritsa",
	}
}
