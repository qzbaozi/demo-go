package main

type container struct {
	items [3]item
}

func (c *container) Len() int {
	return 10
}

type item struct {
}

func (c *container) Iter() <-chan item {
	ch := make(chan item)
	go func() {
		for i := 0; i < c.Len(); i++ { // or use a for-range loop
			ch <- c.items[i]
		}
	}()
	return ch
}
