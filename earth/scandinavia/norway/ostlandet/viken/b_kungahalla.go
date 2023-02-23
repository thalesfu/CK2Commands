package viken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库耶赫拉KungahallaBarony struct {
	feud.BaseBarony
}

var BKungahalla库耶赫拉 feud.Barony = &库耶赫拉KungahallaBarony{}

func init() {
	f := BKungahalla库耶赫拉.(*库耶赫拉KungahallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kungahalla",
		TitleName: "库耶赫拉",
		TitleCode: "b_kungahalla",
	}
}
