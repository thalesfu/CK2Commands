package trapezous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉佩祖斯TrapezousBarony struct {
	feud.BaseBarony
}

var BTrapezous特拉佩祖斯 feud.Barony = &特拉佩祖斯TrapezousBarony{}

func init() {
    f := BTrapezous特拉佩祖斯.(*特拉佩祖斯TrapezousBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trapezous",
		TitleName: "特拉佩祖斯",
		TitleCode: "b_trapezous",
	}
}
