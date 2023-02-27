package diskit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 地里特ThiritBarony struct {
	feud.BaseBarony
}

var BThirit地里特 feud.Barony = &地里特ThiritBarony{}

func init() {
    f := BThirit地里特.(*地里特ThiritBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thirit",
		TitleName: "地里特",
		TitleCode: "b_thirit",
	}
}
