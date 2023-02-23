package tis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕拉克ParakBarony struct {
	feud.BaseBarony
}

var BParak帕拉克 feud.Barony = &帕拉克ParakBarony{}

func init() {
	f := BParak帕拉克.(*帕拉克ParakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parak",
		TitleName: "帕拉克",
		TitleCode: "b_parak",
	}
}
