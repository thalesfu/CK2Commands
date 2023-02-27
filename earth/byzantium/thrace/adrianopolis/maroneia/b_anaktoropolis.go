package maroneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿纳克托罗波利斯AnaktoropolisBarony struct {
	feud.BaseBarony
}

var BAnaktoropolis阿纳克托罗波利斯 feud.Barony = &阿纳克托罗波利斯AnaktoropolisBarony{}

func init() {
    f := BAnaktoropolis阿纳克托罗波利斯.(*阿纳克托罗波利斯AnaktoropolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anaktoropolis",
		TitleName: "阿纳克托罗波利斯",
		TitleCode: "b_anaktoropolis",
	}
}
