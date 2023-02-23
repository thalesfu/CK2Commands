package eu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄镇EuBarony struct {
	feud.BaseBarony
}

var BEu厄镇 feud.Barony = &厄镇EuBarony{}

func init() {
	f := BEu厄镇.(*厄镇EuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eu",
		TitleName: "厄镇",
		TitleCode: "b_eu",
	}
}
