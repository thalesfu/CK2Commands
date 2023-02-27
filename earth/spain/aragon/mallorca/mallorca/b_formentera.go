package mallorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福门特拉FormenteraBarony struct {
	feud.BaseBarony
}

var BFormentera福门特拉 feud.Barony = &福门特拉FormenteraBarony{}

func init() {
    f := BFormentera福门特拉.(*福门特拉FormenteraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "formentera",
		TitleName: "福门特拉",
		TitleCode: "b_formentera",
	}
}
