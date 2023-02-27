package akhmim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 居诺波利斯CynopolisBarony struct {
	feud.BaseBarony
}

var BCynopolis居诺波利斯 feud.Barony = &居诺波利斯CynopolisBarony{}

func init() {
    f := BCynopolis居诺波利斯.(*居诺波利斯CynopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cynopolis",
		TitleName: "居诺波利斯",
		TitleCode: "b_cynopolis",
	}
}
