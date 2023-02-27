package medog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德兴DeshingBarony struct {
	feud.BaseBarony
}

var BDeshing德兴 feud.Barony = &德兴DeshingBarony{}

func init() {
    f := BDeshing德兴.(*德兴DeshingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deshing",
		TitleName: "德兴",
		TitleCode: "b_deshing",
	}
}
