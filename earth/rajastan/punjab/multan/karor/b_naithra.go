package karor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈多罗NaithraBarony struct {
	feud.BaseBarony
}

var BNaithra奈多罗 feud.Barony = &奈多罗NaithraBarony{}

func init() {
	f := BNaithra奈多罗.(*奈多罗NaithraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naithra",
		TitleName: "奈多罗",
		TitleCode: "b_naithra",
	}
}
