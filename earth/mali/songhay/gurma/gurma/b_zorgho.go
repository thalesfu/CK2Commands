package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佐尔戈ZorghoBarony struct {
	feud.BaseBarony
}

var BZorgho佐尔戈 feud.Barony = &佐尔戈ZorghoBarony{}

func init() {
	f := BZorgho佐尔戈.(*佐尔戈ZorghoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zorgho",
		TitleName: "佐尔戈",
		TitleCode: "b_zorgho",
	}
}
