package carcassonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米内尔夫MinerveBarony struct {
	feud.BaseBarony
}

var BMinerve米内尔夫 feud.Barony = &米内尔夫MinerveBarony{}

func init() {
	f := BMinerve米内尔夫.(*米内尔夫MinerveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "minerve",
		TitleName: "米内尔夫",
		TitleCode: "b_minerve",
	}
}
