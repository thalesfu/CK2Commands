package khangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌里雅苏台UliastaiBarony struct {
	feud.BaseBarony
}

var BUliastai乌里雅苏台 feud.Barony = &乌里雅苏台UliastaiBarony{}

func init() {
    f := BUliastai乌里雅苏台.(*乌里雅苏台UliastaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uliastai",
		TitleName: "乌里雅苏台",
		TitleCode: "b_uliastai",
	}
}
