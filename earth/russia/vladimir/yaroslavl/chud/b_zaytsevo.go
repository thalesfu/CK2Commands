package chud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宰采沃ZaytsevoBarony struct {
	feud.BaseBarony
}

var BZaytsevo宰采沃 feud.Barony = &宰采沃ZaytsevoBarony{}

func init() {
	f := BZaytsevo宰采沃.(*宰采沃ZaytsevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaytsevo",
		TitleName: "宰采沃",
		TitleCode: "b_zaytsevo",
	}
}
