package cherchen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扜零YulingBarony struct {
	feud.BaseBarony
}

var BYuling扜零 feud.Barony = &扜零YulingBarony{}

func init() {
    f := BYuling扜零.(*扜零YulingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuling",
		TitleName: "扜零",
		TitleCode: "b_yuling",
	}
}
