package thuringia

import (
	"github.com/thalesfu/CK2Commands/earth/germany/germany/thuringia/thuringen"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/thuringia/weimar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThuringiaDuke interface {
	feud.Duke
	CThuringen图林根() thuringen.ThuringenCounty
	CWeimar魏玛() weimar.WeimarCounty
}

type 图林根ThuringiaDuke struct {
	feud.BaseDuke
	图林根Thuringen thuringen.ThuringenCounty
	魏玛Weimar     weimar.WeimarCounty
}

func (d *图林根ThuringiaDuke) CThuringen图林根() thuringen.ThuringenCounty {
	return d.图林根Thuringen
}

func (d *图林根ThuringiaDuke) CWeimar魏玛() weimar.WeimarCounty {
	return d.魏玛Weimar
}

var DThuringia图林根 ThuringiaDuke = &图林根ThuringiaDuke{}

func init() {
	f := DThuringia图林根.(*图林根ThuringiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "thuringia",
		TitleName: "图林根",
		TitleCode: "d_thuringia",
		Counties:  map[string]feud.County{},
	}

	f.图林根Thuringen = thuringen.CThuringen图林根
	f.图林根Thuringen.SetParent(f)

	f.魏玛Weimar = weimar.CWeimar魏玛
	f.魏玛Weimar.SetParent(f)

}
