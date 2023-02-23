package sarqihya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SarqihyaCounty interface {
	feud.County
	BAgruda阿鲁达() feud.Barony
	BAlhaifar海法尔() feud.Barony
	BBubastis布巴斯提斯() feud.Barony
	BMinya明亚() feud.Barony
	BQantir坎提尔() feud.Barony
	BSarqinya比勒拜斯() feud.Barony
	BZagazig宰加济格() feud.Barony
}

type 比勒拜斯SarqihyaCounty struct {
	feud.BaseCounty
	阿鲁达Agruda     feud.Barony
	海法尔Alhaifar   feud.Barony
	布巴斯提斯Bubastis feud.Barony
	明亚Minya       feud.Barony
	坎提尔Qantir     feud.Barony
	比勒拜斯Sarqinya  feud.Barony
	宰加济格Zagazig   feud.Barony
}

func (c *比勒拜斯SarqihyaCounty) BAgruda阿鲁达() feud.Barony {
	return c.阿鲁达Agruda
}

func (c *比勒拜斯SarqihyaCounty) BAlhaifar海法尔() feud.Barony {
	return c.海法尔Alhaifar
}

func (c *比勒拜斯SarqihyaCounty) BBubastis布巴斯提斯() feud.Barony {
	return c.布巴斯提斯Bubastis
}

func (c *比勒拜斯SarqihyaCounty) BMinya明亚() feud.Barony {
	return c.明亚Minya
}

func (c *比勒拜斯SarqihyaCounty) BQantir坎提尔() feud.Barony {
	return c.坎提尔Qantir
}

func (c *比勒拜斯SarqihyaCounty) BSarqinya比勒拜斯() feud.Barony {
	return c.比勒拜斯Sarqinya
}

func (c *比勒拜斯SarqihyaCounty) BZagazig宰加济格() feud.Barony {
	return c.宰加济格Zagazig
}

var CSarqihya比勒拜斯 SarqihyaCounty = &比勒拜斯SarqihyaCounty{}

func init() {
	f := CSarqihya比勒拜斯.(*比勒拜斯SarqihyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "790",
		Title:     "sarqihya",
		TitleName: "比勒拜斯",
		TitleCode: "c_sarqihya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿鲁达Agruda = BAgruda阿鲁达
	f.阿鲁达Agruda.SetParent(f)

	f.海法尔Alhaifar = BAlhaifar海法尔
	f.海法尔Alhaifar.SetParent(f)

	f.布巴斯提斯Bubastis = BBubastis布巴斯提斯
	f.布巴斯提斯Bubastis.SetParent(f)

	f.明亚Minya = BMinya明亚
	f.明亚Minya.SetParent(f)

	f.坎提尔Qantir = BQantir坎提尔
	f.坎提尔Qantir.SetParent(f)

	f.比勒拜斯Sarqinya = BSarqinya比勒拜斯
	f.比勒拜斯Sarqinya.SetParent(f)

	f.宰加济格Zagazig = BZagazig宰加济格
	f.宰加济格Zagazig.SetParent(f)

}
