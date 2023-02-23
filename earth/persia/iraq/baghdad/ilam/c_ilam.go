package ilam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IlamCounty interface {
	feud.County
	BAbdanan阿卜达南() feud.Barony
	BChahartaghi察哈尔塔赫() feud.Barony
	BDehloran代赫洛兰() feud.Barony
	BGhalehghiran吉阿尔赫兰() feud.Barony
	BHezardar赫泽拉达尔() feud.Barony
	BIlam伊拉姆() feud.Barony
	BMehran梅赫兰() feud.Barony
	BTowhid塔胡德() feud.Barony
}

type 伊拉姆IlamCounty struct {
	feud.BaseCounty
	阿卜达南Abdanan       feud.Barony
	察哈尔塔赫Chahartaghi  feud.Barony
	代赫洛兰Dehloran      feud.Barony
	吉阿尔赫兰Ghalehghiran feud.Barony
	赫泽拉达尔Hezardar     feud.Barony
	伊拉姆Ilam           feud.Barony
	梅赫兰Mehran         feud.Barony
	塔胡德Towhid         feud.Barony
}

func (c *伊拉姆IlamCounty) BAbdanan阿卜达南() feud.Barony {
	return c.阿卜达南Abdanan
}

func (c *伊拉姆IlamCounty) BChahartaghi察哈尔塔赫() feud.Barony {
	return c.察哈尔塔赫Chahartaghi
}

func (c *伊拉姆IlamCounty) BDehloran代赫洛兰() feud.Barony {
	return c.代赫洛兰Dehloran
}

func (c *伊拉姆IlamCounty) BGhalehghiran吉阿尔赫兰() feud.Barony {
	return c.吉阿尔赫兰Ghalehghiran
}

func (c *伊拉姆IlamCounty) BHezardar赫泽拉达尔() feud.Barony {
	return c.赫泽拉达尔Hezardar
}

func (c *伊拉姆IlamCounty) BIlam伊拉姆() feud.Barony {
	return c.伊拉姆Ilam
}

func (c *伊拉姆IlamCounty) BMehran梅赫兰() feud.Barony {
	return c.梅赫兰Mehran
}

func (c *伊拉姆IlamCounty) BTowhid塔胡德() feud.Barony {
	return c.塔胡德Towhid
}

var CIlam伊拉姆 IlamCounty = &伊拉姆IlamCounty{}

func init() {
	f := CIlam伊拉姆.(*伊拉姆IlamCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "689",
		Title:     "ilam",
		TitleName: "伊拉姆",
		TitleCode: "c_ilam",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卜达南Abdanan = BAbdanan阿卜达南
	f.阿卜达南Abdanan.SetParent(f)

	f.察哈尔塔赫Chahartaghi = BChahartaghi察哈尔塔赫
	f.察哈尔塔赫Chahartaghi.SetParent(f)

	f.代赫洛兰Dehloran = BDehloran代赫洛兰
	f.代赫洛兰Dehloran.SetParent(f)

	f.吉阿尔赫兰Ghalehghiran = BGhalehghiran吉阿尔赫兰
	f.吉阿尔赫兰Ghalehghiran.SetParent(f)

	f.赫泽拉达尔Hezardar = BHezardar赫泽拉达尔
	f.赫泽拉达尔Hezardar.SetParent(f)

	f.伊拉姆Ilam = BIlam伊拉姆
	f.伊拉姆Ilam.SetParent(f)

	f.梅赫兰Mehran = BMehran梅赫兰
	f.梅赫兰Mehran.SetParent(f)

	f.塔胡德Towhid = BTowhid塔胡德
	f.塔胡德Towhid.SetParent(f)

}
