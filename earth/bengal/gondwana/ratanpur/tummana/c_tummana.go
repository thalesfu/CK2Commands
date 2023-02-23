package tummana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TummanaCounty interface {
	feud.County
	BAmarkantak阿摩罗建吒迦() feud.Barony
	BManendragarh摩你因陀罗姞利呬() feud.Barony
	BShumhi首摩醯() feud.Barony
	BTorxem道先() feud.Barony
	BTualbung堵尔布恩() feud.Barony
	BTummana都摩那() feud.Barony
	BVij毗恃() feud.Barony
}

type 都摩那TummanaCounty struct {
	feud.BaseCounty
	阿摩罗建吒迦Amarkantak     feud.Barony
	摩你因陀罗姞利呬Manendragarh feud.Barony
	首摩醯Shumhi            feud.Barony
	道先Torxem             feud.Barony
	堵尔布恩Tualbung         feud.Barony
	都摩那Tummana           feud.Barony
	毗恃Vij                feud.Barony
}

func (c *都摩那TummanaCounty) BAmarkantak阿摩罗建吒迦() feud.Barony {
	return c.阿摩罗建吒迦Amarkantak
}

func (c *都摩那TummanaCounty) BManendragarh摩你因陀罗姞利呬() feud.Barony {
	return c.摩你因陀罗姞利呬Manendragarh
}

func (c *都摩那TummanaCounty) BShumhi首摩醯() feud.Barony {
	return c.首摩醯Shumhi
}

func (c *都摩那TummanaCounty) BTorxem道先() feud.Barony {
	return c.道先Torxem
}

func (c *都摩那TummanaCounty) BTualbung堵尔布恩() feud.Barony {
	return c.堵尔布恩Tualbung
}

func (c *都摩那TummanaCounty) BTummana都摩那() feud.Barony {
	return c.都摩那Tummana
}

func (c *都摩那TummanaCounty) BVij毗恃() feud.Barony {
	return c.毗恃Vij
}

var CTummana都摩那 TummanaCounty = &都摩那TummanaCounty{}

func init() {
	f := CTummana都摩那.(*都摩那TummanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1277",
		Title:     "tummana",
		TitleName: "都摩那",
		TitleCode: "c_tummana",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿摩罗建吒迦Amarkantak = BAmarkantak阿摩罗建吒迦
	f.阿摩罗建吒迦Amarkantak.SetParent(f)

	f.摩你因陀罗姞利呬Manendragarh = BManendragarh摩你因陀罗姞利呬
	f.摩你因陀罗姞利呬Manendragarh.SetParent(f)

	f.首摩醯Shumhi = BShumhi首摩醯
	f.首摩醯Shumhi.SetParent(f)

	f.道先Torxem = BTorxem道先
	f.道先Torxem.SetParent(f)

	f.堵尔布恩Tualbung = BTualbung堵尔布恩
	f.堵尔布恩Tualbung.SetParent(f)

	f.都摩那Tummana = BTummana都摩那
	f.都摩那Tummana.SetParent(f)

	f.毗恃Vij = BVij毗恃
	f.毗恃Vij.SetParent(f)

}
