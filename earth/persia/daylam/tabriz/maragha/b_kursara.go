package maragha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔萨拉KursaraBarony struct {
	feud.BaseBarony
}

var BKursara库尔萨拉 feud.Barony = &库尔萨拉KursaraBarony{}

func init() {
    f := BKursara库尔萨拉.(*库尔萨拉KursaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kursara",
		TitleName: "库尔萨拉",
		TitleCode: "b_kursara",
	}
}
