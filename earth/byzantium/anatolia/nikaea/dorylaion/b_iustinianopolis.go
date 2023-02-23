package dorylaion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 优士丁尼波利斯IustinianopolisBarony struct {
	feud.BaseBarony
}

var BIustinianopolis优士丁尼波利斯 feud.Barony = &优士丁尼波利斯IustinianopolisBarony{}

func init() {
	f := BIustinianopolis优士丁尼波利斯.(*优士丁尼波利斯IustinianopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iustinianopolis",
		TitleName: "优士丁尼波利斯",
		TitleCode: "b_iustinianopolis",
	}
}
