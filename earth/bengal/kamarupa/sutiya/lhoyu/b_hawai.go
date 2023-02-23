package lhoyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈威HawaiBarony struct {
	feud.BaseBarony
}

var BHawai哈威 feud.Barony = &哈威HawaiBarony{}

func init() {
	f := BHawai哈威.(*哈威HawaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hawai",
		TitleName: "哈威",
		TitleCode: "b_hawai",
	}
}
