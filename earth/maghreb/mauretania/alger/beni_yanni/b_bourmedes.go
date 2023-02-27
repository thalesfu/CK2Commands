package beni_yanni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布米尔达斯BourmedesBarony struct {
	feud.BaseBarony
}

var BBourmedes布米尔达斯 feud.Barony = &布米尔达斯BourmedesBarony{}

func init() {
    f := BBourmedes布米尔达斯.(*布米尔达斯BourmedesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bourmedes",
		TitleName: "布米尔达斯",
		TitleCode: "b_bourmedes",
	}
}
