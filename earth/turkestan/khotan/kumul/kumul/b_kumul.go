package kumul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈密KumulBarony struct {
	feud.BaseBarony
}

var BKumul哈密 feud.Barony = &哈密KumulBarony{}

func init() {
	f := BKumul哈密.(*哈密KumulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumul",
		TitleName: "哈密",
		TitleCode: "b_kumul",
	}
}
