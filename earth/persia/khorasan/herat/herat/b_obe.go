package herat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥比ObeBarony struct {
	feud.BaseBarony
}

var BObe奥比 feud.Barony = &奥比ObeBarony{}

func init() {
    f := BObe奥比.(*奥比ObeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "obe",
		TitleName: "奥比",
		TitleCode: "b_obe",
	}
}
