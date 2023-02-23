package nakhshab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜坎德PaykandBarony struct {
	feud.BaseBarony
}

var BPaykand拜坎德 feud.Barony = &拜坎德PaykandBarony{}

func init() {
	f := BPaykand拜坎德.(*拜坎德PaykandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paykand",
		TitleName: "拜坎德",
		TitleCode: "b_paykand",
	}
}
