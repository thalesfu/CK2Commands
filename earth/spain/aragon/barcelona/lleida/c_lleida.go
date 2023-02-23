package lleida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LleidaCounty interface {
	feud.County
	BAgramunt阿格拉蒙() feud.Barony
	BBalaguer巴拉格尔() feud.Barony
	BBorgesblanques莱斯博尔热斯布兰克斯() feud.Barony
	BCervera塞尔韦拉() feud.Barony
	BLleida列伊达() feud.Barony
	BSolsona索尔索纳() feud.Barony
	BTarrega塔雷加() feud.Barony
	BVerdu贝尔杜() feud.Barony
}

type 列伊达LleidaCounty struct {
	feud.BaseCounty
	阿格拉蒙Agramunt             feud.Barony
	巴拉格尔Balaguer             feud.Barony
	莱斯博尔热斯布兰克斯Borgesblanques feud.Barony
	塞尔韦拉Cervera              feud.Barony
	列伊达Lleida                feud.Barony
	索尔索纳Solsona              feud.Barony
	塔雷加Tarrega               feud.Barony
	贝尔杜Verdu                 feud.Barony
}

func (c *列伊达LleidaCounty) BAgramunt阿格拉蒙() feud.Barony {
	return c.阿格拉蒙Agramunt
}

func (c *列伊达LleidaCounty) BBalaguer巴拉格尔() feud.Barony {
	return c.巴拉格尔Balaguer
}

func (c *列伊达LleidaCounty) BBorgesblanques莱斯博尔热斯布兰克斯() feud.Barony {
	return c.莱斯博尔热斯布兰克斯Borgesblanques
}

func (c *列伊达LleidaCounty) BCervera塞尔韦拉() feud.Barony {
	return c.塞尔韦拉Cervera
}

func (c *列伊达LleidaCounty) BLleida列伊达() feud.Barony {
	return c.列伊达Lleida
}

func (c *列伊达LleidaCounty) BSolsona索尔索纳() feud.Barony {
	return c.索尔索纳Solsona
}

func (c *列伊达LleidaCounty) BTarrega塔雷加() feud.Barony {
	return c.塔雷加Tarrega
}

func (c *列伊达LleidaCounty) BVerdu贝尔杜() feud.Barony {
	return c.贝尔杜Verdu
}

var CLleida列伊达 LleidaCounty = &列伊达LleidaCounty{}

func init() {
	f := CLleida列伊达.(*列伊达LleidaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "203",
		Title:     "lleida",
		TitleName: "列伊达",
		TitleCode: "c_lleida",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格拉蒙Agramunt = BAgramunt阿格拉蒙
	f.阿格拉蒙Agramunt.SetParent(f)

	f.巴拉格尔Balaguer = BBalaguer巴拉格尔
	f.巴拉格尔Balaguer.SetParent(f)

	f.莱斯博尔热斯布兰克斯Borgesblanques = BBorgesblanques莱斯博尔热斯布兰克斯
	f.莱斯博尔热斯布兰克斯Borgesblanques.SetParent(f)

	f.塞尔韦拉Cervera = BCervera塞尔韦拉
	f.塞尔韦拉Cervera.SetParent(f)

	f.列伊达Lleida = BLleida列伊达
	f.列伊达Lleida.SetParent(f)

	f.索尔索纳Solsona = BSolsona索尔索纳
	f.索尔索纳Solsona.SetParent(f)

	f.塔雷加Tarrega = BTarrega塔雷加
	f.塔雷加Tarrega.SetParent(f)

	f.贝尔杜Verdu = BVerdu贝尔杜
	f.贝尔杜Verdu.SetParent(f)

}
