package samtho

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玉盘YulpunBarony struct {
	feud.BaseBarony
}

var BYulpun玉盘 feud.Barony = &玉盘YulpunBarony{}

func init() {
	f := BYulpun玉盘.(*玉盘YulpunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yulpun",
		TitleName: "玉盘",
		TitleCode: "b_yulpun",
	}
}
