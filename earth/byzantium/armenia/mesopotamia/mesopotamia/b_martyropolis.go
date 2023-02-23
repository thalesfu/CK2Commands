package mesopotamia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马图罗波利斯MartyropolisBarony struct {
	feud.BaseBarony
}

var BMartyropolis马图罗波利斯 feud.Barony = &马图罗波利斯MartyropolisBarony{}

func init() {
	f := BMartyropolis马图罗波利斯.(*马图罗波利斯MartyropolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "martyropolis",
		TitleName: "马图罗波利斯",
		TitleCode: "b_martyropolis",
	}
}
