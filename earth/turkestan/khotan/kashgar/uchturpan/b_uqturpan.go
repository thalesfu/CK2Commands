package uchturpan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌什UqturpanBarony struct {
	feud.BaseBarony
}

var BUqturpan乌什 feud.Barony = &乌什UqturpanBarony{}

func init() {
    f := BUqturpan乌什.(*乌什UqturpanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uqturpan",
		TitleName: "乌什",
		TitleCode: "b_uqturpan",
	}
}
