package uvs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特斯TesBarony struct {
	feud.BaseBarony
}

var BTes特斯 feud.Barony = &特斯TesBarony{}

func init() {
	f := BTes特斯.(*特斯TesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tes",
		TitleName: "特斯",
		TitleCode: "b_tes",
	}
}
