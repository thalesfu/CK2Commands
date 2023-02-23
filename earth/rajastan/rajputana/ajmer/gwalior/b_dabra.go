package gwalior

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德布拉DabraBarony struct {
	feud.BaseBarony
}

var BDabra德布拉 feud.Barony = &德布拉DabraBarony{}

func init() {
	f := BDabra德布拉.(*德布拉DabraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dabra",
		TitleName: "德布拉",
		TitleCode: "b_dabra",
	}
}
