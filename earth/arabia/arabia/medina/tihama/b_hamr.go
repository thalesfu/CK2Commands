package tihama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈穆伦HamrBarony struct {
	feud.BaseBarony
}

var BHamr哈穆伦 feud.Barony = &哈穆伦HamrBarony{}

func init() {
	f := BHamr哈穆伦.(*哈穆伦HamrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamr",
		TitleName: "哈穆伦",
		TitleCode: "b_hamr",
	}
}
