package malaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝莱斯_马拉加VelezmalagaBarony struct {
	feud.BaseBarony
}

var BVelezmalaga贝莱斯_马拉加 feud.Barony = &贝莱斯_马拉加VelezmalagaBarony{}

func init() {
    f := BVelezmalaga贝莱斯_马拉加.(*贝莱斯_马拉加VelezmalagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velezmalaga",
		TitleName: "贝莱斯_马拉加",
		TitleCode: "b_velezmalaga",
	}
}
