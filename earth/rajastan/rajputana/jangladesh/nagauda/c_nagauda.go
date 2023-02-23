package nagauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NagaudaCounty interface {
	feud.County
	BAhichhatrapur恶醯掣呾逻补罗() feud.Barony
	BKuchaman古贾曼() feud.Barony
	BLadnu罗努() feud.Barony
	BLadnun罗努恩() feud.Barony
	BMakrana摩伽罗那() feud.Barony
	BMundwa蒙德拉() feud.Barony
	BNagauda那乔佗() feud.Barony
}

type 那乔佗NagaudaCounty struct {
	feud.BaseCounty
	恶醯掣呾逻补罗Ahichhatrapur feud.Barony
	古贾曼Kuchaman          feud.Barony
	罗努Ladnu              feud.Barony
	罗努恩Ladnun            feud.Barony
	摩伽罗那Makrana          feud.Barony
	蒙德拉Mundwa            feud.Barony
	那乔佗Nagauda           feud.Barony
}

func (c *那乔佗NagaudaCounty) BAhichhatrapur恶醯掣呾逻补罗() feud.Barony {
	return c.恶醯掣呾逻补罗Ahichhatrapur
}

func (c *那乔佗NagaudaCounty) BKuchaman古贾曼() feud.Barony {
	return c.古贾曼Kuchaman
}

func (c *那乔佗NagaudaCounty) BLadnu罗努() feud.Barony {
	return c.罗努Ladnu
}

func (c *那乔佗NagaudaCounty) BLadnun罗努恩() feud.Barony {
	return c.罗努恩Ladnun
}

func (c *那乔佗NagaudaCounty) BMakrana摩伽罗那() feud.Barony {
	return c.摩伽罗那Makrana
}

func (c *那乔佗NagaudaCounty) BMundwa蒙德拉() feud.Barony {
	return c.蒙德拉Mundwa
}

func (c *那乔佗NagaudaCounty) BNagauda那乔佗() feud.Barony {
	return c.那乔佗Nagauda
}

var CNagauda那乔佗 NagaudaCounty = &那乔佗NagaudaCounty{}

func init() {
	f := CNagauda那乔佗.(*那乔佗NagaudaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1354",
		Title:     "nagauda",
		TitleName: "那乔佗",
		TitleCode: "c_nagauda",
		Baronies:  map[string]feud.Barony{},
	}

	f.恶醯掣呾逻补罗Ahichhatrapur = BAhichhatrapur恶醯掣呾逻补罗
	f.恶醯掣呾逻补罗Ahichhatrapur.SetParent(f)

	f.古贾曼Kuchaman = BKuchaman古贾曼
	f.古贾曼Kuchaman.SetParent(f)

	f.罗努Ladnu = BLadnu罗努
	f.罗努Ladnu.SetParent(f)

	f.罗努恩Ladnun = BLadnun罗努恩
	f.罗努恩Ladnun.SetParent(f)

	f.摩伽罗那Makrana = BMakrana摩伽罗那
	f.摩伽罗那Makrana.SetParent(f)

	f.蒙德拉Mundwa = BMundwa蒙德拉
	f.蒙德拉Mundwa.SetParent(f)

	f.那乔佗Nagauda = BNagauda那乔佗
	f.那乔佗Nagauda.SetParent(f)

}
