package aksu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柯坪KelpinBarony struct {
	feud.BaseBarony
}

var BKelpin柯坪 feud.Barony = &柯坪KelpinBarony{}

func init() {
	f := BKelpin柯坪.(*柯坪KelpinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kelpin",
		TitleName: "柯坪",
		TitleCode: "b_kelpin",
	}
}
