package gafsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛达斯MadasBarony struct {
	feud.BaseBarony
}

var BMadas玛达斯 feud.Barony = &玛达斯MadasBarony{}

func init() {
    f := BMadas玛达斯.(*玛达斯MadasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madas",
		TitleName: "玛达斯",
		TitleCode: "b_madas",
	}
}
