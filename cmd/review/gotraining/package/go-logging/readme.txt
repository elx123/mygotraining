首先最主要的结构有2个
type Logger struct {
	Module      string
	backend     LeveledBackend
	haveBackend bool

	// ExtraCallDepth can be used to add additional call depth when getting the
	// calling function. This is normally used when wrapping a logger.
	ExtraCalldepth int
}

和LevelBackend，其中LevelBackend有默认值，也就是包中有默认值
