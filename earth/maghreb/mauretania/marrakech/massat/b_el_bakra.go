package massat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝克雷El_bakraBarony struct {
	feud.BaseBarony
}

var BEl_bakra贝克雷 feud.Barony = &贝克雷El_bakraBarony{}

func init() {
    f := BEl_bakra贝克雷.(*贝克雷El_bakraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_bakra",
		TitleName: "贝克雷",
		TitleCode: "b_el_bakra",
	}
}
