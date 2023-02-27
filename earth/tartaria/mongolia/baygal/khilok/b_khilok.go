package khilok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勤勒豁KhilokBarony struct {
	feud.BaseBarony
}

var BKhilok勤勒豁 feud.Barony = &勤勒豁KhilokBarony{}

func init() {
    f := BKhilok勤勒豁.(*勤勒豁KhilokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khilok",
		TitleName: "勤勒豁",
		TitleCode: "b_khilok",
	}
}
