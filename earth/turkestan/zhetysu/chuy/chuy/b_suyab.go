package chuy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 碎叶SuyabBarony struct {
	feud.BaseBarony
}

var BSuyab碎叶 feud.Barony = &碎叶SuyabBarony{}

func init() {
	f := BSuyab碎叶.(*碎叶SuyabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suyab",
		TitleName: "碎叶",
		TitleCode: "b_suyab",
	}
}
