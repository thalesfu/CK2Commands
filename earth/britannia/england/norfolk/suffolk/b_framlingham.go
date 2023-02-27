package suffolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗瑞林姆FramlinghamBarony struct {
	feud.BaseBarony
}

var BFramlingham弗瑞林姆 feud.Barony = &弗瑞林姆FramlinghamBarony{}

func init() {
    f := BFramlingham弗瑞林姆.(*弗瑞林姆FramlinghamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "framlingham",
		TitleName: "弗瑞林姆",
		TitleCode: "b_framlingham",
	}
}
