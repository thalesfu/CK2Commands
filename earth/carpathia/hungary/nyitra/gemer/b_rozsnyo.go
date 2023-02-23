package gemer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗日纽RozsnyoBarony struct {
	feud.BaseBarony
}

var BRozsnyo罗日纽 feud.Barony = &罗日纽RozsnyoBarony{}

func init() {
	f := BRozsnyo罗日纽.(*罗日纽RozsnyoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rozsnyo",
		TitleName: "罗日纽",
		TitleCode: "b_rozsnyo",
	}
}
