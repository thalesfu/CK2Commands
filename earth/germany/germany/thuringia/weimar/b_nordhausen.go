package weimar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺德豪森NordhausenBarony struct {
	feud.BaseBarony
}

var BNordhausen诺德豪森 feud.Barony = &诺德豪森NordhausenBarony{}

func init() {
    f := BNordhausen诺德豪森.(*诺德豪森NordhausenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nordhausen",
		TitleName: "诺德豪森",
		TitleCode: "b_nordhausen",
	}
}
