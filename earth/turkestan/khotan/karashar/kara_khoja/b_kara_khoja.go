package kara_khoja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈剌火州Kara_khojaBarony struct {
	feud.BaseBarony
}

var BKara_khoja哈剌火州 feud.Barony = &哈剌火州Kara_khojaBarony{}

func init() {
    f := BKara_khoja哈剌火州.(*哈剌火州Kara_khojaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kara_khoja",
		TitleName: "哈剌火州",
		TitleCode: "b_kara_khoja",
	}
}
