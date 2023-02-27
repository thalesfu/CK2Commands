package tmutarakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采梅斯TsemesBarony struct {
	feud.BaseBarony
}

var BTsemes采梅斯 feud.Barony = &采梅斯TsemesBarony{}

func init() {
    f := BTsemes采梅斯.(*采梅斯TsemesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsemes",
		TitleName: "采梅斯",
		TitleCode: "b_tsemes",
	}
}
