package safed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多隆ToronBarony struct {
	feud.BaseBarony
}

var BToron多隆 feud.Barony = &多隆ToronBarony{}

func init() {
	f := BToron多隆.(*多隆ToronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toron",
		TitleName: "多隆",
		TitleCode: "b_toron",
	}
}
