package rimini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加泰奥GatteoBarony struct {
	feud.BaseBarony
}

var BGatteo加泰奥 feud.Barony = &加泰奥GatteoBarony{}

func init() {
	f := BGatteo加泰奥.(*加泰奥GatteoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gatteo",
		TitleName: "加泰奥",
		TitleCode: "b_gatteo",
	}
}
