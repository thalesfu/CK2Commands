package qwivir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达夫瓦DarvarBarony struct {
	feud.BaseBarony
}

var BDarvar达夫瓦 feud.Barony = &达夫瓦DarvarBarony{}

func init() {
	f := BDarvar达夫瓦.(*达夫瓦DarvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darvar",
		TitleName: "达夫瓦",
		TitleCode: "b_darvar",
	}
}
