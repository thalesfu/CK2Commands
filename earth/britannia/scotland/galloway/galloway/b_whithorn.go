package galloway

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 惠特霍恩WhithornBarony struct {
	feud.BaseBarony
}

var BWhithorn惠特霍恩 feud.Barony = &惠特霍恩WhithornBarony{}

func init() {
	f := BWhithorn惠特霍恩.(*惠特霍恩WhithornBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "whithorn",
		TitleName: "惠特霍恩",
		TitleCode: "b_whithorn",
	}
}
