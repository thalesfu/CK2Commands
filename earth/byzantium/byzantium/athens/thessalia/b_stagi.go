package thessalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔基StagiBarony struct {
	feud.BaseBarony
}

var BStagi斯塔基 feud.Barony = &斯塔基StagiBarony{}

func init() {
    f := BStagi斯塔基.(*斯塔基StagiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stagi",
		TitleName: "斯塔基",
		TitleCode: "b_stagi",
	}
}
