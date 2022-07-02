package handlers

import (
	"cloudsign/fakeserver/fakedata"
	"cloudsign/fakeserver/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DocumentsDocumentIDAttributeGet -
func (c *Container) DocumentsDocumentIDAttributeGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDAttributePut -
func (c *Container) DocumentsDocumentIDAttributePut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDCertificateGet -
func (c *Container) DocumentsDocumentIDCertificateGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDDeclinePut -
func (c *Container) DocumentsDocumentIDDeclinePut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDDelete -
func (c *Container) DocumentsDocumentIDDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDFilesFileIDDelete -
func (c *Container) DocumentsDocumentIDFilesFileIDDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDFilesFileIDGet -
func (c *Container) DocumentsDocumentIDFilesFileIDGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDFilesFileIDWidgetsPost -
func (c *Container) DocumentsDocumentIDFilesFileIDWidgetsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDFilesFileIDWidgetsWidgetIDDelete -
func (c *Container) DocumentsDocumentIDFilesFileIDWidgetsWidgetIDDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDFilesFileIDWidgetsWidgetIDPut -
func (c *Container) DocumentsDocumentIDFilesFileIDWidgetsWidgetIDPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDFilesPost -
func (c *Container) DocumentsDocumentIDFilesPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDGet -
func (c *Container) DocumentsDocumentIDGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDParticipantsParticipantIDDelete -
func (c *Container) DocumentsDocumentIDParticipantsParticipantIDDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDParticipantsParticipantIDPut -
func (c *Container) DocumentsDocumentIDParticipantsParticipantIDPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDParticipantsPost -
func (c *Container) DocumentsDocumentIDParticipantsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDPost -
func (c *Container) DocumentsDocumentIDPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDPut -
func (c *Container) DocumentsDocumentIDPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDReporteesPost -
func (c *Container) DocumentsDocumentIDReporteesPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDReporteesReporteeIDDelete -
func (c *Container) DocumentsDocumentIDReporteesReporteeIDDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsDocumentIDReporteesReporteeIDPut -
func (c *Container) DocumentsDocumentIDReporteesReporteeIDPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DocumentsGet -
func (c *Container) DocumentsGet(ctx echo.Context) error {
	model, err := fakedata.GenerateFakeData(models.DocumentListModel{})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, model)

}

// DocumentsPost -
func (c *Container) DocumentsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// MeGet -
func (c *Container) MeGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// TeamDocumentsGet -
func (c *Container) TeamDocumentsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// TokenGet -
// Deprecated
func (c *Container) TokenGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// TokenPost -
func (c *Container) TokenPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
