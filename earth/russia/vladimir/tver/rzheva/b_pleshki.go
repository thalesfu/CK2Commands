package rzheva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普列什基PleshkiBarony struct {
	feud.BaseBarony
}

var BPleshki普列什基 feud.Barony = &普列什基PleshkiBarony{}

func init() {
	f := BPleshki普列什基.(*普列什基PleshkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pleshki",
		TitleName: "普列什基",
		TitleCode: "b_pleshki",
	}
}
