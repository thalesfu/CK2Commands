package nakhshab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌拉迟UkrachBarony struct {
	feud.BaseBarony
}

var BUkrach乌拉迟 feud.Barony = &乌拉迟UkrachBarony{}

func init() {
    f := BUkrach乌拉迟.(*乌拉迟UkrachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ukrach",
		TitleName: "乌拉迟",
		TitleCode: "b_ukrach",
	}
}
