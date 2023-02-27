package durdzukia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉吉尔AlagirBarony struct {
	feud.BaseBarony
}

var BAlagir阿拉吉尔 feud.Barony = &阿拉吉尔AlagirBarony{}

func init() {
    f := BAlagir阿拉吉尔.(*阿拉吉尔AlagirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alagir",
		TitleName: "阿拉吉尔",
		TitleCode: "b_alagir",
	}
}
