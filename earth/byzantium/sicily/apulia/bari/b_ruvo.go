package bari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁沃RuvoBarony struct {
	feud.BaseBarony
}

var BRuvo鲁沃 feud.Barony = &鲁沃RuvoBarony{}

func init() {
    f := BRuvo鲁沃.(*鲁沃RuvoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruvo",
		TitleName: "鲁沃",
		TitleCode: "b_ruvo",
	}
}
