package tobruk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎布特KambutBarony struct {
	feud.BaseBarony
}

var BKambut坎布特 feud.Barony = &坎布特KambutBarony{}

func init() {
    f := BKambut坎布特.(*坎布特KambutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kambut",
		TitleName: "坎布特",
		TitleCode: "b_kambut",
	}
}
