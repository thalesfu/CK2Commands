package kucha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库木吐喇KumturaBarony struct {
	feud.BaseBarony
}

var BKumtura库木吐喇 feud.Barony = &库木吐喇KumturaBarony{}

func init() {
	f := BKumtura库木吐喇.(*库木吐喇KumturaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumtura",
		TitleName: "库木吐喇",
		TitleCode: "b_kumtura",
	}
}
