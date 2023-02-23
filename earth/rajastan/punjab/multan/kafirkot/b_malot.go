package kafirkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马诺特MalotBarony struct {
	feud.BaseBarony
}

var BMalot马诺特 feud.Barony = &马诺特MalotBarony{}

func init() {
	f := BMalot马诺特.(*马诺特MalotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malot",
		TitleName: "马诺特",
		TitleCode: "b_malot",
	}
}
