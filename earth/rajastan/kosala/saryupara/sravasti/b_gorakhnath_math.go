package sravasti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿罗叉那他神庙Gorakhnath_mathBarony struct {
	feud.BaseBarony
}

var BGorakhnath_math瞿罗叉那他神庙 feud.Barony = &瞿罗叉那他神庙Gorakhnath_mathBarony{}

func init() {
    f := BGorakhnath_math瞿罗叉那他神庙.(*瞿罗叉那他神庙Gorakhnath_mathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorakhnath_math",
		TitleName: "瞿罗叉那他神庙",
		TitleCode: "b_gorakhnath_math",
	}
}
