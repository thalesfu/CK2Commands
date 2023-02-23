package salzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔多夫MuhldorfBarony struct {
	feud.BaseBarony
}

var BMuhldorf米尔多夫 feud.Barony = &米尔多夫MuhldorfBarony{}

func init() {
	f := BMuhldorf米尔多夫.(*米尔多夫MuhldorfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muhldorf",
		TitleName: "米尔多夫",
		TitleCode: "b_muhldorf",
	}
}
