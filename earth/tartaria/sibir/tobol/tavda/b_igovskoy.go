package tavda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊戈夫斯科伊IgovskoyBarony struct {
	feud.BaseBarony
}

var BIgovskoy伊戈夫斯科伊 feud.Barony = &伊戈夫斯科伊IgovskoyBarony{}

func init() {
    f := BIgovskoy伊戈夫斯科伊.(*伊戈夫斯科伊IgovskoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "igovskoy",
		TitleName: "伊戈夫斯科伊",
		TitleCode: "b_igovskoy",
	}
}
