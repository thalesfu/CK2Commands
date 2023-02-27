package suf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法特纳萨FatnassaBarony struct {
	feud.BaseBarony
}

var BFatnassa法特纳萨 feud.Barony = &法特纳萨FatnassaBarony{}

func init() {
    f := BFatnassa法特纳萨.(*法特纳萨FatnassaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fatnassa",
		TitleName: "法特纳萨",
		TitleCode: "b_fatnassa",
	}
}
