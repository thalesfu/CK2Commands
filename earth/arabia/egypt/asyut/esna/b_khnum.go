package esna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫努姆KhnumBarony struct {
	feud.BaseBarony
}

var BKhnum赫努姆 feud.Barony = &赫努姆KhnumBarony{}

func init() {
    f := BKhnum赫努姆.(*赫努姆KhnumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khnum",
		TitleName: "赫努姆",
		TitleCode: "b_khnum",
	}
}
