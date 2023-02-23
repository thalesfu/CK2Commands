package prayaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 憍赏弥KausambiBarony struct {
	feud.BaseBarony
}

var BKausambi憍赏弥 feud.Barony = &憍赏弥KausambiBarony{}

func init() {
	f := BKausambi憍赏弥.(*憍赏弥KausambiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kausambi",
		TitleName: "憍赏弥",
		TitleCode: "b_kausambi",
	}
}
