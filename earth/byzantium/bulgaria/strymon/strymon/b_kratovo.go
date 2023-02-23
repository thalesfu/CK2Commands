package strymon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉托沃KratovoBarony struct {
	feud.BaseBarony
}

var BKratovo克拉托沃 feud.Barony = &克拉托沃KratovoBarony{}

func init() {
	f := BKratovo克拉托沃.(*克拉托沃KratovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kratovo",
		TitleName: "克拉托沃",
		TitleCode: "b_kratovo",
	}
}
