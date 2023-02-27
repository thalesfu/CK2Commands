package aosta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费尼斯FenisBarony struct {
	feud.BaseBarony
}

var BFenis费尼斯 feud.Barony = &费尼斯FenisBarony{}

func init() {
    f := BFenis费尼斯.(*费尼斯FenisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fenis",
		TitleName: "费尼斯",
		TitleCode: "b_fenis",
	}
}
