package containers

type UnparsedServerConf struct {
	KindHint string
	Content  []byte
}

type Container struct {
	Kind        string
	Metadata    map[string]string
	ServerSpecs []UnparsedServerConf
}

type SubscriptionContainerDocumentParser interface {
	ParseSubscriptionContainerDocument(rawConfig []byte) (*Container, error)
}

type SubscriptionContainerDocumentWrapper interface {
	WrapSubscriptionContainerDocument(Config *Container) ([]byte, error)
}
