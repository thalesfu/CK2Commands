package trier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维特利希WittlichBarony struct {
	feud.BaseBarony
}

var BWittlich维特利希 feud.Barony = &维特利希WittlichBarony{}

func init() {
    f := BWittlich维特利希.(*维特利希WittlichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wittlich",
		TitleName: "维特利希",
		TitleCode: "b_wittlich",
	}
}
