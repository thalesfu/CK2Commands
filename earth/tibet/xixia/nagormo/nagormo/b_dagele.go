package nagormo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大格勒DageleBarony struct {
	feud.BaseBarony
}

var BDagele大格勒 feud.Barony = &大格勒DageleBarony{}

func init() {
	f := BDagele大格勒.(*大格勒DageleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dagele",
		TitleName: "大格勒",
		TitleCode: "b_dagele",
	}
}
