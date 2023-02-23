package nasikya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 底哩阎婆迦TryambakaBarony struct {
	feud.BaseBarony
}

var BTryambaka底哩阎婆迦 feud.Barony = &底哩阎婆迦TryambakaBarony{}

func init() {
	f := BTryambaka底哩阎婆迦.(*底哩阎婆迦TryambakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tryambaka",
		TitleName: "底哩阎婆迦",
		TitleCode: "b_tryambaka",
	}
}
