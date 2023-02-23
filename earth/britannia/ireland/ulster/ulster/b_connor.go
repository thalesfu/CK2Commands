package ulster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康诺尔ConnorBarony struct {
	feud.BaseBarony
}

var BConnor康诺尔 feud.Barony = &康诺尔ConnorBarony{}

func init() {
	f := BConnor康诺尔.(*康诺尔ConnorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "connor",
		TitleName: "康诺尔",
		TitleCode: "b_connor",
	}
}
