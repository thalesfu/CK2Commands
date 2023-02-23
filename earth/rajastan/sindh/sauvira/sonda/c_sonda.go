package sonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SondaCounty interface {
	feud.County
	BAmarkot阿莫乔特() feud.Barony
	BBadin婆陈() feud.Barony
	BKaroonjar迦卢恩奢() feud.Barony
	BMamhal蒙诃尔() feud.Barony
	BMatli秣尼() feud.Barony
	BNagarparkar那揭罗波罗迦罗() feud.Barony
	BSonda孙陀() feud.Barony
}

type 孙陀SondaCounty struct {
	feud.BaseCounty
	阿莫乔特Amarkot        feud.Barony
	婆陈Badin            feud.Barony
	迦卢恩奢Karoonjar      feud.Barony
	蒙诃尔Mamhal          feud.Barony
	秣尼Matli            feud.Barony
	那揭罗波罗迦罗Nagarparkar feud.Barony
	孙陀Sonda            feud.Barony
}

func (c *孙陀SondaCounty) BAmarkot阿莫乔特() feud.Barony {
	return c.阿莫乔特Amarkot
}

func (c *孙陀SondaCounty) BBadin婆陈() feud.Barony {
	return c.婆陈Badin
}

func (c *孙陀SondaCounty) BKaroonjar迦卢恩奢() feud.Barony {
	return c.迦卢恩奢Karoonjar
}

func (c *孙陀SondaCounty) BMamhal蒙诃尔() feud.Barony {
	return c.蒙诃尔Mamhal
}

func (c *孙陀SondaCounty) BMatli秣尼() feud.Barony {
	return c.秣尼Matli
}

func (c *孙陀SondaCounty) BNagarparkar那揭罗波罗迦罗() feud.Barony {
	return c.那揭罗波罗迦罗Nagarparkar
}

func (c *孙陀SondaCounty) BSonda孙陀() feud.Barony {
	return c.孙陀Sonda
}

var CSonda孙陀 SondaCounty = &孙陀SondaCounty{}

func init() {
	f := CSonda孙陀.(*孙陀SondaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1331",
		Title:     "sonda",
		TitleName: "孙陀",
		TitleCode: "c_sonda",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿莫乔特Amarkot = BAmarkot阿莫乔特
	f.阿莫乔特Amarkot.SetParent(f)

	f.婆陈Badin = BBadin婆陈
	f.婆陈Badin.SetParent(f)

	f.迦卢恩奢Karoonjar = BKaroonjar迦卢恩奢
	f.迦卢恩奢Karoonjar.SetParent(f)

	f.蒙诃尔Mamhal = BMamhal蒙诃尔
	f.蒙诃尔Mamhal.SetParent(f)

	f.秣尼Matli = BMatli秣尼
	f.秣尼Matli.SetParent(f)

	f.那揭罗波罗迦罗Nagarparkar = BNagarparkar那揭罗波罗迦罗
	f.那揭罗波罗迦罗Nagarparkar.SetParent(f)

	f.孙陀Sonda = BSonda孙陀
	f.孙陀Sonda.SetParent(f)

}
