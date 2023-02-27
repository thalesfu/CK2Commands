package suvraga_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 臣赫尔TsenherBarony struct {
	feud.BaseBarony
}

var BTsenher臣赫尔 feud.Barony = &臣赫尔TsenherBarony{}

func init() {
    f := BTsenher臣赫尔.(*臣赫尔TsenherBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsenher",
		TitleName: "臣赫尔",
		TitleCode: "b_tsenher",
	}
}
