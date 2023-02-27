package bayda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈班HabbanBarony struct {
	feud.BaseBarony
}

var BHabban哈班 feud.Barony = &哈班HabbanBarony{}

func init() {
    f := BHabban哈班.(*哈班HabbanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "habban",
		TitleName: "哈班",
		TitleCode: "b_habban",
	}
}
