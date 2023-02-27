package roslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊马斯拉夫尔ImaslavlBarony struct {
	feud.BaseBarony
}

var BImaslavl伊马斯拉夫尔 feud.Barony = &伊马斯拉夫尔ImaslavlBarony{}

func init() {
    f := BImaslavl伊马斯拉夫尔.(*伊马斯拉夫尔ImaslavlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "imaslavl",
		TitleName: "伊马斯拉夫尔",
		TitleCode: "b_imaslavl",
	}
}
