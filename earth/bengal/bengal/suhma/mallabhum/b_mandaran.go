package mallabhum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼陀兰MandaranBarony struct {
	feud.BaseBarony
}

var BMandaran曼陀兰 feud.Barony = &曼陀兰MandaranBarony{}

func init() {
	f := BMandaran曼陀兰.(*曼陀兰MandaranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandaran",
		TitleName: "曼陀兰",
		TitleCode: "b_mandaran",
	}
}
