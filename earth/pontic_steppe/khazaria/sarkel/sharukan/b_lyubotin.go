package sharukan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳博京LyubotinBarony struct {
	feud.BaseBarony
}

var BLyubotin柳博京 feud.Barony = &柳博京LyubotinBarony{}

func init() {
    f := BLyubotin柳博京.(*柳博京LyubotinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyubotin",
		TitleName: "柳博京",
		TitleCode: "b_lyubotin",
	}
}
