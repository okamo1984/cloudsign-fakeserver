package main

import (
	"cloudsign/fakeserver/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// DocumentsDocumentIDAttributeGet -
	e.GET("/documents/:documentID/attribute", c.DocumentsDocumentIDAttributeGet)

	// DocumentsDocumentIDAttributePut -
	e.PUT("/documents/:documentID/attribute", c.DocumentsDocumentIDAttributePut)

	// DocumentsDocumentIDCertificateGet -
	e.GET("/documents/:documentID/certificate", c.DocumentsDocumentIDCertificateGet)

	// DocumentsDocumentIDDeclinePut -
	e.PUT("/documents/:documentID/decline", c.DocumentsDocumentIDDeclinePut)

	// DocumentsDocumentIDDelete -
	e.DELETE("/documents/:documentID", c.DocumentsDocumentIDDelete)

	// DocumentsDocumentIDFilesFileIDDelete -
	e.DELETE("/documents/:documentID/files/:fileID", c.DocumentsDocumentIDFilesFileIDDelete)

	// DocumentsDocumentIDFilesFileIDGet -
	e.GET("/documents/:documentID/files/:fileID", c.DocumentsDocumentIDFilesFileIDGet)

	// DocumentsDocumentIDFilesFileIDWidgetsPost -
	e.POST("/documents/:documentID/files/:fileID/widgets", c.DocumentsDocumentIDFilesFileIDWidgetsPost)

	// DocumentsDocumentIDFilesFileIDWidgetsWidgetIDDelete -
	e.DELETE("/documents/:documentID/files/:fileID/widgets/:widgetID", c.DocumentsDocumentIDFilesFileIDWidgetsWidgetIDDelete)

	// DocumentsDocumentIDFilesFileIDWidgetsWidgetIDPut -
	e.PUT("/documents/:documentID/files/:fileID/widgets/:widgetID", c.DocumentsDocumentIDFilesFileIDWidgetsWidgetIDPut)

	// DocumentsDocumentIDFilesPost -
	e.POST("/documents/:documentID/files", c.DocumentsDocumentIDFilesPost)

	// DocumentsDocumentIDGet -
	e.GET("/documents/:documentID", c.DocumentsDocumentIDGet)

	// DocumentsDocumentIDParticipantsParticipantIDDelete -
	e.DELETE("/documents/:documentID/participants/:participantID", c.DocumentsDocumentIDParticipantsParticipantIDDelete)

	// DocumentsDocumentIDParticipantsParticipantIDPut -
	e.PUT("/documents/:documentID/participants/:participantID", c.DocumentsDocumentIDParticipantsParticipantIDPut)

	// DocumentsDocumentIDParticipantsPost -
	e.POST("/documents/:documentID/participants", c.DocumentsDocumentIDParticipantsPost)

	// DocumentsDocumentIDPost -
	e.POST("/documents/:documentID", c.DocumentsDocumentIDPost)

	// DocumentsDocumentIDPut -
	e.PUT("/documents/:documentID", c.DocumentsDocumentIDPut)

	// DocumentsDocumentIDReporteesPost -
	e.POST("/documents/:documentID/reportees", c.DocumentsDocumentIDReporteesPost)

	// DocumentsDocumentIDReporteesReporteeIDDelete -
	e.DELETE("/documents/:documentID/reportees/:reporteeID", c.DocumentsDocumentIDReporteesReporteeIDDelete)

	// DocumentsDocumentIDReporteesReporteeIDPut -
	e.PUT("/documents/:documentID/reportees/:reporteeID", c.DocumentsDocumentIDReporteesReporteeIDPut)

	// DocumentsGet -
	e.GET("/documents", c.DocumentsGet)

	// DocumentsPost -
	e.POST("/documents", c.DocumentsPost)

	// MeGet -
	e.GET("/me", c.MeGet)

	// TeamDocumentsGet -
	e.GET("/team_documents", c.TeamDocumentsGet)

	// TokenGet -  (deprecated)
	e.GET("/token", c.TokenGet)

	// TokenPost -
	e.POST("/token", c.TokenPost)

	// Start server
	e.Logger.Fatal(e.Start(":18080"))
}
