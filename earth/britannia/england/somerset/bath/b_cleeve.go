package bath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克利夫CleeveBarony struct {
	feud.BaseBarony
}

var BCleeve克利夫 feud.Barony = &克利夫CleeveBarony{}

func init() {
    f := BCleeve克利夫.(*克利夫CleeveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cleeve",
		TitleName: "克利夫",
		TitleCode: "b_cleeve",
	}
}
