package kara_khorum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈剌和林Kara_khorumBarony struct {
	feud.BaseBarony
}

var BKara_khorum哈剌和林 feud.Barony = &哈剌和林Kara_khorumBarony{}

func init() {
    f := BKara_khorum哈剌和林.(*哈剌和林Kara_khorumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kara_khorum",
		TitleName: "哈剌和林",
		TitleCode: "b_kara_khorum",
	}
}
