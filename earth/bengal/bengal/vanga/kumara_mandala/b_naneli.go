package kumara_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那内梨NaneliBarony struct {
	feud.BaseBarony
}

var BNaneli那内梨 feud.Barony = &那内梨NaneliBarony{}

func init() {
    f := BNaneli那内梨.(*那内梨NaneliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naneli",
		TitleName: "那内梨",
		TitleCode: "b_naneli",
	}
}
