package sarysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库姆肯特KumkentBarony struct {
	feud.BaseBarony
}

var BKumkent库姆肯特 feud.Barony = &库姆肯特KumkentBarony{}

func init() {
	f := BKumkent库姆肯特.(*库姆肯特KumkentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumkent",
		TitleName: "库姆肯特",
		TitleCode: "b_kumkent",
	}
}
