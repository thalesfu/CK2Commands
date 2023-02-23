package cephalonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕兰法里兰PalaiofrourioBarony struct {
	feud.BaseBarony
}

var BPalaiofrourio帕兰法里兰 feud.Barony = &帕兰法里兰PalaiofrourioBarony{}

func init() {
	f := BPalaiofrourio帕兰法里兰.(*帕兰法里兰PalaiofrourioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palaiofrourio",
		TitleName: "帕兰法里兰",
		TitleCode: "b_palaiofrourio",
	}
}
