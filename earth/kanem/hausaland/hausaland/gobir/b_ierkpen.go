package gobir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔克彭IerkpenBarony struct {
	feud.BaseBarony
}

var BIerkpen伊尔克彭 feud.Barony = &伊尔克彭IerkpenBarony{}

func init() {
	f := BIerkpen伊尔克彭.(*伊尔克彭IerkpenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ierkpen",
		TitleName: "伊尔克彭",
		TitleCode: "b_ierkpen",
	}
}
