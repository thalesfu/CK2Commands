package goalpara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔阿罗波罗GoalparaBarony struct {
	feud.BaseBarony
}

var BGoalpara乔阿罗波罗 feud.Barony = &乔阿罗波罗GoalparaBarony{}

func init() {
    f := BGoalpara乔阿罗波罗.(*乔阿罗波罗GoalparaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goalpara",
		TitleName: "乔阿罗波罗",
		TitleCode: "b_goalpara",
	}
}
