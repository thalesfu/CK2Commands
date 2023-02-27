package chach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别斯干PskentBarony struct {
	feud.BaseBarony
}

var BPskent别斯干 feud.Barony = &别斯干PskentBarony{}

func init() {
    f := BPskent别斯干.(*别斯干PskentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pskent",
		TitleName: "别斯干",
		TitleCode: "b_pskent",
	}
}
