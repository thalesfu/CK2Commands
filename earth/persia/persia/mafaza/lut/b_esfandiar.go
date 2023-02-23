package lut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯凡迪亚尔EsfandiarBarony struct {
	feud.BaseBarony
}

var BEsfandiar埃斯凡迪亚尔 feud.Barony = &埃斯凡迪亚尔EsfandiarBarony{}

func init() {
	f := BEsfandiar埃斯凡迪亚尔.(*埃斯凡迪亚尔EsfandiarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esfandiar",
		TitleName: "埃斯凡迪亚尔",
		TitleCode: "b_esfandiar",
	}
}
