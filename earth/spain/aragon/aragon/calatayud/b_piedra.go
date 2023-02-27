package calatayud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼德拉PiedraBarony struct {
	feud.BaseBarony
}

var BPiedra彼德拉 feud.Barony = &彼德拉PiedraBarony{}

func init() {
    f := BPiedra彼德拉.(*彼德拉PiedraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piedra",
		TitleName: "彼德拉",
		TitleCode: "b_piedra",
	}
}
