package tarvagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松吉瑙SonginoBarony struct {
	feud.BaseBarony
}

var BSongino松吉瑙 feud.Barony = &松吉瑙SonginoBarony{}

func init() {
    f := BSongino松吉瑙.(*松吉瑙SonginoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "songino",
		TitleName: "松吉瑙",
		TitleCode: "b_songino",
	}
}
