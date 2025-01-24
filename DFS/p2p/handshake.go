package p2p

type handShakeFunc func(any) error

func NOPHandShakeFunc(any) error { return nil }
