package dijon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟米尔SemurBarony struct {
	feud.BaseBarony
}

var BSemur瑟米尔 feud.Barony = &瑟米尔SemurBarony{}

func init() {
	f := BSemur瑟米尔.(*瑟米尔SemurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "semur",
		TitleName: "瑟米尔",
		TitleCode: "b_semur",
	}
}
