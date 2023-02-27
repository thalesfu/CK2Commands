package selija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜贝纽DubenouBarony struct {
	feud.BaseBarony
}

var BDubenou杜贝纽 feud.Barony = &杜贝纽DubenouBarony{}

func init() {
    f := BDubenou杜贝纽.(*杜贝纽DubenouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dubenou",
		TitleName: "杜贝纽",
		TitleCode: "b_dubenou",
	}
}
