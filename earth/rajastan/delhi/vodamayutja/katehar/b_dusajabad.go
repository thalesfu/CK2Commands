package katehar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菟娑阇跋陀DusajabadBarony struct {
	feud.BaseBarony
}

var BDusajabad菟娑阇跋陀 feud.Barony = &菟娑阇跋陀DusajabadBarony{}

func init() {
	f := BDusajabad菟娑阇跋陀.(*菟娑阇跋陀DusajabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dusajabad",
		TitleName: "菟娑阇跋陀",
		TitleCode: "b_dusajabad",
	}
}
