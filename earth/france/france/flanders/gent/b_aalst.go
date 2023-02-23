package gent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔斯特AalstBarony struct {
	feud.BaseBarony
}

var BAalst阿尔斯特 feud.Barony = &阿尔斯特AalstBarony{}

func init() {
	f := BAalst阿尔斯特.(*阿尔斯特AalstBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aalst",
		TitleName: "阿尔斯特",
		TitleCode: "b_aalst",
	}
}
