package kathmandu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗梨多补罗LalitpurBarony struct {
	feud.BaseBarony
}

var BLalitpur罗梨多补罗 feud.Barony = &罗梨多补罗LalitpurBarony{}

func init() {
    f := BLalitpur罗梨多补罗.(*罗梨多补罗LalitpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lalitpur",
		TitleName: "罗梨多补罗",
		TitleCode: "b_lalitpur",
	}
}
