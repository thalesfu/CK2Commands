package oland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌滕比OttenbyBarony struct {
	feud.BaseBarony
}

var BOttenby乌滕比 feud.Barony = &乌滕比OttenbyBarony{}

func init() {
	f := BOttenby乌滕比.(*乌滕比OttenbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ottenby",
		TitleName: "乌滕比",
		TitleCode: "b_ottenby",
	}
}
