package urgench

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌烈昆UlkunBarony struct {
	feud.BaseBarony
}

var BUlkun乌烈昆 feud.Barony = &乌烈昆UlkunBarony{}

func init() {
    f := BUlkun乌烈昆.(*乌烈昆UlkunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulkun",
		TitleName: "乌烈昆",
		TitleCode: "b_ulkun",
	}
}
