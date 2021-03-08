package services

import (
	"time"

	"github.com/mercadolibre/pruebas/configs"

	"github.com/mercadolibre/fury_goauditclient/src/api/audittrail"
)

var auditClient audittrail.Client

func init() {
	auditConfig := audittrail.MakeAuditConfig()
	auditConfig.SetTimeout(configs.AuditTimeout * time.Millisecond)
	auditConfig.SetConnectionTimeout(configs.AuditConnectionTimeout * time.Millisecond)
	auditConfig.SetMaxWait(configs.AuditMaxWait * time.Millisecond)
	auditConfig.SetMaxConnections(configs.AuditMaxConnections)
	auditConfig.SetMaxConnectionsPerRoute(configs.AuditMaxConnectionsPerRoute)
	auditConfig.SetMaxRetries(configs.AuditMaxRetries)
	auditConfig.SetRetryDelay(configs.AuditRetryDelay * time.Millisecond)
	auditClient = audittrail.MakeAuditClient(configs.AuditName, auditConfig)
}

//SaveAuditSync ...
func SaveAuditSync(eventName string, userName string, resourceType string, resourceID string, currentData map[string]interface{}, previousData map[string]interface{}, requestID string, tags []string) error {
	auditLog := audittrail.MakeAudit(eventName, userName, resourceType)
	auditLog.SetResourceID(resourceID)
	auditLog.SetCurrentData(currentData)
	auditLog.SetPreviousData(previousData)
	auditLog.SetRequestID(requestID)
	auditLog.SetTags(tags)

	//ATTENTION - this id is auto generated, you can change it, but not recommended, use the SetResourceID function to customize an ID
	//auditLog.SetID("NOT_CHANGE")

	err := auditClient.SaveAuditSync(auditLog)
	if err != nil {
		return err
	}

	return nil
}
