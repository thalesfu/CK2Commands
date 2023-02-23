package korosten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳罗季奇NarodytchiBarony struct {
	feud.BaseBarony
}

var BNarodytchi纳罗季奇 feud.Barony = &纳罗季奇NarodytchiBarony{}

func init() {
	f := BNarodytchi纳罗季奇.(*纳罗季奇NarodytchiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narodytchi",
		TitleName: "纳罗季奇",
		TitleCode: "b_narodytchi",
	}
}
