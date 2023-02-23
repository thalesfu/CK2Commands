package hama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞吉拉SerjillaBarony struct {
	feud.BaseBarony
}

var BSerjilla塞吉拉 feud.Barony = &塞吉拉SerjillaBarony{}

func init() {
	f := BSerjilla塞吉拉.(*塞吉拉SerjillaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serjilla",
		TitleName: "塞吉拉",
		TitleCode: "b_serjilla",
	}
}
