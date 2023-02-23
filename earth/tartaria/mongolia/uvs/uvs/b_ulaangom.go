package uvs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌兰固木UlaangomBarony struct {
	feud.BaseBarony
}

var BUlaangom乌兰固木 feud.Barony = &乌兰固木UlaangomBarony{}

func init() {
	f := BUlaangom乌兰固木.(*乌兰固木UlaangomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulaangom",
		TitleName: "乌兰固木",
		TitleCode: "b_ulaangom",
	}
}
