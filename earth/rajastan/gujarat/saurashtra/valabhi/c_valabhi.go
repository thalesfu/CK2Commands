package valabhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ValabhiCounty interface {
	feud.County
	BDhandhuka弹度迦() feud.Barony
	BGundigar古蒂迦() feud.Barony
	BMadya摩地耶() feud.Barony
	BPalatina波利旦那() feud.Barony
	BShatrunjaya设咄路阇耶() feud.Barony
	BSihor锡霍尔() feud.Barony
	BValabhi伐腊毗() feud.Barony
}

type 伐腊毗ValabhiCounty struct {
	feud.BaseCounty
	弹度迦Dhandhuka     feud.Barony
	古蒂迦Gundigar      feud.Barony
	摩地耶Madya         feud.Barony
	波利旦那Palatina     feud.Barony
	设咄路阇耶Shatrunjaya feud.Barony
	锡霍尔Sihor         feud.Barony
	伐腊毗Valabhi       feud.Barony
}

func (c *伐腊毗ValabhiCounty) BDhandhuka弹度迦() feud.Barony {
	return c.弹度迦Dhandhuka
}

func (c *伐腊毗ValabhiCounty) BGundigar古蒂迦() feud.Barony {
	return c.古蒂迦Gundigar
}

func (c *伐腊毗ValabhiCounty) BMadya摩地耶() feud.Barony {
	return c.摩地耶Madya
}

func (c *伐腊毗ValabhiCounty) BPalatina波利旦那() feud.Barony {
	return c.波利旦那Palatina
}

func (c *伐腊毗ValabhiCounty) BShatrunjaya设咄路阇耶() feud.Barony {
	return c.设咄路阇耶Shatrunjaya
}

func (c *伐腊毗ValabhiCounty) BSihor锡霍尔() feud.Barony {
	return c.锡霍尔Sihor
}

func (c *伐腊毗ValabhiCounty) BValabhi伐腊毗() feud.Barony {
	return c.伐腊毗Valabhi
}

var CValabhi伐腊毗 ValabhiCounty = &伐腊毗ValabhiCounty{}

func init() {
	f := CValabhi伐腊毗.(*伐腊毗ValabhiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1135",
		Title:     "valabhi",
		TitleName: "伐腊毗",
		TitleCode: "c_valabhi",
		Baronies:  map[string]feud.Barony{},
	}

	f.弹度迦Dhandhuka = BDhandhuka弹度迦
	f.弹度迦Dhandhuka.SetParent(f)

	f.古蒂迦Gundigar = BGundigar古蒂迦
	f.古蒂迦Gundigar.SetParent(f)

	f.摩地耶Madya = BMadya摩地耶
	f.摩地耶Madya.SetParent(f)

	f.波利旦那Palatina = BPalatina波利旦那
	f.波利旦那Palatina.SetParent(f)

	f.设咄路阇耶Shatrunjaya = BShatrunjaya设咄路阇耶
	f.设咄路阇耶Shatrunjaya.SetParent(f)

	f.锡霍尔Sihor = BSihor锡霍尔
	f.锡霍尔Sihor.SetParent(f)

	f.伐腊毗Valabhi = BValabhi伐腊毗
	f.伐腊毗Valabhi.SetParent(f)

}
