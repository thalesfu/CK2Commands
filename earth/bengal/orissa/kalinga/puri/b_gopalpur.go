package puri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿波罗补罗GopalpurBarony struct {
	feud.BaseBarony
}

var BGopalpur瞿波罗补罗 feud.Barony = &瞿波罗补罗GopalpurBarony{}

func init() {
    f := BGopalpur瞿波罗补罗.(*瞿波罗补罗GopalpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gopalpur",
		TitleName: "瞿波罗补罗",
		TitleCode: "b_gopalpur",
	}
}
