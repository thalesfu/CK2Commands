package vidin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VidinCounty interface {
	feud.County
	BBolvan博尔万() feud.Barony
	BKucevo库切沃() feud.Barony
	BKula库拉() feud.Barony
	BPirot皮罗特() feud.Barony
	BSrvljig斯夫尔利格() feud.Barony
	BVidin维丁() feud.Barony
	BViseslav维舍斯拉夫() feud.Barony
	BZajecar扎耶查尔() feud.Barony
}

type 维丁VidinCounty struct {
	feud.BaseCounty
	博尔万Bolvan     feud.Barony
	库切沃Kucevo     feud.Barony
	库拉Kula        feud.Barony
	皮罗特Pirot      feud.Barony
	斯夫尔利格Srvljig  feud.Barony
	维丁Vidin       feud.Barony
	维舍斯拉夫Viseslav feud.Barony
	扎耶查尔Zajecar   feud.Barony
}

func (c *维丁VidinCounty) BBolvan博尔万() feud.Barony {
	return c.博尔万Bolvan
}

func (c *维丁VidinCounty) BKucevo库切沃() feud.Barony {
	return c.库切沃Kucevo
}

func (c *维丁VidinCounty) BKula库拉() feud.Barony {
	return c.库拉Kula
}

func (c *维丁VidinCounty) BPirot皮罗特() feud.Barony {
	return c.皮罗特Pirot
}

func (c *维丁VidinCounty) BSrvljig斯夫尔利格() feud.Barony {
	return c.斯夫尔利格Srvljig
}

func (c *维丁VidinCounty) BVidin维丁() feud.Barony {
	return c.维丁Vidin
}

func (c *维丁VidinCounty) BViseslav维舍斯拉夫() feud.Barony {
	return c.维舍斯拉夫Viseslav
}

func (c *维丁VidinCounty) BZajecar扎耶查尔() feud.Barony {
	return c.扎耶查尔Zajecar
}

var CVidin维丁 VidinCounty = &维丁VidinCounty{}

func init() {
	f := CVidin维丁.(*维丁VidinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "506",
		Title:     "vidin",
		TitleName: "维丁",
		TitleCode: "c_vidin",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔万Bolvan = BBolvan博尔万
	f.博尔万Bolvan.SetParent(f)

	f.库切沃Kucevo = BKucevo库切沃
	f.库切沃Kucevo.SetParent(f)

	f.库拉Kula = BKula库拉
	f.库拉Kula.SetParent(f)

	f.皮罗特Pirot = BPirot皮罗特
	f.皮罗特Pirot.SetParent(f)

	f.斯夫尔利格Srvljig = BSrvljig斯夫尔利格
	f.斯夫尔利格Srvljig.SetParent(f)

	f.维丁Vidin = BVidin维丁
	f.维丁Vidin.SetParent(f)

	f.维舍斯拉夫Viseslav = BViseslav维舍斯拉夫
	f.维舍斯拉夫Viseslav.SetParent(f)

	f.扎耶查尔Zajecar = BZajecar扎耶查尔
	f.扎耶查尔Zajecar.SetParent(f)

}
