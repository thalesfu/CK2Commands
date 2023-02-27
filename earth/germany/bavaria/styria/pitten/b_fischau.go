package pitten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲绍FischauBarony struct {
	feud.BaseBarony
}

var BFischau菲绍 feud.Barony = &菲绍FischauBarony{}

func init() {
    f := BFischau菲绍.(*菲绍FischauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fischau",
		TitleName: "菲绍",
		TitleCode: "b_fischau",
	}
}
