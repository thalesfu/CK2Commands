package starodub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰斯尼扬西克DesnianskeBarony struct {
	feud.BaseBarony
}

var BDesnianske杰斯尼扬西克 feud.Barony = &杰斯尼扬西克DesnianskeBarony{}

func init() {
    f := BDesnianske杰斯尼扬西克.(*杰斯尼扬西克DesnianskeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "desnianske",
		TitleName: "杰斯尼扬西克",
		TitleCode: "b_desnianske",
	}
}
