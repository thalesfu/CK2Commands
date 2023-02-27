package srihatta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡奴KunnauBarony struct {
	feud.BaseBarony
}

var BKunnau贡奴 feud.Barony = &贡奴KunnauBarony{}

func init() {
    f := BKunnau贡奴.(*贡奴KunnauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunnau",
		TitleName: "贡奴",
		TitleCode: "b_kunnau",
	}
}
