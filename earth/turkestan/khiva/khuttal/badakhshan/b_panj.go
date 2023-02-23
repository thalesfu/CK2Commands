package badakhshan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喷赤PanjBarony struct {
	feud.BaseBarony
}

var BPanj喷赤 feud.Barony = &喷赤PanjBarony{}

func init() {
	f := BPanj喷赤.(*喷赤PanjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panj",
		TitleName: "喷赤",
		TitleCode: "b_panj",
	}
}
