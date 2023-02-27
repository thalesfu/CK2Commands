package ascalon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博赫梅BothmeBarony struct {
	feud.BaseBarony
}

var BBothme博赫梅 feud.Barony = &博赫梅BothmeBarony{}

func init() {
    f := BBothme博赫梅.(*博赫梅BothmeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bothme",
		TitleName: "博赫梅",
		TitleCode: "b_bothme",
	}
}
