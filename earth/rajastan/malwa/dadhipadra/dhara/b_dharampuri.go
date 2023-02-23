package dhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃愣补罗DharampuriBarony struct {
	feud.BaseBarony
}

var BDharampuri诃愣补罗 feud.Barony = &诃愣补罗DharampuriBarony{}

func init() {
	f := BDharampuri诃愣补罗.(*诃愣补罗DharampuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dharampuri",
		TitleName: "诃愣补罗",
		TitleCode: "b_dharampuri",
	}
}
