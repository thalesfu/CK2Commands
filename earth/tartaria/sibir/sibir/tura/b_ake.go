package tura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克AkeBarony struct {
	feud.BaseBarony
}

var BAke阿克 feud.Barony = &阿克AkeBarony{}

func init() {
	f := BAke阿克.(*阿克AkeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ake",
		TitleName: "阿克",
		TitleCode: "b_ake",
	}
}
