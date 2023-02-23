package agder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊夫兰IvelandBarony struct {
	feud.BaseBarony
}

var BIveland伊夫兰 feud.Barony = &伊夫兰IvelandBarony{}

func init() {
	f := BIveland伊夫兰.(*伊夫兰IvelandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iveland",
		TitleName: "伊夫兰",
		TitleCode: "b_iveland",
	}
}
