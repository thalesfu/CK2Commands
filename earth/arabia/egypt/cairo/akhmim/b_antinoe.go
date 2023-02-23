package akhmim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安提倪AntinoeBarony struct {
	feud.BaseBarony
}

var BAntinoe安提倪 feud.Barony = &安提倪AntinoeBarony{}

func init() {
	f := BAntinoe安提倪.(*安提倪AntinoeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antinoe",
		TitleName: "安提倪",
		TitleCode: "b_antinoe",
	}
}
