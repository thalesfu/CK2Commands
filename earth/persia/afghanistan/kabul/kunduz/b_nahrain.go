package kunduz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那赫林NahrainBarony struct {
	feud.BaseBarony
}

var BNahrain那赫林 feud.Barony = &那赫林NahrainBarony{}

func init() {
    f := BNahrain那赫林.(*那赫林NahrainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nahrain",
		TitleName: "那赫林",
		TitleCode: "b_nahrain",
	}
}
