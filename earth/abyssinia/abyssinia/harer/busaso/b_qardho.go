package busaso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加尔多QardhoBarony struct {
	feud.BaseBarony
}

var BQardho加尔多 feud.Barony = &加尔多QardhoBarony{}

func init() {
	f := BQardho加尔多.(*加尔多QardhoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qardho",
		TitleName: "加尔多",
		TitleCode: "b_qardho",
	}
}
