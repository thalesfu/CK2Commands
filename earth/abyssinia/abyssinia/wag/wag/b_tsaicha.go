package wag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎伊查TsaichaBarony struct {
	feud.BaseBarony
}

var BTsaicha扎伊查 feud.Barony = &扎伊查TsaichaBarony{}

func init() {
	f := BTsaicha扎伊查.(*扎伊查TsaichaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsaicha",
		TitleName: "扎伊查",
		TitleCode: "b_tsaicha",
	}
}
