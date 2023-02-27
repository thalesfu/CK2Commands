package rottenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥克森富特OchsenfurtBarony struct {
	feud.BaseBarony
}

var BOchsenfurt奥克森富特 feud.Barony = &奥克森富特OchsenfurtBarony{}

func init() {
    f := BOchsenfurt奥克森富特.(*奥克森富特OchsenfurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ochsenfurt",
		TitleName: "奥克森富特",
		TitleCode: "b_ochsenfurt",
	}
}
