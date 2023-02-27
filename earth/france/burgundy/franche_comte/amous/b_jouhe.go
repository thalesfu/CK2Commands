package amous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茹埃JouheBarony struct {
	feud.BaseBarony
}

var BJouhe茹埃 feud.Barony = &茹埃JouheBarony{}

func init() {
    f := BJouhe茹埃.(*茹埃JouheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jouhe",
		TitleName: "茹埃",
		TitleCode: "b_jouhe",
	}
}
