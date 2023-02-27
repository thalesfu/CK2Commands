package holland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 默伊登MuidenBarony struct {
	feud.BaseBarony
}

var BMuiden默伊登 feud.Barony = &默伊登MuidenBarony{}

func init() {
    f := BMuiden默伊登.(*默伊登MuidenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muiden",
		TitleName: "默伊登",
		TitleCode: "b_muiden",
	}
}
