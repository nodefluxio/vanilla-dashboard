package event

import (
	"github.com/gorilla/websocket"
	"gitlab.com/nodefluxio/vanilla-dashboard/internal/pkg/util"
	"gitlab.com/nodefluxio/vanilla-dashboard/internal/presenter"
	"golang.org/x/net/context"
)

// Service interface abstracts the controller layer and should be implemented in controller directory. Controller contains business logics and is independent of any database connection.
type Service interface {
	// AddAdditionalData(data map[string]interface{}) (map[string]interface{}, error)
	InitiateDataStream(StreamID string, conn *websocket.Conn)
	CronjobPartition(ctx context.Context) error
	Partition(ctx context.Context) error
	Dumping(ctx context.Context)
	GetHistory(ctx context.Context, paging *util.Pagination) (*presenter.EventHistoryPaging, error)
	ExportEvent(ctx context.Context, paging *util.Pagination) error
	CheckExportedEvent(ctx context.Context) (string, error)
	UpdateStatusExportDownload(ctx context.Context) error
}
