package valladolid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ValladolidCounty interface {
	feud.County
	BAvila阿维拉() feud.Barony
	BIscar伊斯卡尔() feud.Barony
	BMedinadelcampo坎波城() feud.Barony
	BPenafiel佩尼亚菲耶尔() feud.Barony
	BSegovia塞哥维亚() feud.Barony
	BSimancas锡曼卡斯() feud.Barony
	BTordesillas托德西利亚斯() feud.Barony
	BValladolid巴利亚多利德() feud.Barony
}

type 巴利亚多利德ValladolidCounty struct {
	feud.BaseCounty
	阿维拉Avila          feud.Barony
	伊斯卡尔Iscar         feud.Barony
	坎波城Medinadelcampo feud.Barony
	佩尼亚菲耶尔Penafiel    feud.Barony
	塞哥维亚Segovia       feud.Barony
	锡曼卡斯Simancas      feud.Barony
	托德西利亚斯Tordesillas feud.Barony
	巴利亚多利德Valladolid  feud.Barony
}

func (c *巴利亚多利德ValladolidCounty) BAvila阿维拉() feud.Barony {
	return c.阿维拉Avila
}

func (c *巴利亚多利德ValladolidCounty) BIscar伊斯卡尔() feud.Barony {
	return c.伊斯卡尔Iscar
}

func (c *巴利亚多利德ValladolidCounty) BMedinadelcampo坎波城() feud.Barony {
	return c.坎波城Medinadelcampo
}

func (c *巴利亚多利德ValladolidCounty) BPenafiel佩尼亚菲耶尔() feud.Barony {
	return c.佩尼亚菲耶尔Penafiel
}

func (c *巴利亚多利德ValladolidCounty) BSegovia塞哥维亚() feud.Barony {
	return c.塞哥维亚Segovia
}

func (c *巴利亚多利德ValladolidCounty) BSimancas锡曼卡斯() feud.Barony {
	return c.锡曼卡斯Simancas
}

func (c *巴利亚多利德ValladolidCounty) BTordesillas托德西利亚斯() feud.Barony {
	return c.托德西利亚斯Tordesillas
}

func (c *巴利亚多利德ValladolidCounty) BValladolid巴利亚多利德() feud.Barony {
	return c.巴利亚多利德Valladolid
}

var CValladolid巴利亚多利德 ValladolidCounty = &巴利亚多利德ValladolidCounty{}

func init() {
	f := CValladolid巴利亚多利德.(*巴利亚多利德ValladolidCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "198",
		Title:     "valladolid",
		TitleName: "巴利亚多利德",
		TitleCode: "c_valladolid",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿维拉Avila = BAvila阿维拉
	f.阿维拉Avila.SetParent(f)

	f.伊斯卡尔Iscar = BIscar伊斯卡尔
	f.伊斯卡尔Iscar.SetParent(f)

	f.坎波城Medinadelcampo = BMedinadelcampo坎波城
	f.坎波城Medinadelcampo.SetParent(f)

	f.佩尼亚菲耶尔Penafiel = BPenafiel佩尼亚菲耶尔
	f.佩尼亚菲耶尔Penafiel.SetParent(f)

	f.塞哥维亚Segovia = BSegovia塞哥维亚
	f.塞哥维亚Segovia.SetParent(f)

	f.锡曼卡斯Simancas = BSimancas锡曼卡斯
	f.锡曼卡斯Simancas.SetParent(f)

	f.托德西利亚斯Tordesillas = BTordesillas托德西利亚斯
	f.托德西利亚斯Tordesillas.SetParent(f)

	f.巴利亚多利德Valladolid = BValladolid巴利亚多利德
	f.巴利亚多利德Valladolid.SetParent(f)

}
