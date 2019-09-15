package interfaceMock

type ReceiverMock struct {
	Content string
}

func (r ReceiverMock) Get(url string)  string{
	return r.Content
}