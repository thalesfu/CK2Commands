package tarbagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌沙拉尔UsharalBarony struct {
	feud.BaseBarony
}

var BUsharal乌沙拉尔 feud.Barony = &乌沙拉尔UsharalBarony{}

func init() {
	f := BUsharal乌沙拉尔.(*乌沙拉尔UsharalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usharal",
		TitleName: "乌沙拉尔",
		TitleCode: "b_usharal",
	}
}
