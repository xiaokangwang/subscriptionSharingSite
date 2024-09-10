package base64urlline

import (
	"bytes"
	"encoding/base64"
	"io"
	"strings"

	"github.com/xiaokangwang/subscriptionSharingSite/subscription/containers"
)

func NewBase64URLLineWrapper() containers.SubscriptionContainerDocumentWrapper {
	return &wrapper{}
}

type wrapper struct {
}

func (w *wrapper) WrapSubscriptionContainerDocument(config *containers.Container) ([]byte, error) {
	var sb strings.Builder
	for _, v := range config.ServerSpecs {
		trimedContent :=
			strings.ReplaceAll(strings.Trim(string(v.Content), " \n\t"), "\n", " ")
		sb.WriteString(trimedContent)
		sb.WriteString("\n")
	}
	var base64edContent bytes.Buffer
	encoder := base64.NewEncoder(base64.URLEncoding, &base64edContent)
	_, err := io.Copy(encoder, strings.NewReader(sb.String()))
	if err != nil {
		return nil, err
	}
	err = encoder.Close()
	if err != nil {
		return nil, err
	}
	return base64edContent.Bytes(), nil
}
