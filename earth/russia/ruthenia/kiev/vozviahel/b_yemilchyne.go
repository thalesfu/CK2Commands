package vozviahel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃米利奇涅YemilchyneBarony struct {
	feud.BaseBarony
}

var BYemilchyne埃米利奇涅 feud.Barony = &埃米利奇涅YemilchyneBarony{}

func init() {
    f := BYemilchyne埃米利奇涅.(*埃米利奇涅YemilchyneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yemilchyne",
		TitleName: "埃米利奇涅",
		TitleCode: "b_yemilchyne",
	}
}
