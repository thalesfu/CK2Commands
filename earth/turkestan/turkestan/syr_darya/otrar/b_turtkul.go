package otrar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔特库尔TurtkulBarony struct {
	feud.BaseBarony
}

var BTurtkul图尔特库尔 feud.Barony = &图尔特库尔TurtkulBarony{}

func init() {
    f := BTurtkul图尔特库尔.(*图尔特库尔TurtkulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turtkul",
		TitleName: "图尔特库尔",
		TitleCode: "b_turtkul",
	}
}
