package perigord

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧布罗什AuberocheBarony struct {
	feud.BaseBarony
}

var BAuberoche欧布罗什 feud.Barony = &欧布罗什AuberocheBarony{}

func init() {
    f := BAuberoche欧布罗什.(*欧布罗什AuberocheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auberoche",
		TitleName: "欧布罗什",
		TitleCode: "b_auberoche",
	}
}
