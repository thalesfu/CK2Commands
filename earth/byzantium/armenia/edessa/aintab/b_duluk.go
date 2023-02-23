package aintab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜鲁克DulukBarony struct {
	feud.BaseBarony
}

var BDuluk杜鲁克 feud.Barony = &杜鲁克DulukBarony{}

func init() {
	f := BDuluk杜鲁克.(*杜鲁克DulukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duluk",
		TitleName: "杜鲁克",
		TitleCode: "b_duluk",
	}
}
