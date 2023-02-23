package burhanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 章伽提婆ChangdevBarony struct {
	feud.BaseBarony
}

var BChangdev章伽提婆 feud.Barony = &章伽提婆ChangdevBarony{}

func init() {
	f := BChangdev章伽提婆.(*章伽提婆ChangdevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "changdev",
		TitleName: "章伽提婆",
		TitleCode: "b_changdev",
	}
}
