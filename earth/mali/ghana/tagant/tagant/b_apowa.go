package tagant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿波瓦ApowaBarony struct {
	feud.BaseBarony
}

var BApowa阿波瓦 feud.Barony = &阿波瓦ApowaBarony{}

func init() {
    f := BApowa阿波瓦.(*阿波瓦ApowaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "apowa",
		TitleName: "阿波瓦",
		TitleCode: "b_apowa",
	}
}
