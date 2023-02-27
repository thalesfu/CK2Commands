package birjand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔詹德BirjandBarony struct {
	feud.BaseBarony
}

var BBirjand比尔詹德 feud.Barony = &比尔詹德BirjandBarony{}

func init() {
    f := BBirjand比尔詹德.(*比尔詹德BirjandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birjand",
		TitleName: "比尔詹德",
		TitleCode: "b_birjand",
	}
}
