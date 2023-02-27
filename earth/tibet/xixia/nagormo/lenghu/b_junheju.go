package lenghu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 君和居JunhejuBarony struct {
	feud.BaseBarony
}

var BJunheju君和居 feud.Barony = &君和居JunhejuBarony{}

func init() {
    f := BJunheju君和居.(*君和居JunhejuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "junheju",
		TitleName: "君和居",
		TitleCode: "b_junheju",
	}
}
