package alcacer_do_sal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥德米拉OdemiraBarony struct {
	feud.BaseBarony
}

var BOdemira奥德米拉 feud.Barony = &奥德米拉OdemiraBarony{}

func init() {
    f := BOdemira奥德米拉.(*奥德米拉OdemiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odemira",
		TitleName: "奥德米拉",
		TitleCode: "b_odemira",
	}
}
