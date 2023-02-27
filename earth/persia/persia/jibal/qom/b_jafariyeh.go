package qom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾法里耶赫JafariyehBarony struct {
	feud.BaseBarony
}

var BJafariyeh贾法里耶赫 feud.Barony = &贾法里耶赫JafariyehBarony{}

func init() {
    f := BJafariyeh贾法里耶赫.(*贾法里耶赫JafariyehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jafariyeh",
		TitleName: "贾法里耶赫",
		TitleCode: "b_jafariyeh",
	}
}
