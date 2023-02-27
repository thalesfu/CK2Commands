package nakhshab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐兹尔查QizilchaBarony struct {
	feud.BaseBarony
}

var BQizilcha齐兹尔查 feud.Barony = &齐兹尔查QizilchaBarony{}

func init() {
    f := BQizilcha齐兹尔查.(*齐兹尔查QizilchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qizilcha",
		TitleName: "齐兹尔查",
		TitleCode: "b_qizilcha",
	}
}
