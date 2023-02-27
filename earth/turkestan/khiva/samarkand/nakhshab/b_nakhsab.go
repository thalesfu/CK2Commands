package nakhshab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那色波NakhsabBarony struct {
	feud.BaseBarony
}

var BNakhsab那色波 feud.Barony = &那色波NakhsabBarony{}

func init() {
    f := BNakhsab那色波.(*那色波NakhsabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nakhsab",
		TitleName: "那色波",
		TitleCode: "b_nakhsab",
	}
}
