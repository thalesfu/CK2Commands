package forcalquier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加普GapBarony struct {
	feud.BaseBarony
}

var BGap加普 feud.Barony = &加普GapBarony{}

func init() {
    f := BGap加普.(*加普GapBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gap",
		TitleName: "加普",
		TitleCode: "b_gap",
	}
}
