package buhairya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BuhairyaCounty interface {
	feud.County
	BAbuminqar阿布民加尔() feud.Barony
	BBudhkula布赫库拉() feud.Barony
	BBuhairya拜哈里耶() feud.Barony
	BIbshawai伊布沙韦() feud.Barony
	BSineita西奈塔() feud.Barony
}

type c_buhairyaBuhairyaCounty struct {
	feud.BaseCounty
	阿布民加尔Abuminqar feud.Barony
	布赫库拉Budhkula   feud.Barony
	拜哈里耶Buhairya   feud.Barony
	伊布沙韦Ibshawai   feud.Barony
	西奈塔Sineita     feud.Barony
}

func (c *c_buhairyaBuhairyaCounty) BAbuminqar阿布民加尔() feud.Barony {
	return c.阿布民加尔Abuminqar
}

func (c *c_buhairyaBuhairyaCounty) BBudhkula布赫库拉() feud.Barony {
	return c.布赫库拉Budhkula
}

func (c *c_buhairyaBuhairyaCounty) BBuhairya拜哈里耶() feud.Barony {
	return c.拜哈里耶Buhairya
}

func (c *c_buhairyaBuhairyaCounty) BIbshawai伊布沙韦() feud.Barony {
	return c.伊布沙韦Ibshawai
}

func (c *c_buhairyaBuhairyaCounty) BSineita西奈塔() feud.Barony {
	return c.西奈塔Sineita
}

var CBuhairyac_buhairya BuhairyaCounty = &c_buhairyaBuhairyaCounty{}

func init() {
	f := CBuhairyac_buhairya.(*c_buhairyaBuhairyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "801",
		Title:     "buhairya",
		TitleName: "c_buhairya",
		TitleCode: "c_buhairya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布民加尔Abuminqar = BAbuminqar阿布民加尔
	f.阿布民加尔Abuminqar.SetParent(f)

	f.布赫库拉Budhkula = BBudhkula布赫库拉
	f.布赫库拉Budhkula.SetParent(f)

	f.拜哈里耶Buhairya = BBuhairya拜哈里耶
	f.拜哈里耶Buhairya.SetParent(f)

	f.伊布沙韦Ibshawai = BIbshawai伊布沙韦
	f.伊布沙韦Ibshawai.SetParent(f)

	f.西奈塔Sineita = BSineita西奈塔
	f.西奈塔Sineita.SetParent(f)

}
