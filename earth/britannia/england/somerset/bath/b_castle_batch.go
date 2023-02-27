package bath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴奇Castle_batchBarony struct {
	feud.BaseBarony
}

var BCastle_batch巴奇 feud.Barony = &巴奇Castle_batchBarony{}

func init() {
    f := BCastle_batch巴奇.(*巴奇Castle_batchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castle_batch",
		TitleName: "巴奇",
		TitleCode: "b_castle_batch",
	}
}
