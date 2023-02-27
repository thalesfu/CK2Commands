package atlas_mnt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因斯马拉AinzmilaBarony struct {
	feud.BaseBarony
}

var BAinzmila艾因斯马拉 feud.Barony = &艾因斯马拉AinzmilaBarony{}

func init() {
    f := BAinzmila艾因斯马拉.(*艾因斯马拉AinzmilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ainzmila",
		TitleName: "艾因斯马拉",
		TitleCode: "b_ainzmila",
	}
}
