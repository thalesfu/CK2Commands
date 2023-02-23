package phiti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿吒摩萨傥那AtamasthanaBarony struct {
	feud.BaseBarony
}

var BAtamasthana阿吒摩萨傥那 feud.Barony = &阿吒摩萨傥那AtamasthanaBarony{}

func init() {
	f := BAtamasthana阿吒摩萨傥那.(*阿吒摩萨傥那AtamasthanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atamasthana",
		TitleName: "阿吒摩萨傥那",
		TitleCode: "b_atamasthana",
	}
}
