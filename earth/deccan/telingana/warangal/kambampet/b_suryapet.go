package kambampet

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏利耶佩特SuryapetBarony struct {
	feud.BaseBarony
}

var BSuryapet苏利耶佩特 feud.Barony = &苏利耶佩特SuryapetBarony{}

func init() {
	f := BSuryapet苏利耶佩特.(*苏利耶佩特SuryapetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suryapet",
		TitleName: "苏利耶佩特",
		TitleCode: "b_suryapet",
	}
}
