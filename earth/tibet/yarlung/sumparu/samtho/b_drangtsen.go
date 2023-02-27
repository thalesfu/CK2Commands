package samtho

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 长颈DrangtsenBarony struct {
	feud.BaseBarony
}

var BDrangtsen长颈 feud.Barony = &长颈DrangtsenBarony{}

func init() {
    f := BDrangtsen长颈.(*长颈DrangtsenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drangtsen",
		TitleName: "长颈",
		TitleCode: "b_drangtsen",
	}
}
