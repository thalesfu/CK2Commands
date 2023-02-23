package ishim

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ishim/ishim"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ishim/kipchak"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ishim/seletyteniz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IshimDuke interface {
	feud.Duke
	CIshim伊希姆() ishim.IshimCounty
	CKipchak田吉兹() kipchak.KipchakCounty
	CSeletyteniz谢列特捷尼兹() seletyteniz.SeletytenizCounty
}

type 伊希姆IshimDuke struct {
	feud.BaseDuke
	伊希姆Ishim          ishim.IshimCounty
	田吉兹Kipchak        kipchak.KipchakCounty
	谢列特捷尼兹Seletyteniz seletyteniz.SeletytenizCounty
}

func (d *伊希姆IshimDuke) CIshim伊希姆() ishim.IshimCounty {
	return d.伊希姆Ishim
}

func (d *伊希姆IshimDuke) CKipchak田吉兹() kipchak.KipchakCounty {
	return d.田吉兹Kipchak
}

func (d *伊希姆IshimDuke) CSeletyteniz谢列特捷尼兹() seletyteniz.SeletytenizCounty {
	return d.谢列特捷尼兹Seletyteniz
}

var DIshim伊希姆 IshimDuke = &伊希姆IshimDuke{}

func init() {
	f := DIshim伊希姆.(*伊希姆IshimDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ishim",
		TitleName: "伊希姆",
		TitleCode: "d_ishim",
		Counties:  map[string]feud.County{},
	}

	f.伊希姆Ishim = ishim.CIshim伊希姆
	f.伊希姆Ishim.SetParent(f)

	f.田吉兹Kipchak = kipchak.CKipchak田吉兹
	f.田吉兹Kipchak.SetParent(f)

	f.谢列特捷尼兹Seletyteniz = seletyteniz.CSeletyteniz谢列特捷尼兹
	f.谢列特捷尼兹Seletyteniz.SetParent(f)

}
