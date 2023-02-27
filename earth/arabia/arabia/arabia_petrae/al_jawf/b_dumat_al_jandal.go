package al_jawf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜玛占达Dumat_al_jandalBarony struct {
	feud.BaseBarony
}

var BDumat_al_jandal杜玛占达 feud.Barony = &杜玛占达Dumat_al_jandalBarony{}

func init() {
    f := BDumat_al_jandal杜玛占达.(*杜玛占达Dumat_al_jandalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dumat_al_jandal",
		TitleName: "杜玛占达",
		TitleCode: "b_dumat_al_jandal",
	}
}
