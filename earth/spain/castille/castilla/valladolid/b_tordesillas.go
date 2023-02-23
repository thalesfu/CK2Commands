package valladolid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托德西利亚斯TordesillasBarony struct {
	feud.BaseBarony
}

var BTordesillas托德西利亚斯 feud.Barony = &托德西利亚斯TordesillasBarony{}

func init() {
	f := BTordesillas托德西利亚斯.(*托德西利亚斯TordesillasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tordesillas",
		TitleName: "托德西利亚斯",
		TitleCode: "b_tordesillas",
	}
}
