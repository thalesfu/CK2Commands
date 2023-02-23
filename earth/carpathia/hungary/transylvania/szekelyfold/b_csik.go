package szekelyfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇克CsikBarony struct {
	feud.BaseBarony
}

var BCsik奇克 feud.Barony = &奇克CsikBarony{}

func init() {
	f := BCsik奇克.(*奇克CsikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "csik",
		TitleName: "奇克",
		TitleCode: "b_csik",
	}
}
