package yangikent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 养吉干YangikentBarony struct {
	feud.BaseBarony
}

var BYangikent养吉干 feud.Barony = &养吉干YangikentBarony{}

func init() {
    f := BYangikent养吉干.(*养吉干YangikentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yangikent",
		TitleName: "养吉干",
		TitleCode: "b_yangikent",
	}
}
