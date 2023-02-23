package senoussi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布宰迈BuzaymanBarony struct {
	feud.BaseBarony
}

var BBuzayman布宰迈 feud.Barony = &布宰迈BuzaymanBarony{}

func init() {
	f := BBuzayman布宰迈.(*布宰迈BuzaymanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buzayman",
		TitleName: "布宰迈",
		TitleCode: "b_buzayman",
	}
}
