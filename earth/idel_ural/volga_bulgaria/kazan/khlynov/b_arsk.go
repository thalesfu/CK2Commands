package khlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔斯克ArskBarony struct {
	feud.BaseBarony
}

var BArsk阿尔斯克 feud.Barony = &阿尔斯克ArskBarony{}

func init() {
    f := BArsk阿尔斯克.(*阿尔斯克ArskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arsk",
		TitleName: "阿尔斯克",
		TitleCode: "b_arsk",
	}
}
