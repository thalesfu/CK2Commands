package litomerice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利托梅日采LitomericeBarony struct {
	feud.BaseBarony
}

var BLitomerice利托梅日采 feud.Barony = &利托梅日采LitomericeBarony{}

func init() {
    f := BLitomerice利托梅日采.(*利托梅日采LitomericeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "litomerice",
		TitleName: "利托梅日采",
		TitleCode: "b_litomerice",
	}
}
