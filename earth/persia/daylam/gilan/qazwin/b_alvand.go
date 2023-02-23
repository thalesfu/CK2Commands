package qazwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔万德AlvandBarony struct {
	feud.BaseBarony
}

var BAlvand阿尔万德 feud.Barony = &阿尔万德AlvandBarony{}

func init() {
	f := BAlvand阿尔万德.(*阿尔万德AlvandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alvand",
		TitleName: "阿尔万德",
		TitleCode: "b_alvand",
	}
}
