package erchis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 也儿的石ErtisBarony struct {
	feud.BaseBarony
}

var BErtis也儿的石 feud.Barony = &也儿的石ErtisBarony{}

func init() {
	f := BErtis也儿的石.(*也儿的石ErtisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ertis",
		TitleName: "也儿的石",
		TitleCode: "b_ertis",
	}
}
