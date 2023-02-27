package samatata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽罗摩梨GarmaliBarony struct {
	feud.BaseBarony
}

var BGarmali伽罗摩梨 feud.Barony = &伽罗摩梨GarmaliBarony{}

func init() {
    f := BGarmali伽罗摩梨.(*伽罗摩梨GarmaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garmali",
		TitleName: "伽罗摩梨",
		TitleCode: "b_garmali",
	}
}
