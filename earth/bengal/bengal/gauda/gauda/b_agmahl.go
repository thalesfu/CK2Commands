package gauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿迦摩诃AgmahlBarony struct {
	feud.BaseBarony
}

var BAgmahl阿迦摩诃 feud.Barony = &阿迦摩诃AgmahlBarony{}

func init() {
	f := BAgmahl阿迦摩诃.(*阿迦摩诃AgmahlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agmahl",
		TitleName: "阿迦摩诃",
		TitleCode: "b_agmahl",
	}
}
