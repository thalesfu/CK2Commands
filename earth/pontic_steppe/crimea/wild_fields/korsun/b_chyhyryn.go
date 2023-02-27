package korsun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇吉林ChyhyrynBarony struct {
	feud.BaseBarony
}

var BChyhyryn奇吉林 feud.Barony = &奇吉林ChyhyrynBarony{}

func init() {
    f := BChyhyryn奇吉林.(*奇吉林ChyhyrynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chyhyryn",
		TitleName: "奇吉林",
		TitleCode: "b_chyhyryn",
	}
}
