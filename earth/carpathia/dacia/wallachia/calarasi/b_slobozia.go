package calarasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯洛博齐亚SloboziaBarony struct {
	feud.BaseBarony
}

var BSlobozia斯洛博齐亚 feud.Barony = &斯洛博齐亚SloboziaBarony{}

func init() {
    f := BSlobozia斯洛博齐亚.(*斯洛博齐亚SloboziaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slobozia",
		TitleName: "斯洛博齐亚",
		TitleCode: "b_slobozia",
	}
}
