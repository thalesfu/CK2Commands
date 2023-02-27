package leptis_magna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉拉达赫QirdahBarony struct {
	feud.BaseBarony
}

var BQirdah吉拉达赫 feud.Barony = &吉拉达赫QirdahBarony{}

func init() {
    f := BQirdah吉拉达赫.(*吉拉达赫QirdahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qirdah",
		TitleName: "吉拉达赫",
		TitleCode: "b_qirdah",
	}
}
