package ranthambore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌德吉尔UtgirBarony struct {
	feud.BaseBarony
}

var BUtgir乌德吉尔 feud.Barony = &乌德吉尔UtgirBarony{}

func init() {
    f := BUtgir乌德吉尔.(*乌德吉尔UtgirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "utgir",
		TitleName: "乌德吉尔",
		TitleCode: "b_utgir",
	}
}
