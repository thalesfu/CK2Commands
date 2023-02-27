package diskit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德吉DiskitBarony struct {
	feud.BaseBarony
}

var BDiskit德吉 feud.Barony = &德吉DiskitBarony{}

func init() {
    f := BDiskit德吉.(*德吉DiskitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diskit",
		TitleName: "德吉",
		TitleCode: "b_diskit",
	}
}
