package damman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈季迈NajmahBarony struct {
	feud.BaseBarony
}

var BNajmah奈季迈 feud.Barony = &奈季迈NajmahBarony{}

func init() {
	f := BNajmah奈季迈.(*奈季迈NajmahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "najmah",
		TitleName: "奈季迈",
		TitleCode: "b_najmah",
	}
}
