package lower_dniepr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥列什基OleshyeBarony struct {
	feud.BaseBarony
}

var BOleshye奥列什基 feud.Barony = &奥列什基OleshyeBarony{}

func init() {
    f := BOleshye奥列什基.(*奥列什基OleshyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oleshye",
		TitleName: "奥列什基",
		TitleCode: "b_oleshye",
	}
}
