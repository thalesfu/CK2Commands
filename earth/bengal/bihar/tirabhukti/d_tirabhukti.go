package tirabhukti

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/tirabhukti/kusinagara"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/tirabhukti/mithila"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/tirabhukti/simaramapura"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TirabhuktiDuke interface {
	feud.Duke
	CKusinagara拘尸那揭罗() kusinagara.KusinagaraCounty
	CMithila弥梯罗() mithila.MithilaCounty
	CSimaramapura僧摩罗摩城() simaramapura.SimaramapuraCounty
}

type 帝那伏帝TirabhuktiDuke struct {
	feud.BaseDuke
	拘尸那揭罗Kusinagara   kusinagara.KusinagaraCounty
	弥梯罗Mithila        mithila.MithilaCounty
	僧摩罗摩城Simaramapura simaramapura.SimaramapuraCounty
}

func (d *帝那伏帝TirabhuktiDuke) CKusinagara拘尸那揭罗() kusinagara.KusinagaraCounty {
	return d.拘尸那揭罗Kusinagara
}

func (d *帝那伏帝TirabhuktiDuke) CMithila弥梯罗() mithila.MithilaCounty {
	return d.弥梯罗Mithila
}

func (d *帝那伏帝TirabhuktiDuke) CSimaramapura僧摩罗摩城() simaramapura.SimaramapuraCounty {
	return d.僧摩罗摩城Simaramapura
}

var DTirabhukti帝那伏帝 TirabhuktiDuke = &帝那伏帝TirabhuktiDuke{}

func init() {
	f := DTirabhukti帝那伏帝.(*帝那伏帝TirabhuktiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tirabhukti",
		TitleName: "帝那伏帝",
		TitleCode: "d_tirabhukti",
		Counties:  map[string]feud.County{},
	}

	f.拘尸那揭罗Kusinagara = kusinagara.CKusinagara拘尸那揭罗
	f.拘尸那揭罗Kusinagara.SetParent(f)

	f.弥梯罗Mithila = mithila.CMithila弥梯罗
	f.弥梯罗Mithila.SetParent(f)

	f.僧摩罗摩城Simaramapura = simaramapura.CSimaramapura僧摩罗摩城
	f.僧摩罗摩城Simaramapura.SetParent(f)

}
