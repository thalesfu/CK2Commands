package plauen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茨维考ZwickauBarony struct {
	feud.BaseBarony
}

var BZwickau茨维考 feud.Barony = &茨维考ZwickauBarony{}

func init() {
    f := BZwickau茨维考.(*茨维考ZwickauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zwickau",
		TitleName: "茨维考",
		TitleCode: "b_zwickau",
	}
}
