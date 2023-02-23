package asyut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古西耶KusaiBarony struct {
	feud.BaseBarony
}

var BKusai古西耶 feud.Barony = &古西耶KusaiBarony{}

func init() {
	f := BKusai古西耶.(*古西耶KusaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kusai",
		TitleName: "古西耶",
		TitleCode: "b_kusai",
	}
}
