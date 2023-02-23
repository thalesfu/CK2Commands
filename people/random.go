package people

func BuildRandom(peoples ...int) {
	var functions = [][]buildFunc{
		getAmbassadorFunctions(),
		getGeneralFunctions(),
		getManagerFunctions(),
		getSpyFunctions(),
		getFatherFunctions(),
		getGeneralFunctions(),
	}

	buildPeopleRandom("random", functions, peoples...)
}
