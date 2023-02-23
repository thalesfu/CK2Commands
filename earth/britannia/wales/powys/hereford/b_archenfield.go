package hereford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿肯菲尔德ArchenfieldBarony struct {
	feud.BaseBarony
}

var BArchenfield阿肯菲尔德 feud.Barony = &阿肯菲尔德ArchenfieldBarony{}

func init() {
	f := BArchenfield阿肯菲尔德.(*阿肯菲尔德ArchenfieldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "archenfield",
		TitleName: "阿肯菲尔德",
		TitleCode: "b_archenfield",
	}
}
