package diskit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌玛鲁UnmaruBarony struct {
	feud.BaseBarony
}

var BUnmaru乌玛鲁 feud.Barony = &乌玛鲁UnmaruBarony{}

func init() {
    f := BUnmaru乌玛鲁.(*乌玛鲁UnmaruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "unmaru",
		TitleName: "乌玛鲁",
		TitleCode: "b_unmaru",
	}
}
