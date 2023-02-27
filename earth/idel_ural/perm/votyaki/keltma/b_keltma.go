package keltma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克利特马KeltmaBarony struct {
	feud.BaseBarony
}

var BKeltma克利特马 feud.Barony = &克利特马KeltmaBarony{}

func init() {
    f := BKeltma克利特马.(*克利特马KeltmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keltma",
		TitleName: "克利特马",
		TitleCode: "b_keltma",
	}
}
