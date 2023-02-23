package chur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格卢恩斯GlurnsBarony struct {
	feud.BaseBarony
}

var BGlurns格卢恩斯 feud.Barony = &格卢恩斯GlurnsBarony{}

func init() {
	f := BGlurns格卢恩斯.(*格卢恩斯GlurnsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glurns",
		TitleName: "格卢恩斯",
		TitleCode: "b_glurns",
	}
}
