package dyrrachion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 都拉斯DurazzoBarony struct {
	feud.BaseBarony
}

var BDurazzo都拉斯 feud.Barony = &都拉斯DurazzoBarony{}

func init() {
	f := BDurazzo都拉斯.(*都拉斯DurazzoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durazzo",
		TitleName: "都拉斯",
		TitleCode: "b_durazzo",
	}
}
