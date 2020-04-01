package try

type tryCt struct {
	Err interface{}
}

func Try(f func()) tryCt {
	var chatch tryCt
	func() {
		defer func() {
			chatch.Err = recover()
		}()
		f()
	}()
	return chatch
}

func (c tryCt) Catch(hand func(interface{})) tryCt {
	if c.Err != nil {
		hand(c.Err)
	}
	return c
}

func (c tryCt) Final(final func()) {
	final()
}
