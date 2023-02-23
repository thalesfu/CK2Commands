package arta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格南塔AgnantaBarony struct {
	feud.BaseBarony
}

var BAgnanta阿格南塔 feud.Barony = &阿格南塔AgnantaBarony{}

func init() {
	f := BAgnanta阿格南塔.(*阿格南塔AgnantaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agnanta",
		TitleName: "阿格南塔",
		TitleCode: "b_agnanta",
	}
}
