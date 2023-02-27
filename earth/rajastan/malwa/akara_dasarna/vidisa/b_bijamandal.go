package vidisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗阇曼荼罗BijamandalBarony struct {
	feud.BaseBarony
}

var BBijamandal毗阇曼荼罗 feud.Barony = &毗阇曼荼罗BijamandalBarony{}

func init() {
    f := BBijamandal毗阇曼荼罗.(*毗阇曼荼罗BijamandalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bijamandal",
		TitleName: "毗阇曼荼罗",
		TitleCode: "b_bijamandal",
	}
}
