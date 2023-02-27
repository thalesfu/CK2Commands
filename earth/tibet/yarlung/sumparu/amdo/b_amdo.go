package amdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安多AmdoBarony struct {
	feud.BaseBarony
}

var BAmdo安多 feud.Barony = &安多AmdoBarony{}

func init() {
    f := BAmdo安多.(*安多AmdoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amdo",
		TitleName: "安多",
		TitleCode: "b_amdo",
	}
}
