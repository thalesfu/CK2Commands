package harer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库发彻拉KurfacheleBarony struct {
	feud.BaseBarony
}

var BKurfachele库发彻拉 feud.Barony = &库发彻拉KurfacheleBarony{}

func init() {
    f := BKurfachele库发彻拉.(*库发彻拉KurfacheleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kurfachele",
		TitleName: "库发彻拉",
		TitleCode: "b_kurfachele",
	}
}
