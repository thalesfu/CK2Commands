package irbid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆盖斯UmmqaisBarony struct {
	feud.BaseBarony
}

var BUmmqais乌姆盖斯 feud.Barony = &乌姆盖斯UmmqaisBarony{}

func init() {
    f := BUmmqais乌姆盖斯.(*乌姆盖斯UmmqaisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ummqais",
		TitleName: "乌姆盖斯",
		TitleCode: "b_ummqais",
	}
}
