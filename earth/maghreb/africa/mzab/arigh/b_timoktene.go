package arigh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂莫克滕TimokteneBarony struct {
	feud.BaseBarony
}

var BTimoktene蒂莫克滕 feud.Barony = &蒂莫克滕TimokteneBarony{}

func init() {
	f := BTimoktene蒂莫克滕.(*蒂莫克滕TimokteneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "timoktene",
		TitleName: "蒂莫克滕",
		TitleCode: "b_timoktene",
	}
}
