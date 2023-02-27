package lhunze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玉麦YumaiBarony struct {
	feud.BaseBarony
}

var BYumai玉麦 feud.Barony = &玉麦YumaiBarony{}

func init() {
    f := BYumai玉麦.(*玉麦YumaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yumai",
		TitleName: "玉麦",
		TitleCode: "b_yumai",
	}
}
