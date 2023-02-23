package taranto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡萨诺CassanoBarony struct {
	feud.BaseBarony
}

var BCassano卡萨诺 feud.Barony = &卡萨诺CassanoBarony{}

func init() {
	f := BCassano卡萨诺.(*卡萨诺CassanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cassano",
		TitleName: "卡萨诺",
		TitleCode: "b_cassano",
	}
}
