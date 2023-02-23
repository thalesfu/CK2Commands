package kempten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肯普滕KemptenBarony struct {
	feud.BaseBarony
}

var BKempten肯普滕 feud.Barony = &肯普滕KemptenBarony{}

func init() {
	f := BKempten肯普滕.(*肯普滕KemptenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kempten",
		TitleName: "肯普滕",
		TitleCode: "b_kempten",
	}
}
