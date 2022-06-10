package ut

type translatorV2 struct {
}

type transTextV2 struct {
	text    string
	indexes []idxParam
}

type idxParam struct {
	idxStart int
	idxEnd   int
	val      int
}
