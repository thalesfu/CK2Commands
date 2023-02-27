package havelberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维特施托克WittstockBarony struct {
	feud.BaseBarony
}

var BWittstock维特施托克 feud.Barony = &维特施托克WittstockBarony{}

func init() {
    f := BWittstock维特施托克.(*维特施托克WittstockBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wittstock",
		TitleName: "维特施托克",
		TitleCode: "b_wittstock",
	}
}
