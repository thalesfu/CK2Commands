package kandalax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮亚奥泽罗PyaozeroBarony struct {
	feud.BaseBarony
}

var BPyaozero皮亚奥泽罗 feud.Barony = &皮亚奥泽罗PyaozeroBarony{}

func init() {
    f := BPyaozero皮亚奥泽罗.(*皮亚奥泽罗PyaozeroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pyaozero",
		TitleName: "皮亚奥泽罗",
		TitleCode: "b_pyaozero",
	}
}
