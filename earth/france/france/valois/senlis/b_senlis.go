package senlis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑利斯SenlisBarony struct {
	feud.BaseBarony
}

var BSenlis桑利斯 feud.Barony = &桑利斯SenlisBarony{}

func init() {
	f := BSenlis桑利斯.(*桑利斯SenlisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "senlis",
		TitleName: "桑利斯",
		TitleCode: "b_senlis",
	}
}
