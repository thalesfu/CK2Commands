package tara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔贝扎KarbyzaBarony struct {
	feud.BaseBarony
}

var BKarbyza卡尔贝扎 feud.Barony = &卡尔贝扎KarbyzaBarony{}

func init() {
    f := BKarbyza卡尔贝扎.(*卡尔贝扎KarbyzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karbyza",
		TitleName: "卡尔贝扎",
		TitleCode: "b_karbyza",
	}
}
