package samarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔古特UrgutBarony struct {
	feud.BaseBarony
}

var BUrgut乌尔古特 feud.Barony = &乌尔古特UrgutBarony{}

func init() {
    f := BUrgut乌尔古特.(*乌尔古特UrgutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urgut",
		TitleName: "乌尔古特",
		TitleCode: "b_urgut",
	}
}
