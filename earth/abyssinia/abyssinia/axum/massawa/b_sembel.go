package massawa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟姆贝SembelBarony struct {
	feud.BaseBarony
}

var BSembel瑟姆贝 feud.Barony = &瑟姆贝SembelBarony{}

func init() {
    f := BSembel瑟姆贝.(*瑟姆贝SembelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sembel",
		TitleName: "瑟姆贝",
		TitleCode: "b_sembel",
	}
}
