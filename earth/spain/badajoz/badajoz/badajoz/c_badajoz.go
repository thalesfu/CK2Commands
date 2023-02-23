package badajoz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BadajozCounty interface {
	feud.County
	BAlmendralejo阿尔门德拉莱霍() feud.Barony
	BBadajoz巴达霍斯() feud.Barony
	BFuentedelmaestre丰特德尔马埃斯特雷() feud.Barony
	BGuarena瓜雷尼亚() feud.Barony
	BJerezdeloscaballeros赫雷斯德洛斯卡瓦列罗斯() feud.Barony
	BMerida梅里达() feud.Barony
	BVillalbadelosbarros比利亚尔瓦德洛斯巴罗斯() feud.Barony
	BZafra萨夫拉() feud.Barony
}

type 巴达霍斯BadajozCounty struct {
	feud.BaseCounty
	阿尔门德拉莱霍Almendralejo             feud.Barony
	巴达霍斯Badajoz                     feud.Barony
	丰特德尔马埃斯特雷Fuentedelmaestre       feud.Barony
	瓜雷尼亚Guarena                     feud.Barony
	赫雷斯德洛斯卡瓦列罗斯Jerezdeloscaballeros feud.Barony
	梅里达Merida                       feud.Barony
	比利亚尔瓦德洛斯巴罗斯Villalbadelosbarros  feud.Barony
	萨夫拉Zafra                        feud.Barony
}

func (c *巴达霍斯BadajozCounty) BAlmendralejo阿尔门德拉莱霍() feud.Barony {
	return c.阿尔门德拉莱霍Almendralejo
}

func (c *巴达霍斯BadajozCounty) BBadajoz巴达霍斯() feud.Barony {
	return c.巴达霍斯Badajoz
}

func (c *巴达霍斯BadajozCounty) BFuentedelmaestre丰特德尔马埃斯特雷() feud.Barony {
	return c.丰特德尔马埃斯特雷Fuentedelmaestre
}

func (c *巴达霍斯BadajozCounty) BGuarena瓜雷尼亚() feud.Barony {
	return c.瓜雷尼亚Guarena
}

func (c *巴达霍斯BadajozCounty) BJerezdeloscaballeros赫雷斯德洛斯卡瓦列罗斯() feud.Barony {
	return c.赫雷斯德洛斯卡瓦列罗斯Jerezdeloscaballeros
}

func (c *巴达霍斯BadajozCounty) BMerida梅里达() feud.Barony {
	return c.梅里达Merida
}

func (c *巴达霍斯BadajozCounty) BVillalbadelosbarros比利亚尔瓦德洛斯巴罗斯() feud.Barony {
	return c.比利亚尔瓦德洛斯巴罗斯Villalbadelosbarros
}

func (c *巴达霍斯BadajozCounty) BZafra萨夫拉() feud.Barony {
	return c.萨夫拉Zafra
}

var CBadajoz巴达霍斯 BadajozCounty = &巴达霍斯BadajozCounty{}

func init() {
	f := CBadajoz巴达霍斯.(*巴达霍斯BadajozCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "184",
		Title:     "badajoz",
		TitleName: "巴达霍斯",
		TitleCode: "c_badajoz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔门德拉莱霍Almendralejo = BAlmendralejo阿尔门德拉莱霍
	f.阿尔门德拉莱霍Almendralejo.SetParent(f)

	f.巴达霍斯Badajoz = BBadajoz巴达霍斯
	f.巴达霍斯Badajoz.SetParent(f)

	f.丰特德尔马埃斯特雷Fuentedelmaestre = BFuentedelmaestre丰特德尔马埃斯特雷
	f.丰特德尔马埃斯特雷Fuentedelmaestre.SetParent(f)

	f.瓜雷尼亚Guarena = BGuarena瓜雷尼亚
	f.瓜雷尼亚Guarena.SetParent(f)

	f.赫雷斯德洛斯卡瓦列罗斯Jerezdeloscaballeros = BJerezdeloscaballeros赫雷斯德洛斯卡瓦列罗斯
	f.赫雷斯德洛斯卡瓦列罗斯Jerezdeloscaballeros.SetParent(f)

	f.梅里达Merida = BMerida梅里达
	f.梅里达Merida.SetParent(f)

	f.比利亚尔瓦德洛斯巴罗斯Villalbadelosbarros = BVillalbadelosbarros比利亚尔瓦德洛斯巴罗斯
	f.比利亚尔瓦德洛斯巴罗斯Villalbadelosbarros.SetParent(f)

	f.萨夫拉Zafra = BZafra萨夫拉
	f.萨夫拉Zafra.SetParent(f)

}
