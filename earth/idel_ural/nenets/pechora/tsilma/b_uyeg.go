package tsilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌耶格UyegBarony struct {
	feud.BaseBarony
}

var BUyeg乌耶格 feud.Barony = &乌耶格UyegBarony{}

func init() {
    f := BUyeg乌耶格.(*乌耶格UyegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uyeg",
		TitleName: "乌耶格",
		TitleCode: "b_uyeg",
	}
}
