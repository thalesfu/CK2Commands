package tukums

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔西TalsiBarony struct {
	feud.BaseBarony
}

var BTalsi塔尔西 feud.Barony = &塔尔西TalsiBarony{}

func init() {
    f := BTalsi塔尔西.(*塔尔西TalsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talsi",
		TitleName: "塔尔西",
		TitleCode: "b_talsi",
	}
}
