package napoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯基亚IschiaBarony struct {
	feud.BaseBarony
}

var BIschia伊斯基亚 feud.Barony = &伊斯基亚IschiaBarony{}

func init() {
	f := BIschia伊斯基亚.(*伊斯基亚IschiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ischia",
		TitleName: "伊斯基亚",
		TitleCode: "b_ischia",
	}
}
