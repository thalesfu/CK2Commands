package cadota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼雅NiyaBarony struct {
	feud.BaseBarony
}

var BNiya尼雅 feud.Barony = &尼雅NiyaBarony{}

func init() {
    f := BNiya尼雅.(*尼雅NiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niya",
		TitleName: "尼雅",
		TitleCode: "b_niya",
	}
}
