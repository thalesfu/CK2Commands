package tyrone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图里霍格TullyhogueBarony struct {
	feud.BaseBarony
}

var BTullyhogue图里霍格 feud.Barony = &图里霍格TullyhogueBarony{}

func init() {
    f := BTullyhogue图里霍格.(*图里霍格TullyhogueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tullyhogue",
		TitleName: "图里霍格",
		TitleCode: "b_tullyhogue",
	}
}
