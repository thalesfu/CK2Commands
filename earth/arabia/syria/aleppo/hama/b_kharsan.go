package hama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉桑KharsanBarony struct {
	feud.BaseBarony
}

var BKharsan哈拉桑 feud.Barony = &哈拉桑KharsanBarony{}

func init() {
    f := BKharsan哈拉桑.(*哈拉桑KharsanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharsan",
		TitleName: "哈拉桑",
		TitleCode: "b_kharsan",
	}
}
