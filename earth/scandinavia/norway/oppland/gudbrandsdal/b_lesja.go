package gudbrandsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱沙LesjaBarony struct {
	feud.BaseBarony
}

var BLesja莱沙 feud.Barony = &莱沙LesjaBarony{}

func init() {
	f := BLesja莱沙.(*莱沙LesjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lesja",
		TitleName: "莱沙",
		TitleCode: "b_lesja",
	}
}
