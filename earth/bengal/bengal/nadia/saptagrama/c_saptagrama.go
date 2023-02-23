package saptagrama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaptagramaCounty interface {
	feud.County
	BBansberia般娑吠利阿() feud.Barony
	BChinsura真苏罗() feud.Barony
	BMahanad摩诃那陀() feud.Barony
	BPandua般菟阿() feud.Barony
	BSaptagrama七村() feud.Barony
	BShantipur扇底补罗() feud.Barony
	BTribeni帝利吠尼() feud.Barony
}

type 七村SaptagramaCounty struct {
	feud.BaseCounty
	般娑吠利阿Bansberia feud.Barony
	真苏罗Chinsura    feud.Barony
	摩诃那陀Mahanad    feud.Barony
	般菟阿Pandua      feud.Barony
	七村Saptagrama   feud.Barony
	扇底补罗Shantipur  feud.Barony
	帝利吠尼Tribeni    feud.Barony
}

func (c *七村SaptagramaCounty) BBansberia般娑吠利阿() feud.Barony {
	return c.般娑吠利阿Bansberia
}

func (c *七村SaptagramaCounty) BChinsura真苏罗() feud.Barony {
	return c.真苏罗Chinsura
}

func (c *七村SaptagramaCounty) BMahanad摩诃那陀() feud.Barony {
	return c.摩诃那陀Mahanad
}

func (c *七村SaptagramaCounty) BPandua般菟阿() feud.Barony {
	return c.般菟阿Pandua
}

func (c *七村SaptagramaCounty) BSaptagrama七村() feud.Barony {
	return c.七村Saptagrama
}

func (c *七村SaptagramaCounty) BShantipur扇底补罗() feud.Barony {
	return c.扇底补罗Shantipur
}

func (c *七村SaptagramaCounty) BTribeni帝利吠尼() feud.Barony {
	return c.帝利吠尼Tribeni
}

var CSaptagrama七村 SaptagramaCounty = &七村SaptagramaCounty{}

func init() {
	f := CSaptagrama七村.(*七村SaptagramaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1233",
		Title:     "saptagrama",
		TitleName: "七村",
		TitleCode: "c_saptagrama",
		Baronies:  map[string]feud.Barony{},
	}

	f.般娑吠利阿Bansberia = BBansberia般娑吠利阿
	f.般娑吠利阿Bansberia.SetParent(f)

	f.真苏罗Chinsura = BChinsura真苏罗
	f.真苏罗Chinsura.SetParent(f)

	f.摩诃那陀Mahanad = BMahanad摩诃那陀
	f.摩诃那陀Mahanad.SetParent(f)

	f.般菟阿Pandua = BPandua般菟阿
	f.般菟阿Pandua.SetParent(f)

	f.七村Saptagrama = BSaptagrama七村
	f.七村Saptagrama.SetParent(f)

	f.扇底补罗Shantipur = BShantipur扇底补罗
	f.扇底补罗Shantipur.SetParent(f)

	f.帝利吠尼Tribeni = BTribeni帝利吠尼
	f.帝利吠尼Tribeni.SetParent(f)

}
