package poitiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢丹LoudunBarony struct {
	feud.BaseBarony
}

var BLoudun卢丹 feud.Barony = &卢丹LoudunBarony{}

func init() {
	f := BLoudun卢丹.(*卢丹LoudunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loudun",
		TitleName: "卢丹",
		TitleCode: "b_loudun",
	}
}
