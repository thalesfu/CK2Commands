package besancon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔绍MarchauxBarony struct {
	feud.BaseBarony
}

var BMarchaux马尔绍 feud.Barony = &马尔绍MarchauxBarony{}

func init() {
    f := BMarchaux马尔绍.(*马尔绍MarchauxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marchaux",
		TitleName: "马尔绍",
		TitleCode: "b_marchaux",
	}
}
