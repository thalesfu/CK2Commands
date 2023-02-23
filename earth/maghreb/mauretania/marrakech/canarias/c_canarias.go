package canarias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CanariasCounty interface {
	feud.County
	BAlegranza阿莱格兰萨() feud.Barony
	BFuerteventura富埃特文图拉() feud.Barony
	BLagomera戈梅拉() feud.Barony
	BLagraciosa格拉西奥萨() feud.Barony
	BLanzarote兰萨罗特() feud.Barony
	BLapalma拉帕尔马() feud.Barony
	BTenerife特内里费() feud.Barony
}

type 加那利CanariasCounty struct {
	feud.BaseCounty
	阿莱格兰萨Alegranza      feud.Barony
	富埃特文图拉Fuerteventura feud.Barony
	戈梅拉Lagomera         feud.Barony
	格拉西奥萨Lagraciosa     feud.Barony
	兰萨罗特Lanzarote       feud.Barony
	拉帕尔马Lapalma         feud.Barony
	特内里费Tenerife        feud.Barony
}

func (c *加那利CanariasCounty) BAlegranza阿莱格兰萨() feud.Barony {
	return c.阿莱格兰萨Alegranza
}

func (c *加那利CanariasCounty) BFuerteventura富埃特文图拉() feud.Barony {
	return c.富埃特文图拉Fuerteventura
}

func (c *加那利CanariasCounty) BLagomera戈梅拉() feud.Barony {
	return c.戈梅拉Lagomera
}

func (c *加那利CanariasCounty) BLagraciosa格拉西奥萨() feud.Barony {
	return c.格拉西奥萨Lagraciosa
}

func (c *加那利CanariasCounty) BLanzarote兰萨罗特() feud.Barony {
	return c.兰萨罗特Lanzarote
}

func (c *加那利CanariasCounty) BLapalma拉帕尔马() feud.Barony {
	return c.拉帕尔马Lapalma
}

func (c *加那利CanariasCounty) BTenerife特内里费() feud.Barony {
	return c.特内里费Tenerife
}

var CCanarias加那利 CanariasCounty = &加那利CanariasCounty{}

func init() {
	f := CCanarias加那利.(*加那利CanariasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "849",
		Title:     "canarias",
		TitleName: "加那利",
		TitleCode: "c_canarias",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿莱格兰萨Alegranza = BAlegranza阿莱格兰萨
	f.阿莱格兰萨Alegranza.SetParent(f)

	f.富埃特文图拉Fuerteventura = BFuerteventura富埃特文图拉
	f.富埃特文图拉Fuerteventura.SetParent(f)

	f.戈梅拉Lagomera = BLagomera戈梅拉
	f.戈梅拉Lagomera.SetParent(f)

	f.格拉西奥萨Lagraciosa = BLagraciosa格拉西奥萨
	f.格拉西奥萨Lagraciosa.SetParent(f)

	f.兰萨罗特Lanzarote = BLanzarote兰萨罗特
	f.兰萨罗特Lanzarote.SetParent(f)

	f.拉帕尔马Lapalma = BLapalma拉帕尔马
	f.拉帕尔马Lapalma.SetParent(f)

	f.特内里费Tenerife = BTenerife特内里费
	f.特内里费Tenerife.SetParent(f)

}
