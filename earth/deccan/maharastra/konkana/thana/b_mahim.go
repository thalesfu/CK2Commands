package thana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马哈音MahimBarony struct {
	feud.BaseBarony
}

var BMahim马哈音 feud.Barony = &马哈音MahimBarony{}

func init() {
    f := BMahim马哈音.(*马哈音MahimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahim",
		TitleName: "马哈音",
		TitleCode: "b_mahim",
	}
}
