package qangtang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雪源ZhoyonBarony struct {
	feud.BaseBarony
}

var BZhoyon雪源 feud.Barony = &雪源ZhoyonBarony{}

func init() {
    f := BZhoyon雪源.(*雪源ZhoyonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhoyon",
		TitleName: "雪源",
		TitleCode: "b_zhoyon",
	}
}
