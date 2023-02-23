package taranto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔西TursiBarony struct {
	feud.BaseBarony
}

var BTursi图尔西 feud.Barony = &图尔西TursiBarony{}

func init() {
	f := BTursi图尔西.(*图尔西TursiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tursi",
		TitleName: "图尔西",
		TitleCode: "b_tursi",
	}
}
