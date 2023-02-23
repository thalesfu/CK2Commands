package boqtybay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔卡什UrkashBarony struct {
	feud.BaseBarony
}

var BUrkash乌尔卡什 feud.Barony = &乌尔卡什UrkashBarony{}

func init() {
	f := BUrkash乌尔卡什.(*乌尔卡什UrkashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urkash",
		TitleName: "乌尔卡什",
		TitleCode: "b_urkash",
	}
}
