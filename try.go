package main

type tryCt struct {
	Err interface{}
}

func Try(f func()) *tryCt {
	var catch = &tryCt{}
	func() {
		defer func() {
			catch.Err = recover()
		}()
		f()
	}()
	return catch
}

func (c *tryCt) Catch(hand func(err interface{})) *tryCt {
	if c.Err != nil {
		hand(c.Err)
	}
	return c
}

func (c *tryCt) Final(final func()) {
	final()
}
