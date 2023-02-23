package shrewsbury

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 惠特彻奇WhitchurchBarony struct {
	feud.BaseBarony
}

var BWhitchurch惠特彻奇 feud.Barony = &惠特彻奇WhitchurchBarony{}

func init() {
	f := BWhitchurch惠特彻奇.(*惠特彻奇WhitchurchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "whitchurch",
		TitleName: "惠特彻奇",
		TitleCode: "b_whitchurch",
	}
}
