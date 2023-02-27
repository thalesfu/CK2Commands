package euboia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔基斯ChalkisBarony struct {
	feud.BaseBarony
}

var BChalkis哈尔基斯 feud.Barony = &哈尔基斯ChalkisBarony{}

func init() {
    f := BChalkis哈尔基斯.(*哈尔基斯ChalkisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chalkis",
		TitleName: "哈尔基斯",
		TitleCode: "b_chalkis",
	}
}
