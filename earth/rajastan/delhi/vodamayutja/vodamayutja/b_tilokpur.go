package vodamayutja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝卢迦补罗TilokpurBarony struct {
	feud.BaseBarony
}

var BTilokpur帝卢迦补罗 feud.Barony = &帝卢迦补罗TilokpurBarony{}

func init() {
	f := BTilokpur帝卢迦补罗.(*帝卢迦补罗TilokpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tilokpur",
		TitleName: "帝卢迦补罗",
		TitleCode: "b_tilokpur",
	}
}
