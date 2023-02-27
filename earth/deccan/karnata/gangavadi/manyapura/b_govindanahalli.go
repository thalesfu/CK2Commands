package manyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿频陀纳诃梨GovindanahalliBarony struct {
	feud.BaseBarony
}

var BGovindanahalli瞿频陀纳诃梨 feud.Barony = &瞿频陀纳诃梨GovindanahalliBarony{}

func init() {
    f := BGovindanahalli瞿频陀纳诃梨.(*瞿频陀纳诃梨GovindanahalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "govindanahalli",
		TitleName: "瞿频陀纳诃梨",
		TitleCode: "b_govindanahalli",
	}
}
