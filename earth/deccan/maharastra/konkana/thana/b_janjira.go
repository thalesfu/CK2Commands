package thana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 缮恃罗JanjiraBarony struct {
	feud.BaseBarony
}

var BJanjira缮恃罗 feud.Barony = &缮恃罗JanjiraBarony{}

func init() {
	f := BJanjira缮恃罗.(*缮恃罗JanjiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "janjira",
		TitleName: "缮恃罗",
		TitleCode: "b_janjira",
	}
}
