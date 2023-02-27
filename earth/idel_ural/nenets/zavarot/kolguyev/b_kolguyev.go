package kolguyev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔古耶夫KolguyevBarony struct {
	feud.BaseBarony
}

var BKolguyev科尔古耶夫 feud.Barony = &科尔古耶夫KolguyevBarony{}

func init() {
    f := BKolguyev科尔古耶夫.(*科尔古耶夫KolguyevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolguyev",
		TitleName: "科尔古耶夫",
		TitleCode: "b_kolguyev",
	}
}
