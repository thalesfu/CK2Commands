package ulm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔姆UlmBarony struct {
	feud.BaseBarony
}

var BUlm乌尔姆 feud.Barony = &乌尔姆UlmBarony{}

func init() {
    f := BUlm乌尔姆.(*乌尔姆UlmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulm",
		TitleName: "乌尔姆",
		TitleCode: "b_ulm",
	}
}
