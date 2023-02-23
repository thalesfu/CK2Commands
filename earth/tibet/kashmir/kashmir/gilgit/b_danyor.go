package gilgit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丹约尔DanyorBarony struct {
	feud.BaseBarony
}

var BDanyor丹约尔 feud.Barony = &丹约尔DanyorBarony{}

func init() {
	f := BDanyor丹约尔.(*丹约尔DanyorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "danyor",
		TitleName: "丹约尔",
		TitleCode: "b_danyor",
	}
}
