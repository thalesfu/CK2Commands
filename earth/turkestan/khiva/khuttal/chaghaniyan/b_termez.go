package chaghaniyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呾蜜TermezBarony struct {
	feud.BaseBarony
}

var BTermez呾蜜 feud.Barony = &呾蜜TermezBarony{}

func init() {
	f := BTermez呾蜜.(*呾蜜TermezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "termez",
		TitleName: "呾蜜",
		TitleCode: "b_termez",
	}
}
