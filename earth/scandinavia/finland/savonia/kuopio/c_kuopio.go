package kuopio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KuopioCounty interface {
	feud.County
	BJuurusvesu尤鲁斯湖() feud.Barony
	BKallavesi卡拉韦西() feud.Barony
	BKotasaari科塔萨里() feud.Barony
	BKukkarinselka库卡林湖() feud.Barony
	BKuopio库奥皮奥() feud.Barony
	BRuokovesi鲁奥韦西() feud.Barony
	BVaajasalo瓦亚萨洛() feud.Barony
}

type 卡拉韦西KuopioCounty struct {
	feud.BaseCounty
	尤鲁斯湖Juurusvesu    feud.Barony
	卡拉韦西Kallavesi     feud.Barony
	科塔萨里Kotasaari     feud.Barony
	库卡林湖Kukkarinselka feud.Barony
	库奥皮奥Kuopio        feud.Barony
	鲁奥韦西Ruokovesi     feud.Barony
	瓦亚萨洛Vaajasalo     feud.Barony
}

func (c *卡拉韦西KuopioCounty) BJuurusvesu尤鲁斯湖() feud.Barony {
	return c.尤鲁斯湖Juurusvesu
}

func (c *卡拉韦西KuopioCounty) BKallavesi卡拉韦西() feud.Barony {
	return c.卡拉韦西Kallavesi
}

func (c *卡拉韦西KuopioCounty) BKotasaari科塔萨里() feud.Barony {
	return c.科塔萨里Kotasaari
}

func (c *卡拉韦西KuopioCounty) BKukkarinselka库卡林湖() feud.Barony {
	return c.库卡林湖Kukkarinselka
}

func (c *卡拉韦西KuopioCounty) BKuopio库奥皮奥() feud.Barony {
	return c.库奥皮奥Kuopio
}

func (c *卡拉韦西KuopioCounty) BRuokovesi鲁奥韦西() feud.Barony {
	return c.鲁奥韦西Ruokovesi
}

func (c *卡拉韦西KuopioCounty) BVaajasalo瓦亚萨洛() feud.Barony {
	return c.瓦亚萨洛Vaajasalo
}

var CKuopio卡拉韦西 KuopioCounty = &卡拉韦西KuopioCounty{}

func init() {
	f := CKuopio卡拉韦西.(*卡拉韦西KuopioCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1605",
		Title:     "kuopio",
		TitleName: "卡拉韦西",
		TitleCode: "c_kuopio",
		Baronies:  map[string]feud.Barony{},
	}

	f.尤鲁斯湖Juurusvesu = BJuurusvesu尤鲁斯湖
	f.尤鲁斯湖Juurusvesu.SetParent(f)

	f.卡拉韦西Kallavesi = BKallavesi卡拉韦西
	f.卡拉韦西Kallavesi.SetParent(f)

	f.科塔萨里Kotasaari = BKotasaari科塔萨里
	f.科塔萨里Kotasaari.SetParent(f)

	f.库卡林湖Kukkarinselka = BKukkarinselka库卡林湖
	f.库卡林湖Kukkarinselka.SetParent(f)

	f.库奥皮奥Kuopio = BKuopio库奥皮奥
	f.库奥皮奥Kuopio.SetParent(f)

	f.鲁奥韦西Ruokovesi = BRuokovesi鲁奥韦西
	f.鲁奥韦西Ruokovesi.SetParent(f)

	f.瓦亚萨洛Vaajasalo = BVaajasalo瓦亚萨洛
	f.瓦亚萨洛Vaajasalo.SetParent(f)

}
