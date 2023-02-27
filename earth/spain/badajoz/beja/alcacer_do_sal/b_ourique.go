package alcacer_do_sal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧里基OuriqueBarony struct {
	feud.BaseBarony
}

var BOurique欧里基 feud.Barony = &欧里基OuriqueBarony{}

func init() {
    f := BOurique欧里基.(*欧里基OuriqueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ourique",
		TitleName: "欧里基",
		TitleCode: "b_ourique",
	}
}
