package gwalior

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗阇耶补罗VijaypurBarony struct {
	feud.BaseBarony
}

var BVijaypur毗阇耶补罗 feud.Barony = &毗阇耶补罗VijaypurBarony{}

func init() {
    f := BVijaypur毗阇耶补罗.(*毗阇耶补罗VijaypurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vijaypur",
		TitleName: "毗阇耶补罗",
		TitleCode: "b_vijaypur",
	}
}
