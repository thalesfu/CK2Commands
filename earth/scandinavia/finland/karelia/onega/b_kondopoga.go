package onega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔多波加KondopogaBarony struct {
	feud.BaseBarony
}

var BKondopoga孔多波加 feud.Barony = &孔多波加KondopogaBarony{}

func init() {
    f := BKondopoga孔多波加.(*孔多波加KondopogaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kondopoga",
		TitleName: "孔多波加",
		TitleCode: "b_kondopoga",
	}
}
