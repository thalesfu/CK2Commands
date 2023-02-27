package ashli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图伊马济TuimasyBarony struct {
	feud.BaseBarony
}

var BTuimasy图伊马济 feud.Barony = &图伊马济TuimasyBarony{}

func init() {
    f := BTuimasy图伊马济.(*图伊马济TuimasyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuimasy",
		TitleName: "图伊马济",
		TitleCode: "b_tuimasy",
	}
}
