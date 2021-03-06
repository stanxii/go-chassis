package resource

import (
	"github.com/ServiceComb/go-chassis/core/handler"
	"github.com/ServiceComb/go-chassis/core/invocation"
)

//MetadataHandler
type MetadataHandler struct {
}

//Handle
func (h *MetadataHandler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	inv.SetMetadata("auth", "user1")
	chain.Next(inv, cb)
}

//Name
func (h *MetadataHandler) Name() string {
	return "test"
}
func newMetadataHandler() handler.Handler {
	//call next chain

	return &MetadataHandler{}
}

func init() {
	handler.RegisterHandler("metadata-handler", newMetadataHandler)
}
