package udabhanda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌铎迦汉荼UdabhandaBarony struct {
	feud.BaseBarony
}

var BUdabhanda乌铎迦汉荼 feud.Barony = &乌铎迦汉荼UdabhandaBarony{}

func init() {
	f := BUdabhanda乌铎迦汉荼.(*乌铎迦汉荼UdabhandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udabhanda",
		TitleName: "乌铎迦汉荼",
		TitleCode: "b_udabhanda",
	}
}
