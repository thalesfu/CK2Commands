package yaroslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗曼诺夫RomanovBarony struct {
	feud.BaseBarony
}

var BRomanov罗曼诺夫 feud.Barony = &罗曼诺夫RomanovBarony{}

func init() {
	f := BRomanov罗曼诺夫.(*罗曼诺夫RomanovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "romanov",
		TitleName: "罗曼诺夫",
		TitleCode: "b_romanov",
	}
}
