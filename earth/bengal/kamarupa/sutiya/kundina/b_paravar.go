package kundina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 钵罗跋罗ParavarBarony struct {
	feud.BaseBarony
}

var BParavar钵罗跋罗 feud.Barony = &钵罗跋罗ParavarBarony{}

func init() {
	f := BParavar钵罗跋罗.(*钵罗跋罗ParavarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paravar",
		TitleName: "钵罗跋罗",
		TitleCode: "b_paravar",
	}
}
