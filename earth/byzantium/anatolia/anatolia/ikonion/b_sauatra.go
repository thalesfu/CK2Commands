package ikonion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扫阿特拉SauatraBarony struct {
	feud.BaseBarony
}

var BSauatra扫阿特拉 feud.Barony = &扫阿特拉SauatraBarony{}

func init() {
	f := BSauatra扫阿特拉.(*扫阿特拉SauatraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sauatra",
		TitleName: "扫阿特拉",
		TitleCode: "b_sauatra",
	}
}
