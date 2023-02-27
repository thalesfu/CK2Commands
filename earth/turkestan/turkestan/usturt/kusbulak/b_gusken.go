package kusbulak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古斯肯GuskenBarony struct {
	feud.BaseBarony
}

var BGusken古斯肯 feud.Barony = &古斯肯GuskenBarony{}

func init() {
    f := BGusken古斯肯.(*古斯肯GuskenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gusken",
		TitleName: "古斯肯",
		TitleCode: "b_gusken",
	}
}
