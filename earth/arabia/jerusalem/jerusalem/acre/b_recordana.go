package acre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷科丹那RecordanaBarony struct {
	feud.BaseBarony
}

var BRecordana雷科丹那 feud.Barony = &雷科丹那RecordanaBarony{}

func init() {
    f := BRecordana雷科丹那.(*雷科丹那RecordanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "recordana",
		TitleName: "雷科丹那",
		TitleCode: "b_recordana",
	}
}
