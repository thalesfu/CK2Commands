package rummah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏费里AssufayriBarony struct {
	feud.BaseBarony
}

var BAssufayri苏费里 feud.Barony = &苏费里AssufayriBarony{}

func init() {
	f := BAssufayri苏费里.(*苏费里AssufayriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assufayri",
		TitleName: "苏费里",
		TitleCode: "b_assufayri",
	}
}
