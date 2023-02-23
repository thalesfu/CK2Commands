package cakrakuta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富波婆菟吉PurbbadulkiBarony struct {
	feud.BaseBarony
}

var BPurbbadulki富波婆菟吉 feud.Barony = &富波婆菟吉PurbbadulkiBarony{}

func init() {
	f := BPurbbadulki富波婆菟吉.(*富波婆菟吉PurbbadulkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "purbbadulki",
		TitleName: "富波婆菟吉",
		TitleCode: "b_purbbadulki",
	}
}
