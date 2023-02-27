package euboia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇米KymiBarony struct {
	feud.BaseBarony
}

var BKymi奇米 feud.Barony = &奇米KymiBarony{}

func init() {
    f := BKymi奇米.(*奇米KymiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kymi",
		TitleName: "奇米",
		TitleCode: "b_kymi",
	}
}
