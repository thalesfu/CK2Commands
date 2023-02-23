package varanasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆罗利阿BalliaBarony struct {
	feud.BaseBarony
}

var BBallia婆罗利阿 feud.Barony = &婆罗利阿BalliaBarony{}

func init() {
	f := BBallia婆罗利阿.(*婆罗利阿BalliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ballia",
		TitleName: "婆罗利阿",
		TitleCode: "b_ballia",
	}
}
