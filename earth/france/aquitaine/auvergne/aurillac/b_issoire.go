package aurillac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊苏瓦尔IssoireBarony struct {
	feud.BaseBarony
}

var BIssoire伊苏瓦尔 feud.Barony = &伊苏瓦尔IssoireBarony{}

func init() {
	f := BIssoire伊苏瓦尔.(*伊苏瓦尔IssoireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "issoire",
		TitleName: "伊苏瓦尔",
		TitleCode: "b_issoire",
	}
}
