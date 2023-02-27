package kara_khorum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈木黑KhamagBarony struct {
	feud.BaseBarony
}

var BKhamag哈木黑 feud.Barony = &哈木黑KhamagBarony{}

func init() {
    f := BKhamag哈木黑.(*哈木黑KhamagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khamag",
		TitleName: "哈木黑",
		TitleCode: "b_khamag",
	}
}
