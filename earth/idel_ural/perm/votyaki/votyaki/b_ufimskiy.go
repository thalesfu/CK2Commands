package votyaki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌菲姆斯基UfimskiyBarony struct {
	feud.BaseBarony
}

var BUfimskiy乌菲姆斯基 feud.Barony = &乌菲姆斯基UfimskiyBarony{}

func init() {
    f := BUfimskiy乌菲姆斯基.(*乌菲姆斯基UfimskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ufimskiy",
		TitleName: "乌菲姆斯基",
		TitleCode: "b_ufimskiy",
	}
}
