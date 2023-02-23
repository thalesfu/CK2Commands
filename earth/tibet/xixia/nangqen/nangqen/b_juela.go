package nangqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 觉拉JuelaBarony struct {
	feud.BaseBarony
}

var BJuela觉拉 feud.Barony = &觉拉JuelaBarony{}

func init() {
	f := BJuela觉拉.(*觉拉JuelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "juela",
		TitleName: "觉拉",
		TitleCode: "b_juela",
	}
}
