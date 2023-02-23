package naumadal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 廷海于根TinghaugenBarony struct {
	feud.BaseBarony
}

var BTinghaugen廷海于根 feud.Barony = &廷海于根TinghaugenBarony{}

func init() {
	f := BTinghaugen廷海于根.(*廷海于根TinghaugenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tinghaugen",
		TitleName: "廷海于根",
		TitleCode: "b_tinghaugen",
	}
}
