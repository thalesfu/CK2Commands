package peresechen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图扎拉TuzaraBarony struct {
	feud.BaseBarony
}

var BTuzara图扎拉 feud.Barony = &图扎拉TuzaraBarony{}

func init() {
	f := BTuzara图扎拉.(*图扎拉TuzaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuzara",
		TitleName: "图扎拉",
		TitleCode: "b_tuzara",
	}
}
