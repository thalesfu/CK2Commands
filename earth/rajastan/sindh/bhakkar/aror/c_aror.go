package aror

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArorCounty interface {
	feud.County
	BAror阿卢梨() feud.Barony
	BBasmak巴斯马克() feud.Barony
	BChak查克() feud.Barony
	BMahokhar摩呼伽() feud.Barony
	BSukkur苏库尔() feud.Barony
}

type 阿卢梨ArorCounty struct {
	feud.BaseCounty
	阿卢梨Aror     feud.Barony
	巴斯马克Basmak  feud.Barony
	查克Chak      feud.Barony
	摩呼伽Mahokhar feud.Barony
	苏库尔Sukkur   feud.Barony
}

func (c *阿卢梨ArorCounty) BAror阿卢梨() feud.Barony {
	return c.阿卢梨Aror
}

func (c *阿卢梨ArorCounty) BBasmak巴斯马克() feud.Barony {
	return c.巴斯马克Basmak
}

func (c *阿卢梨ArorCounty) BChak查克() feud.Barony {
	return c.查克Chak
}

func (c *阿卢梨ArorCounty) BMahokhar摩呼伽() feud.Barony {
	return c.摩呼伽Mahokhar
}

func (c *阿卢梨ArorCounty) BSukkur苏库尔() feud.Barony {
	return c.苏库尔Sukkur
}

var CAror阿卢梨 ArorCounty = &阿卢梨ArorCounty{}

func init() {
	f := CAror阿卢梨.(*阿卢梨ArorCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1175",
		Title:     "aror",
		TitleName: "阿卢梨",
		TitleCode: "c_aror",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卢梨Aror = BAror阿卢梨
	f.阿卢梨Aror.SetParent(f)

	f.巴斯马克Basmak = BBasmak巴斯马克
	f.巴斯马克Basmak.SetParent(f)

	f.查克Chak = BChak查克
	f.查克Chak.SetParent(f)

	f.摩呼伽Mahokhar = BMahokhar摩呼伽
	f.摩呼伽Mahokhar.SetParent(f)

	f.苏库尔Sukkur = BSukkur苏库尔
	f.苏库尔Sukkur.SetParent(f)

}
