package kanara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿罗槃迦株AruvankaduBarony struct {
	feud.BaseBarony
}

var BAruvankadu阿罗槃迦株 feud.Barony = &阿罗槃迦株AruvankaduBarony{}

func init() {
    f := BAruvankadu阿罗槃迦株.(*阿罗槃迦株AruvankaduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aruvankadu",
		TitleName: "阿罗槃迦株",
		TitleCode: "b_aruvankadu",
	}
}
