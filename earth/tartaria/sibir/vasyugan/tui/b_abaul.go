package tui

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴乌尔AbaulBarony struct {
	feud.BaseBarony
}

var BAbaul阿巴乌尔 feud.Barony = &阿巴乌尔AbaulBarony{}

func init() {
	f := BAbaul阿巴乌尔.(*阿巴乌尔AbaulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abaul",
		TitleName: "阿巴乌尔",
		TitleCode: "b_abaul",
	}
}
