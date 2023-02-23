package menorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯卡斯特利EscastellBarony struct {
	feud.BaseBarony
}

var BEscastell埃斯卡斯特利 feud.Barony = &埃斯卡斯特利EscastellBarony{}

func init() {
	f := BEscastell埃斯卡斯特利.(*埃斯卡斯特利EscastellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "escastell",
		TitleName: "埃斯卡斯特利",
		TitleCode: "b_escastell",
	}
}
