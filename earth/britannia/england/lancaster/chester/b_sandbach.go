package chester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑德巴奇SandbachBarony struct {
	feud.BaseBarony
}

var BSandbach桑德巴奇 feud.Barony = &桑德巴奇SandbachBarony{}

func init() {
	f := BSandbach桑德巴奇.(*桑德巴奇SandbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sandbach",
		TitleName: "桑德巴奇",
		TitleCode: "b_sandbach",
	}
}
