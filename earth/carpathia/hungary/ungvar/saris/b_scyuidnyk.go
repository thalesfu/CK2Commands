package saris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯奎德尼克ScyuidnykBarony struct {
	feud.BaseBarony
}

var BScyuidnyk斯奎德尼克 feud.Barony = &斯奎德尼克ScyuidnykBarony{}

func init() {
	f := BScyuidnyk斯奎德尼克.(*斯奎德尼克ScyuidnykBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "scyuidnyk",
		TitleName: "斯奎德尼克",
		TitleCode: "b_scyuidnyk",
	}
}
