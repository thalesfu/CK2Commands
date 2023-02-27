package tyrnovo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊里诺波利斯IrinopolisBarony struct {
	feud.BaseBarony
}

var BIrinopolis伊里诺波利斯 feud.Barony = &伊里诺波利斯IrinopolisBarony{}

func init() {
    f := BIrinopolis伊里诺波利斯.(*伊里诺波利斯IrinopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irinopolis",
		TitleName: "伊里诺波利斯",
		TitleCode: "b_irinopolis",
	}
}
