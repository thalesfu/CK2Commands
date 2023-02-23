package charolais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟米尔昂布里奥奈SemurenbrionnaisBarony struct {
	feud.BaseBarony
}

var BSemurenbrionnais瑟米尔昂布里奥奈 feud.Barony = &瑟米尔昂布里奥奈SemurenbrionnaisBarony{}

func init() {
	f := BSemurenbrionnais瑟米尔昂布里奥奈.(*瑟米尔昂布里奥奈SemurenbrionnaisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "semurenbrionnais",
		TitleName: "瑟米尔昂布里奥奈",
		TitleCode: "b_semurenbrionnais",
	}
}
