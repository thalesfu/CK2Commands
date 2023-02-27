package amous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣伊利SaintylieBarony struct {
	feud.BaseBarony
}

var BSaintylie圣伊利 feud.Barony = &圣伊利SaintylieBarony{}

func init() {
    f := BSaintylie圣伊利.(*圣伊利SaintylieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintylie",
		TitleName: "圣伊利",
		TitleCode: "b_saintylie",
	}
}
