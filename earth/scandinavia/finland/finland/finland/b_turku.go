package finland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔库TurkuBarony struct {
	feud.BaseBarony
}

var BTurku图尔库 feud.Barony = &图尔库TurkuBarony{}

func init() {
    f := BTurku图尔库.(*图尔库TurkuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turku",
		TitleName: "图尔库",
		TitleCode: "b_turku",
	}
}
