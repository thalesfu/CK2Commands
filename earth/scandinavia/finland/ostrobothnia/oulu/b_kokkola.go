package oulu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科科拉KokkolaBarony struct {
	feud.BaseBarony
}

var BKokkola科科拉 feud.Barony = &科科拉KokkolaBarony{}

func init() {
	f := BKokkola科科拉.(*科科拉KokkolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kokkola",
		TitleName: "科科拉",
		TitleCode: "b_kokkola",
	}
}
