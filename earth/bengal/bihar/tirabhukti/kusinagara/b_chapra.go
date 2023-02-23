package kusinagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 遮波罗ChapraBarony struct {
	feud.BaseBarony
}

var BChapra遮波罗 feud.Barony = &遮波罗ChapraBarony{}

func init() {
	f := BChapra遮波罗.(*遮波罗ChapraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chapra",
		TitleName: "遮波罗",
		TitleCode: "b_chapra",
	}
}
