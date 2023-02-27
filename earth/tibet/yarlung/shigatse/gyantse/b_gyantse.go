package gyantse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 江孜GyantseBarony struct {
	feud.BaseBarony
}

var BGyantse江孜 feud.Barony = &江孜GyantseBarony{}

func init() {
    f := BGyantse江孜.(*江孜GyantseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyantse",
		TitleName: "江孜",
		TitleCode: "b_gyantse",
	}
}
