package nellore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝利补兰多迦摩TripurantakamBarony struct {
	feud.BaseBarony
}

var BTripurantakam帝利补兰多迦摩 feud.Barony = &帝利补兰多迦摩TripurantakamBarony{}

func init() {
	f := BTripurantakam帝利补兰多迦摩.(*帝利补兰多迦摩TripurantakamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tripurantakam",
		TitleName: "帝利补兰多迦摩",
		TitleCode: "b_tripurantakam",
	}
}
