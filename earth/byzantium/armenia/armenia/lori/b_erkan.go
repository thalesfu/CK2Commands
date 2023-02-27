package lori

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄而罕ErkanBarony struct {
	feud.BaseBarony
}

var BErkan厄而罕 feud.Barony = &厄而罕ErkanBarony{}

func init() {
    f := BErkan厄而罕.(*厄而罕ErkanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erkan",
		TitleName: "厄而罕",
		TitleCode: "b_erkan",
	}
}
