package khojand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhojandCounty interface {
	feud.County
	BAsht阿什特() feud.Barony
	BIlaq义剌克() feud.Barony
	BKendibadam肯德伊巴达姆() feud.Barony
	BKhavakend浩罕() feud.Barony
	BKhojand苦盏() feud.Barony
	BSokh索赫() feud.Barony
	BTunket腾凯特() feud.Barony
}

type 苦盏KhojandCounty struct {
	feud.BaseCounty
	阿什特Asht          feud.Barony
	义剌克Ilaq          feud.Barony
	肯德伊巴达姆Kendibadam feud.Barony
	浩罕Khavakend      feud.Barony
	苦盏Khojand        feud.Barony
	索赫Sokh           feud.Barony
	腾凯特Tunket        feud.Barony
}

func (c *苦盏KhojandCounty) BAsht阿什特() feud.Barony {
	return c.阿什特Asht
}

func (c *苦盏KhojandCounty) BIlaq义剌克() feud.Barony {
	return c.义剌克Ilaq
}

func (c *苦盏KhojandCounty) BKendibadam肯德伊巴达姆() feud.Barony {
	return c.肯德伊巴达姆Kendibadam
}

func (c *苦盏KhojandCounty) BKhavakend浩罕() feud.Barony {
	return c.浩罕Khavakend
}

func (c *苦盏KhojandCounty) BKhojand苦盏() feud.Barony {
	return c.苦盏Khojand
}

func (c *苦盏KhojandCounty) BSokh索赫() feud.Barony {
	return c.索赫Sokh
}

func (c *苦盏KhojandCounty) BTunket腾凯特() feud.Barony {
	return c.腾凯特Tunket
}

var CKhojand苦盏 KhojandCounty = &苦盏KhojandCounty{}

func init() {
	f := CKhojand苦盏.(*苦盏KhojandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1546",
		Title:     "khojand",
		TitleName: "苦盏",
		TitleCode: "c_khojand",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿什特Asht = BAsht阿什特
	f.阿什特Asht.SetParent(f)

	f.义剌克Ilaq = BIlaq义剌克
	f.义剌克Ilaq.SetParent(f)

	f.肯德伊巴达姆Kendibadam = BKendibadam肯德伊巴达姆
	f.肯德伊巴达姆Kendibadam.SetParent(f)

	f.浩罕Khavakend = BKhavakend浩罕
	f.浩罕Khavakend.SetParent(f)

	f.苦盏Khojand = BKhojand苦盏
	f.苦盏Khojand.SetParent(f)

	f.索赫Sokh = BSokh索赫
	f.索赫Sokh.SetParent(f)

	f.腾凯特Tunket = BTunket腾凯特
	f.腾凯特Tunket.SetParent(f)

}
