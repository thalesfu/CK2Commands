package sistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毾㲪山Kuh_taftanBarony struct {
	feud.BaseBarony
}

var BKuh_taftan毾㲪山 feud.Barony = &毾㲪山Kuh_taftanBarony{}

func init() {
    f := BKuh_taftan毾㲪山.(*毾㲪山Kuh_taftanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuh_taftan",
		TitleName: "毾㲪山",
		TitleCode: "b_kuh_taftan",
	}
}
